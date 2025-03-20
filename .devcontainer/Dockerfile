FROM mcr.microsoft.com/devcontainers/base:bookworm   

RUN apt-get update && \
    apt-get install -y makepkg protobuf-compiler clang-format

RUN git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.1
ENV PATH="$PATH:/root/.asdf/bin"
ENV PATH="$PATH:/root/.asdf/shims"

RUN asdf plugin add golang && \
    asdf install golang 1.23.4 && \
    asdf global golang 1.23.4

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2 && \
    go install mvdan.cc/gofumpt@v0.6.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.4

ENV PATH="$PATH:/root/.asdf/installs/golang/1.23.4/packages/bin"