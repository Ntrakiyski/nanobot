package runtime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nanobot-ai/nanobot/pkg/agents"
	"github.com/nanobot-ai/nanobot/pkg/complete"
	"github.com/nanobot-ai/nanobot/pkg/confirm"
	"github.com/nanobot-ai/nanobot/pkg/llm"
	"github.com/nanobot-ai/nanobot/pkg/mcp"
	"github.com/nanobot-ai/nanobot/pkg/sampling"
	"github.com/nanobot-ai/nanobot/pkg/tools"
	"github.com/nanobot-ai/nanobot/pkg/types"
)

type Runtime struct {
	*tools.Service
	config    types.Config
	llmConfig llm.Config
	opt       Options
}

type Options struct {
	Confirmations  *confirm.Service
	Roots          []mcp.Root
	Profiles       []string
	MaxConcurrency int
}

func (o Options) Merge(other Options) (result Options) {
	result.Confirmations = complete.Last(o.Confirmations, other.Confirmations)
	result.MaxConcurrency = complete.Last(o.MaxConcurrency, other.MaxConcurrency)
	result.Profiles = append(o.Profiles, other.Profiles...)
	result.Roots = append(o.Roots, other.Roots...)
	return
}

func NewRuntime(cfg llm.Config, config types.Config, opts ...Options) *Runtime {
	opt := complete.Complete(opts...)
	completer := llm.NewClient(cfg, config)
	registry := tools.NewToolsService(config, tools.RegistryOptions{
		Roots:       opt.Roots,
		Concurrency: opt.MaxConcurrency,
	})
	agents := agents.New(completer, registry, opt.Confirmations, config)
	sampler := sampling.NewSampler(config, agents)

	// This is a circular dependency. Oh well, so much for good design.
	registry.SetSampler(sampler)

	return &Runtime{
		config:    config,
		Service:   registry,
		llmConfig: cfg,
		opt:       opt,
	}
}

func (r *Runtime) Reload(cfg types.Config) {
	newRuntime := NewRuntime(r.llmConfig, cfg, r.opt)
	r.config = cfg
	r.Service = newRuntime.Service
}

func (r *Runtime) GetConfig() types.Config {
	return r.config
}

func (r *Runtime) WithTempSession(ctx context.Context) context.Context {
	return mcp.WithSession(ctx, mcp.NewEmptySession(ctx, "temp"))
}

func (r *Runtime) getToolFromRef(ctx context.Context, serverRef string) (*tools.ListToolsResult, error) {
	var (
		server, tool string
	)

	toolRef := strings.Split(serverRef, "/")
	if len(toolRef) == 1 {
		_, ok := r.config.Agents[toolRef[0]]
		if ok {
			server, tool = toolRef[0], toolRef[0]
		} else {
			server, tool = "", toolRef[0]
		}
	} else if len(toolRef) == 2 {
		server, tool = toolRef[0], toolRef[1]
	} else {
		return nil, fmt.Errorf("invalid tool reference: %s", serverRef)
	}

	toolList, err := r.ListTools(ctx, tools.ListToolsOptions{
		Servers: []string{server},
		Tools:   []string{tool},
	})
	if err != nil {
		return nil, err
	}

	if len(toolList) != 1 || len(toolList[0].Tools) != 1 {
		return nil, fmt.Errorf("found %d tools with name %s on server %s", len(toolList), tool, server)
	}

	return &tools.ListToolsResult{
		Server: toolList[0].Server,
		Tools:  []mcp.Tool{toolList[0].Tools[0]},
	}, nil
}

func (r *Runtime) CallFromCLI(ctx context.Context, serverRef string, args ...string) (*mcp.CallToolResult, error) {
	var (
		argValue any
		argMap   = map[string]string{}
	)

	tools, err := r.getToolFromRef(ctx, serverRef)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(tools.Tools[0].InputSchema, types.ChatInputSchema) {
		argValue = types.SampleCallRequest{
			Prompt: strings.Join(args, " "),
		}
		args = nil
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "--") {
			if len(args) > 1 {
				return nil, fmt.Errorf("if using JSON syntax you must pass one argument: got %d", len(args))
			}
			err := json.Unmarshal([]byte(arg), &argValue)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
			}
			break
		}
		k, v, ok := strings.Cut(arg, "=")
		if !ok {
			if i+1 >= len(args) {
				return nil, fmt.Errorf("missing value for argument %q", arg)
			}
			v = args[i+1]
			i++
		}
		argMap[strings.TrimPrefix(k, "--")] = v
		argValue = argMap
	}

	if argValue == nil {
		argValue = map[string]any{}
	}

	return r.Call(ctx, tools.Server, tools.Tools[0].Name, argValue)
}
