.PHONY: build
build: proto install-rest proto-example

test-runtime:
	go test ./runtime/... -coverprofile=.coverage
	go tool cover -func=.coverage

example-run:
	go run ./example/main.go

install-rest:
	go install ./cmd/protoc-gen-go-rest

proto:
	protoc \
		-I=restapi \
		--go_out=paths=source_relative:restapi \
		restapi/annotations.proto

proto-example:
	rm -rf example/api/*
	protoc \
		-I=. \
		--go_out=example \
		--go-rest_out=example \
		example/proto/*.proto
	swag fmt -d example/api/example_rest.pb.go
	swag init --generalInfo example/api/example_rest.pb.go --output example/api/swagger

lint:
	go mod tidy
	golangci-lint run --fix --show-stats --max-issues-per-linter 0 --max-same-issues 0