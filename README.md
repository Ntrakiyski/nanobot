![Nanobot](docs/header.svg)

# Nanobot - Build MCP Agents

Nanobot enables building agents with MCP and MCP-UI by providing a flexible MCP host.
While existing applications like VSCode, Claude, Cursor, ChatGPT, and Goose all include an MCP host,
Nanobot is designed to be a standalone, open-source MCP host that can be easily deployed or integrated into
your applications. You can use Nanobot to create your own dedicated MCP and MCP-UI powered chatbot.

## What is an MCP Host?

An [MCP host](https://modelcontextprotocol.io/specification/2025-06-18/architecture/index#host) is
the service that combines MCP servers with an LLM and context to present an agent experience to a
consumer. The primary experience today is a chat interface, but it can be many other interfaces such
as voice, SMS, e-mail, AR/VR, Slack, MCP, or any other interface that can be used to interact with
an agent.

![MCP host](docs/mcp-host.png)

## Examples

Here are some examples of Nanobots in action:
- [Blackjack Game](https://blackjack.nanobot.ai) [(config)](./examples/blackjack.yaml) [(mcp source)](https://github.com/nanobot-ai/blackjack)
- [Hugging Face MCP](https://huggingface.nanobot.ai) [(config)](./examples/huggingface.yaml)
- [Shopping/Shopify](https://shopping.nanobot.ai) [(config)](./examples/shopping.yaml)

## Installation

Nanobot can be installed via [Homebrew](https://brew.sh/):

```bash
brew install nanobot-ai/tap/nanobot
```

This will give you the `nanobot` CLI, which you can use to run and manage your MCP host.

---

## Getting Started

---

## Configuration

Nanobot supports the following providers:

- **OpenAI** (e.g. `gpt-4`)
- **Anthropic** (e.g. `claude-3`)

To use them, set the corresponding API key:

```bash
# For OpenAI models
export OPENAI_API_KEY=sk-...

# For Anthropic models
export ANTHROPIC_API_KEY=sk-ant-...
```

Nanobot automatically selects the correct provider based on the model specified.

---

Create a configuration file (e.g. `nanobot.yaml`) that defines your agents and MCP servers.

**Example:**

```yaml
agents:
  dealer:
    name: Blackjack Dealer
    model: gpt-4.1
    mcpServers: blackjackmcp

mcpServers:
  blackjackmcp:
    url: https://blackjack.nanobot.ai/mcp
```

Start Nanobot with:

```bash
nanobot run ./nanobot.yaml
```

The UI will be available at [http://localhost:8080](http://localhost:8080).

---

## Development & Contribution

Contributions are welcome! Nanobot is still in **alpha**, so expect active development and rapid changes.

### Build from Source

```bash
make
```

### Working on the UI

The Nanobot UI lives in the `./ui` directory. To develop against it:

1. Remove the old build artifacts:

   ```bash
   rm -rf ./ui/dist
   ```

2. Rebuild the Nanobot binary:

   ```bash
   make
   ```

3. Start the UI in development mode:

   ```bash
   cd ui
   npm run dev
   ```

4. The UI must be served from port **5173**.\
   Nanobot runs on port **8080** and will forward UI requests to `:5173`.

---

## Features & Roadmap

Nanobot aims to be a **fully compliant MCP Host** and support all MCP + MCP-UI features.

| Feature Category           | Feature | Status        |
| -------------------------- |---------| ------------- |
| **MCP Core**               | TODO    | ✅ Implemented |
|                            | TODO    | 🚧 Partial    |
|                            | TODO    | ❌ Not yet     |
|                            | TODO    | ✅ Implemented |
| **MCP-UI**                 | TODO    | 🚧 Partial    |
|                            | TODO    | ✅ Implemented |
|                            | TODO    | ❌ Not yet     |

✅ = Implemented 🚧 = Partial / WIP ❌ = Not yet ⏳ = Planned

### Roadmap

- Full MCP + MCP-UI compliance
- More robust multi-agent support
- Production-ready UI
- Expanded model provider support
- Expanded authentication and security features
- Frontend integrations (Slack, SMS, email, embedded web agents)
- Easy embedding into existing apps and websites

---

## License

Nanobot is licensed under the [Apache 2.0 License](LICENSE).

---

## Links

- Website: [nanobot.ai](https://nanobot.ai)
- GitHub: [github.com/nanobot-ai/nanobot](https://github.com/nanobot-ai/nanobot)

# Comprehensive Guide: Building and Deploying Multi-Client AI Agents with Nanobot

## Overview

This comprehensive guide provides detailed instructions for building and deploying multi-client AI agents using the nanobot-ai/nanobot framework, covering Docker hosting, secure credential management, and usage tracking.

## Step 1: Environment Setup and Installation

### Task 1.1: Install Nanobot Framework [1](#5-0) 

### Task 1.2: Verify Installation Requirements
- Ensure Go 1.19+ is installed for building from source
- Verify Docker is installed for containerized deployments
- Check Node.js/npm availability for MCP server packages

### Task 1.3: Create Project Directory Structure
```
project/
├── nanobot.yaml
├── nanobot.env
├── profiles/
├── mcps/
└── docker/
```

### Task 1.4: Initialize Base Configuration [2](#5-1) 

### Task 1.5: Set Up Environment Variables [3](#5-2) 

**💡 Tips:**
- Use the curl installation method for quick setup on Linux/macOS
- Always verify installation with `nanobot --version`

**⚠️ Must-Do:**
- Create isolated project directories for each agent deployment
- Always use version tags in production deployments

## Step 2: Secure Credential Management

### Task 2.1: Configure API Credentials [4](#5-3) 

### Task 2.2: Set Up Environment File [5](#5-4) 

### Task 2.3: Configure Credential Isolation [6](#5-5) 

### Task 2.4: Implement Environment Variable Substitution [7](#5-6) 

### Task 2.5: Validate Credential Security
- Test credential loading without exposing sensitive values
- Verify environment isolation between different agent instances

**💡 Tips:**
- Use separate `.env` files for different environments (dev, staging, prod)
- Leverage Docker secrets for production deployments

**🚫 Must-Not-Do:**
- Never commit API keys to version control
- Avoid hardcoding credentials in configuration files

## Step 3: MCP Server Configuration

### Task 3.1: Define MCP Servers [8](#5-7) 

### Task 3.2: Configure Server Parameters [7](#5-6) 

### Task 3.3: Set Up Local MCP Servers
- Install filesystem MCP server: `npx -y @modelcontextprotocol/server-filesystem`
- Configure server paths and permissions
- Test server connectivity

### Task 3.4: Configure Remote MCP Servers [7](#5-6) 

### Task 3.5: Validate MCP Server Integration [9](#5-8) 

**💡 Tips:**
- Use environment variables for dynamic path configuration
- Test MCP servers independently before integration

**⚠️ Must-Do:**
- Always validate MCP server availability before agent deployment
- Implement proper error handling for server connection failures

## Step 4: Agent Configuration and Profiles

### Task 4.1: Define Base Agent Configuration [10](#5-9) 

### Task 4.2: Configure Agent Instructions [11](#5-10) 

### Task 4.3: Set Up Client Profiles [6](#5-5) 

### Task 4.4: Configure Profile Inheritance [12](#5-11) 

### Task 4.5: Validate Agent Configuration [13](#5-12) 

**💡 Tips:**
- Use profiles for environment-specific overrides
- Implement dynamic instructions using prompts for flexibility

**⚠️ Must-Do:**
- Always define an entrypoint in the publish section
- Test agent configurations before deployment

## Step 5: Multi-Client Publishing Setup

### Task 5.1: Configure Publishing Options [14](#5-13) 

### Task 5.2: Define Tool Publishing [15](#5-14) 

### Task 5.3: Configure Resource Publishing [16](#5-15) 

### Task 5.4: Set Up Prompt Publishing [17](#5-16) 

### Task 5.5: Configure Server Metadata [18](#5-17) 

**💡 Tips:**
- Use descriptive names and versions for published agents
- Document all published tools and resources

**🚫 Must-Not-Do:**
- Don't publish sensitive internal tools to external clients
- Avoid version conflicts in published resources

## Step 6: Usage Tracking and Logging

### Task 6.1: Configure Debug Logging [19](#5-18) 

### Task 6.2: Set Up Message Tracking [20](#5-19) 

### Task 6.3: Implement Usage Metrics [21](#5-20) 

### Task 6.4: Configure Log Forwarding [22](#5-21) 

### Task 6.5: Set Up Monitoring
- Configure log aggregation systems
- Set up alerts for error rates
- Implement usage dashboards

**💡 Tips:**
- Use `NANOBOT_DEBUG=messages,progress,log` for comprehensive debugging
- Implement structured logging for better analysis

**⚠️ Must-Do:**
- Always monitor agent performance and usage patterns
- Implement proper log rotation for production deployments

## Step 7: Docker Containerization

### Task 7.1: Create Dockerfile
```dockerfile
FROM ghcr.io/nanobot-ai/nanobot:latest
COPY nanobot.yaml /app/
COPY nanobot.env /app/
WORKDIR /app
ENTRYPOINT ["nanobot", "run", ".", "--mcp"]
```

### Task 7.2: Configure Multi-Stage Build
- Use official Nanobot base image
- Copy only necessary configuration files
- Set appropriate working directory

### Task 7.3: Set Up Environment Variables [23](#5-22) 

### Task 7.4: Configure Health Checks
- Implement readiness probes
- Set up liveness checks
- Configure startup timeouts

### Task 7.5: Optimize Image Size
- Use multi-stage builds for reduced image size
- Remove unnecessary dependencies
- Implement proper layer caching

**💡 Tips:**
- Use official base images for security and stability
- Implement proper signal handling for graceful shutdowns

**⚠️ Must-Do:**
- Always use specific version tags, never `latest` in production
- Implement proper secret management in containers

## Step 8: Container Orchestration and Deployment

### Task 8.1: Create Docker Compose Configuration
```yaml
version: '3.8'
services:
  nanobot-agent:
    image: my-nanobot-agent:latest
    environment:
      - OPENAI_API_KEY=${OPENAI_API_KEY}
    ports:
      - "8099:8099"
```

### Task 8.2: Configure Service Discovery
- Set up container networking
- Configure service mesh if required
- Implement load balancing

### Task 8.3: Set Up Volume Management
- Configure persistent volumes for logs
- Set up configuration file mounting
- Implement backup strategies

### Task 8.4: Configure Container Networking [24](#5-23) 

### Task 8.5: Implement Scaling Strategies
- Configure horizontal pod autoscaling
- Set up resource limits and requests
- Implement circuit breakers

**💡 Tips:**
- Use health checks for proper container lifecycle management
- Implement graceful shutdown handling

**🚫 Must-Not-Do:**
- Don't run containers as root user
- Avoid exposing internal ports unnecessarily

## Step 9: MCP Server Management

### Task 9.1: Configure Server Lifecycle [9](#5-8) 

### Task 9.2: Set Up Server Communication [25](#5-24) 

### Task 9.3: Implement Error Handling [26](#5-25) 

### Task 9.4: Configure Server Routing [27](#5-26) 

### Task 9.5: Set Up Server Monitoring
- Monitor server health and availability
- Track connection metrics
- Implement automatic reconnection

**💡 Tips:**
- Use stdio for local development, HTTP for production
- Implement proper timeout handling

**⚠️ Must-Do:**
- Always validate server configuration before startup
- Implement proper error recovery mechanisms

## Step 10: Production Deployment and Monitoring

### Task 10.1: Configure Production Environment [28](#5-27) 

### Task 10.2: Set Up Load Balancing
- Configure reverse proxy (nginx/traefik)
- Implement session affinity if needed
- Set up SSL termination

### Task 10.3: Configure Monitoring and Alerting [29](#5-28) 

### Task 10.4: Implement Backup and Recovery
- Set up configuration backups
- Implement disaster recovery procedures
- Test recovery processes regularly

### Task 10.5: Configure Auto-scaling
- Set up metrics-based scaling
- Configure resource thresholds
- Implement cost optimization strategies

**💡 Tips:**
- Use blue-green deployments for zero-downtime updates
- Implement comprehensive monitoring from day one

**⚠️ Must-Do:**
- Always test deployments in staging environments first
- Implement proper rollback procedures

## Configuration Templates

### Basic nanobot.yaml Template [2](#5-1) 

### Advanced Configuration Template [30](#5-29) 

## Security Best Practices

1. **Credential Management**: [4](#5-3) 
2. **Environment Isolation**: [23](#5-22) 
3. **Profile Security**: [12](#5-11) 

## Troubleshooting Guide

### Common Issues and Solutions

1. **Configuration Loading Errors**: [31](#5-30) 
2. **MCP Server Connection Issues**: [9](#5-8) 
3. **Environment Variable Problems**: [3](#5-2) 

## Performance Optimization

### Resource Management
- Configure appropriate memory limits
- Set optimal CPU allocations
- Implement caching strategies

### Scaling Considerations [32](#5-31) 

## Notes

This documentation is based on the current Nanobot framework architecture and provides comprehensive coverage of multi-client AI agent deployment. The framework is designed to be lightweight and leverage the Model Context Protocol (MCP) for extensibility. Key features include profile-based configuration management, secure credential handling, and flexible deployment options supporting both stdio and HTTP-based MCP servers.

The configuration system supports inheritance through the `extends` mechanism and profile-based overrides for different environments. Environment variable substitution provides flexibility in configuration management, while the logging system offers comprehensive debugging and monitoring capabilities.

For production deployments, consider implementing proper monitoring, backup strategies, and security measures as outlined in this guide. The framework's modular design allows for easy scaling and customization based on specific requirements.

### Citations

**File:** README.md (L11-27)
```markdown
```yaml
# Example nanobot.yaml

publish:
  entrypoint: fs-agent

agents:
  fs-agent:
    model: gpt-4o
    instructions: You're awesome at file systems things, help the user with files... and things.
    tools: ["filesystem"]

mcpServers:
  filesystem:
    command: npx
    args: ["-y", "@modelcontextprotocol/server-filesystem", "${ROOT_CWD}"]
```
```

**File:** README.md (L37-48)
```markdown
Install with brew (macOS/Linux):

```shell
brew install nanobot-ai/tap/nanobot
```

Install with curl (macOS/Linux):
```shell
curl -sSfL https://get.nanobot.ai | sh -
```

For Windows there are binaries in the [releases](https://github.com/nanobot-ai/nanobot/releases) but they aren't really tested.
```

**File:** README.md (L80-201)
```markdown
```yaml
publish:
  # The main agent to interactive with when using `nanobot run`. This can also be a tool and
  # should be in the format [MCP_SERVER]/[TOOL_NAME], for example `filesystem/ls`. If your
  # are using a tool, `nanobot run` will assume it accepts a string argument named "prompt"
  entrypoint: fs-agent
  
  # Tools to publish when ran as a MCP server. This can refer to tools or agents.
  # If you refer to a MCP Server it will publish all tools from that server.
  tools: [server1/tool2, agent2, server3]
  
  # Prompt to publish when ran as a MCP server. This can be all prompts
  # from a server or a specific prompt.
  prompts: [server2, server1/prompt1]
  
  # Resources to publish when ran as a MCP server. This can be all
  # resources from a server or a specific resource.
  resources: [server1/resource1, server2]
  
  # Resource templates to publish when ran as a MCP server. This can be all
  # resource templates from a server or a specific resource template.
  resourceTemplates: ["server1/resource-template1", server2]
  
  # The name of the MCP Server when running as a MCP server.
  name: My Agent
  
  # The description of the MCP Server when running as a MCP server.
  description: This is a description of my agent
  
  # The version of the MCP Server when running as a MCP server.
  version: 0.0.1
  
  # The instructions for the MCP Server when running as a MCP server.
  instructions: Use wisely and be nice
  
  # The introduction message show in the console run running `nanobot run`
  # This field can also refer to a prompt to create a dynamic messages.
  #   introduction:
  #     prompt: prompt1
  #     mcpServer: [arg1, arg2]
  #     args:
  #       field1: value1
  introduction: I'm a fancy agent that can do things. Nice to mean you
  
agents:
  myAgent:
    # The model name. Currently supported is OpenAI Responses API and
    # Anthropic API. The model string is passed as is to the API.
    model: gpt-4o
    
    # The instructions for the agent. This can be a static string or a dynamic
    # prompt.
    #   instructions:
    #     prompt: prompt1
    #     mcpServer: server1
    #     args:
    #       field1: value1
    instructions: You're awesome at things, help the user with things.
    
    # The tools to use for the agent. This can be a reference to a MCP Server,
    # a specific tool, or an agent.
    tools: [filesystem, server1/tool2, agent2]
    
    # Inform the model on how to choose a tool. This can be "auto", "none", or a specific
    # tool name. If set to "auto", the model will choose the best tool for the task.
    toolChoice: "auto"
    
    toolExtensions:
      # The tool name is the key and the value is a map of options to pass as in the tool API
      # field. Refer to the openai-computer-using-agent example for an example of using this.
      # This is a mechanism to support custom server side tools that are offered in the OpenAI
      # and Anthropic APIs.
      tool1:
        option1: value1
        option2: value2
        # This is a special option that will remove fields that are normal sent in the tool object
        # in the API.
        remove: [description]

    # Whether chat history is saved for this agent in the MCP session. If set to true,
    # which is the default, each invocation of the agent will be added to the same chat
    # thread. If set to false, each invocation will be a new chat thread.
    chatHistory: true
    
    # The temperature to use for the agent. This is a number between 0 and 1. The default
    # is dependent on the model API.
    temperature: 0.7
    
    # The top_p to use for the agent. This is a number between 0 and 1. The default
    # is dependent on the model API. It is always recommend to use top_p or temperature
    # but not both.
    topP: 0.9
    
    # The output schema to use for the agent. This is a map of field names to types.
    output:
      # Whether or not the output must strictly conform. This is only supported in OpenAI API
      strict: true
      schema: # Standard JSON Schema
        type: object
        properties: {}
    
    # Whether the chat history should automatically be truncated to match the context size.
    # Can be set to "auto" or "disabled". "disabled" is the default
    truncation: "auto"
    
    # The maximum number of output tokens allowed by default. This can be overridden by the
    # `maxTokens` field in the CreateMessage sampling request.
    maxToken: 1024
    
    # Aliases used to match model hint name in CreateMessage ModelPreferences. The name of the
    # agent is always used as the first alias.
    alias: ["model1", "model2"]
    
    # The cost of the agent used for the priority evaluation in the CreateMessage ModelPreferences.
    cost: 0.5
    
    # The speed of the agent used for the priority evaluation in the CreateMessage ModelPreferences.
    speed: 0.5
    
    # The intelligence of the agent used for the priority evaluation in the CreateMessage ModelPreferences.
    intelligence: 0.5
```
```

**File:** pkg/cli/root.go (L36-43)
```go
	EnvFile          string            `usage:"Path to the environment file (default: ./nanobot.env)" short:"e"`
	OpenAIAPIKey     string            `usage:"OpenAI API key" env:"OPENAI_API_KEY" name:"openai-api-key"`
	OpenAIBaseURL    string            `usage:"OpenAI API URL" env:"OPENAI_BASE_URL" name:"openai-base-url"`
	OpenAIHeaders    map[string]string `usage:"OpenAI API headers" env:"OPENAI_HEADERS" name:"openai-headers"`
	AnthropicAPIKey  string            `usage:"Anthropic API key" env:"ANTHROPIC_API_KEY" name:"anthropic-api-key"`
	AnthropicBaseURL string            `usage:"Anthropic API URL" env:"ANTHROPIC_BASE_URL" name:"anthropic-base-url"`
	AnthropicHeaders map[string]string `usage:"Anthropic API headers" env:"ANTHROPIC_HEADERS" name:"anthropic-headers"`
}
```

**File:** pkg/cli/root.go (L87-122)
```go
func (n *Nanobot) loadEnv() (map[string]string, error) {
	env := map[string]string{}
	if !n.EmptyEnv {
		for _, e := range os.Environ() {
			k, v, _ := strings.Cut(e, "=")
			env[k] = v
		}
	}

	defaultFile := n.EnvFile == ""
	if defaultFile {
		n.EnvFile = "./nanobot.env"
	}

	data, err := os.ReadFile(n.EnvFile)
	if errors.Is(err, fs.ErrNotExist) && defaultFile {
	} else if err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			k, v, _ := strings.Cut(line, "=")
			env[k] = v
		}
	}

	if _, ok := env["NANOBOT_MCP"]; !ok {
		env["NANOBOT_MCP"] = "true"
	}

	return env, nil
}
```

**File:** pkg/types/config.go (L10-17)
```go
type Config struct {
	Extends    string                   `json:"extends,omitempty"`
	Env        map[string]string        `json:"env,omitempty"`
	Publish    Publish                  `json:"publish,omitempty"`
	Agents     map[string]Agent         `json:"agents,omitempty"`
	MCPServers map[string]mcp.MCPServer `json:"mcpServers,omitempty"`
	Profiles   map[string]Config        `json:"profiles,omitempty"`
}
```

**File:** pkg/types/config.go (L19-30)
```go
type Publish struct {
	Name              string              `json:"name,omitempty"`
	Description       string              `json:"description,omitempty"`
	Introduction      DynamicInstructions `json:"introduction,omitempty"`
	Version           string              `json:"version,omitempty"`
	Instructions      string              `json:"instructions,omitempty"`
	Tools             StringList          `json:"tools,omitzero"`
	Prompts           StringList          `json:"prompts,omitzero"`
	Resources         StringList          `json:"resources,omitzero"`
	ResourceTemplates StringList          `json:"resourceTemplates,omitzero"`
	Entrypoint        string              `json:"entrypoint,omitempty"`
}
```

**File:** pkg/types/config.go (L91-111)
```go
type Agent struct {
	Description    string                    `json:"description,omitempty"`
	Instructions   DynamicInstructions       `json:"instructions,omitempty"`
	Model          string                    `json:"model,omitempty"`
	Tools          []string                  `json:"tools,omitempty"`
	ChatHistory    *bool                     `json:"chatHistory,omitempty"`
	ToolExtensions map[string]map[string]any `json:"toolExtensions,omitempty"`
	ToolChoice     string                    `json:"toolChoice,omitempty"`
	Temperature    *json.Number              `json:"temperature,omitempty"`
	TopP           *json.Number              `json:"topP,omitempty"`
	Output         *OutputSchema             `json:"output,omitempty"`
	Truncation     string                    `json:"truncation,omitempty"`
	MaxTokens      int                       `json:"maxTokens,omitempty"`

	// Selection criteria fields

	Aliases      []string `json:"aliases,omitempty"`
	Cost         float64  `json:"cost,omitempty"`
	Speed        float64  `json:"speed,omitempty"`
	Intelligence float64  `json:"intelligence,omitempty"`
}
```

**File:** pkg/types/config.go (L113-127)
```go
type DynamicInstructions struct {
	Instructions string            `json:"-"`
	MCPServer    string            `json:"mcpServer"`
	Prompt       string            `json:"prompt"`
	Args         map[string]string `json:"args"`
}

func (a DynamicInstructions) IsPrompt() bool {
	return a.MCPServer != "" && a.Prompt != ""
}

func (a DynamicInstructions) IsSet() bool {
	return a.IsPrompt() || a.Instructions != ""
}

```

**File:** pkg/mcp/client.go (L16-29)
```go
type Client struct {
	Session *Session
}

type ClientOption struct {
	OnSampling    func(ctx context.Context, sampling CreateMessageRequest) (CreateMessageResult, error)
	OnRoots       func(ctx context.Context, msg Message) error
	OnLogging     func(ctx context.Context, logMsg LoggingMessage) error
	OnMessage     func(ctx context.Context, msg Message) error
	OnNotify      func(ctx context.Context, msg Message) error
	Env           map[string]string
	ParentSession *Session
	SessionID     string
}
```

**File:** pkg/mcp/client.go (L31-37)
```go
type MCPServer struct {
	Env     map[string]string `json:"env,omitempty"`
	Command string            `json:"command,omitempty"`
	Args    []string          `json:"args,omitempty"`
	BaseURL string            `json:"baseURL,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}
```

**File:** pkg/mcp/client.go (L39-50)
```go
func toHandler(opts ClientOption) MessageHandler {
	return MessageHandlerFunc(func(ctx context.Context, msg Message) {
		if msg.Method == "sampling/createMessage" && opts.OnSampling != nil {
			var param CreateMessageRequest
			if err := json.Unmarshal(msg.Params, &param); err != nil {
				msg.SendUnknownError(ctx, fmt.Errorf("failed to unmarshal sampling/createMessage: %w", err))
				return
			}
			go func() {
				resp, err := opts.OnSampling(ctx, param)
				if err != nil {
					msg.SendUnknownError(ctx, fmt.Errorf("failed to handle sampling/createMessage: %w", err))
```

**File:** pkg/cli/run.go (L61-95)
```go
func (r *Run) getRoots() ([]mcp.Root, error) {
	var (
		rootDefs = r.Roots
		roots    []mcp.Root
	)

	if len(rootDefs) == 0 {
		rootDefs = []string{"cwd:."}
	}

	for _, root := range rootDefs {
		name, directory, ok := strings.Cut(root, ":")
		if !ok {
			name = filepath.Base(root)
			directory = root
		}
		if !filepath.IsAbs(directory) {
			wd, err := os.Getwd()
			if err != nil {
				return nil, fmt.Errorf("failed to get current working directory: %w", err)
			}
			directory = filepath.Join(wd, directory)
		}
		if _, err := os.Stat(directory); err != nil {
			return nil, fmt.Errorf("failed to stat directory root (%s): %w", name, err)
		}

		roots = append(roots, mcp.Root{
			Name: name,
			URI:  "file://" + directory,
		})
	}

	return roots, nil
}
```

**File:** pkg/cli/run.go (L125-139)
```go

	if c := runtime.GetConfig(); c.Publish.Entrypoint == "" {
		var (
			agentName string
			example   string
		)
		for name := range c.Agents {
			agentName = name
			break
		}
		if agentName != "" {
			example = ", for example:\n\npublish:\n  entrypoint: " + agentName + "\nagents:\n  " + agentName + ": ...\n"
		}
		return fmt.Errorf("there are no entrypoints defined in the config file, please add one to the publish section%s", example)
	}
```

**File:** pkg/cli/run.go (L141-165)
```go
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return fmt.Errorf("failed to pick a local port: %w", err)
	}
	r.ListenAddress = l.Addr().String()

	prompt := strings.Join(args[1:], " ")
	if r.Input != "" {
		input, err := os.ReadFile(r.Input)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}
		prompt = strings.TrimSpace(string(input))
	}

	eg, ctx := errgroup.WithContext(cmd.Context())
	ctx, cancel := context.WithCancel(ctx)
	eg.Go(func() error {
		return r.runMCP(ctx, runtime, l)
	})
	eg.Go(func() error {
		defer cancel()
		return chat.Chat(ctx, r.ListenAddress, runtimeOpt.Confirmations, r.AutoConfirm, prompt, r.Output)
	})
	return eg.Wait()
```

**File:** pkg/cli/run.go (L168-186)
```go
func (r *Run) runMCP(ctx context.Context, runtime *runtime.Runtime, l net.Listener) error {
	address := r.ListenAddress
	if strings.HasPrefix("address", "http://") {
		address = strings.TrimPrefix(address, "http://")
	} else if strings.HasPrefix(address, "https://") {
		return fmt.Errorf("https:// is not supported, use http:// instead")
	}

	mcpServer := server.NewServer(runtime)

	if address == "stdio" {
		stdio := mcp.NewStdioServer(mcpServer)
		if err := stdio.Start(ctx, os.Stdin, os.Stdout); err != nil {
			return fmt.Errorf("failed to start stdio server: %w", err)
		}

		stdio.Wait()
		return nil
	}
```

**File:** pkg/cli/run.go (L187-201)
```go

	httpServer := mcp.HTTPServer{
		MessageHandler: mcpServer,
	}

	s := &http.Server{
		Addr:    address,
		Handler: &httpServer,
	}

	context.AfterFunc(ctx, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = s.Shutdown(ctx)
	})
```

**File:** pkg/config/load.go (L18-35)
```go
func Load(path string, profiles ...string) (*types.Config, string, error) {
	var (
		last types.Config
		data []byte
		dir  = "."
	)

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, "", err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
```

**File:** pkg/config/load.go (L57-66)
```go
	if last.Extends != "" {
		parent, _, err := Load(filepath.Join(dir, last.Extends), profiles...)
		if err != nil {
			return nil, "", fmt.Errorf("error reading %s: %w", last.Extends, err)
		}
		last, err = merge(*parent, last)
		if err != nil {
			return nil, "", fmt.Errorf("error merging %s: %w", last.Extends, err)
		}
	}
```

**File:** pkg/config/load.go (L67-81)
```go

	for _, profile := range profiles {
		profile, _, optional := strings.Cut(profile, "?")
		profileConfig, found := last.Profiles[profile]
		if !found && !optional {
			return nil, "", fmt.Errorf("profile %s not found", profile)
		} else if !found {
			continue
		}
		var err error
		last, err = merge(last, profileConfig)
		if err != nil {
			return nil, "", fmt.Errorf("error merging profile %s: %w", profile, err)
		}
	}
```

**File:** pkg/config/load.go (L83-91)
```go
	seen := map[string]string{}
	if err := checkDup(seen, "mcpServer", slices.Collect(maps.Keys(last.MCPServers))); err != nil {
		return nil, "", err
	}
	if err := checkDup(seen, "agent", slices.Collect(maps.Keys(last.Agents))); err != nil {
		return nil, "", err
	}

	return &last, dir, nil
```

**File:** pkg/log/log.go (L13-20)
```go
var (
	debugs            = strings.Split(os.Getenv("NANOBOT_DEBUG"), ",")
	EnableMessages    = slices.Contains(debugs, "messages")
	EnableProgress    = slices.Contains(debugs, "progress")
	DebugLog          = slices.Contains(debugs, "log")
	Base64Replace     = regexp.MustCompile(`((;base64,|")[a-zA-Z0-9+/=]{60})[a-zA-Z0-9+/=]+"`)
	Base64Replacement = []byte(`$1..."`)
)
```

**File:** pkg/log/log.go (L22-35)
```go
func Messages(_ context.Context, server string, out bool, data []byte) {
	if EnableProgress && bytes.Contains(data, []byte(`"notifications/progress"`)) {
	} else if EnableMessages && !bytes.Contains(data, []byte(`"notifications/progress"`)) {
	} else {
		return
	}

	fmtStr := "->(%s) %s\n"
	if !out {
		fmtStr = "<-(%s) %s\n"
	}
	data = Base64Replace.ReplaceAll(data, Base64Replacement)
	_, _ = fmt.Fprintf(os.Stderr, fmtStr, server, strings.ReplaceAll(strings.TrimSpace(string(data)), "\n", " "))
}
```

**File:** pkg/log/log.go (L37-50)
```go
func Errorf(_ context.Context, format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
}

func Infof(_ context.Context, format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "INFO: "+format+"\n", args...)
}

func Debugf(_ context.Context, format string, args ...any) {
	if !DebugLog {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
}
```

