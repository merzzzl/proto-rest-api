.PHONY: build
build:
	buf generate
	go build -o ./bin/protoc-gen-go-rest ./cmd/protoc-gen-go-rest
	go build -o ./bin/protoc-gen-go-oapi ./cmd/protoc-gen-go-oapi
	cd ./example && buf generate

test:
	go test ./... -cover

example-run:
	go run ./example/main.go