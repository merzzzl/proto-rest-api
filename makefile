.PHONY: build
build: proto install-oapi install-rest proto-example

test-runtime:
	go test ./runtime/... -coverprofile=.coverage
	go tool cover -func=.coverage
example-run:
	go run ./example/main.go

install-oapi:
	go install ./cmd/protoc-gen-go-oapi

install-rest:
	go install ./cmd/protoc-gen-go-rest

proto:
	protoc \
		-I=restapi \
		--go_out=paths=source_relative:restapi \
		restapi/annotations.proto

proto-example:
	rm -f example/api/*
	protoc \
		-I=. \
		--go_out=example \
		--go-oapi_out=example \
		--go-rest_out=example \
		example/proto/*.proto