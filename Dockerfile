# Based on debian
FROM docker.io/library/debian:stable-slim
LABEL description="KubeEdge WeChat App"

# Copy from build directory
COPY app /app

# Copy TLS
COPY server.crt /server.crt
COPY server.key /server.key

# The https prot
EXPOSE 443

# WORKDIR
WORKDIR /

# Define default command
ENTRYPOINT ["/app"]

# Run the executable
CMD ["app"]