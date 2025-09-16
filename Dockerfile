# syntax=docker/dockerfile:1
FROM cgr.dev/chainguard/wolfi-base:latest

# Copy the binary
COPY nanobot /usr/local/bin/nanobot

# Create non-root user
RUN adduser -D -s /bin/sh nanobot

# Create data directory and set ownership
RUN mkdir -p /data && chown nanobot:nanobot /data

USER nanobot

# Set common env vars
ENV NANOBOT_STATE=/data/nanobot.db
ENV NANOBOT_RUN_HEALTHZ_PATH=/api/healthz
ENV NANOBOT_RUN_LISTEN_ADDRESS=0.0.0.0:8080

# Define volume for persistent data
VOLUME ["/data"]

ENTRYPOINT ["/usr/local/bin/nanobot"]
CMD ["run"]
