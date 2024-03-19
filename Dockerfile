FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y wget git ca-certificates --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Use dpkg to detect architecture and download the corresponding Golang binary
RUN ARCH="$(dpkg --print-architecture)" && \
    if [ "$ARCH" = "amd64" ]; then \
    GOARCH="amd64"; \
    elif [ "$ARCH" = "arm64" ]; then \
    GOARCH="arm64"; \
    else \
    echo "Unsupported architecture: $ARCH"; \
    exit 1; \
    fi && \
    wget "https://go.dev/dl/go1.20.linux-$GOARCH.tar.gz" -O go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="$GOPATH/bin:$PATH"

WORKDIR $GOPATH

CMD ["sleep", "infinity"]
