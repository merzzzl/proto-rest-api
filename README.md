# Proto REST API Generator

A tool for automatic generation of RESTful APIs based on Protocol Buffers (Protobuf) definitions. The project uses the `protoc-gen-go-rest` plugin to create OpenAPI specifications and other API artifacts.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Usage](#usage)
- [REST API Annotations](#rest-api-annotations)
- [Examples](#examples)
  - [Basic Example](#basic-example)
  - [Advanced Features](#advanced-features)
- [Make Commands](#make-commands)
- [Streaming API Support](#streaming-api-support)
- [Authentication](#authentication)
- [Contributing](#contributing)
- [License](#license)

## Features

- **OpenAPI 3.0 Specification Generation**: Automatically creates OpenAPI documentation based on Protobuf service definitions.
- **Flexible Request and Response Handling**: Maps Protobuf messages to REST request bodies, query parameters, and path parameters.
- **Streaming API Support**: Ability to create streaming APIs with WebSocket support.
- **Customizable HTTP Methods**: Support for GET, POST, PUT, PATCH, and DELETE.
- **Built-in Authentication**: Support for Bearer tokens and customizable access scopes.
- **Customizable HTTP Status Codes**: Ability to specify successful response codes.

## Prerequisites

- Go 1.23 or later
- Protocol Buffers Compiler (`protoc`) version 3.15 or later
- Installed `protoc-gen-go` and `protoc-gen-go-rest` plugins

## Quick Start

1. Install the necessary dependencies:
   ```bash
   sudo apt install -y protobuf-compiler
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest@latest
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Create a simple proto file:
   ```proto
   syntax = "proto3";
   package example;
   
   import "restapi/annotations.proto";
   
   option go_package = "./api;api";
   
   service ExampleService {
     option (restapi.service) = {
       base_path : "/api/v1"
     };
   
     rpc Echo(EchoRequest) returns (EchoResponse) {
       option (restapi.method) = {
         method : "GET"
         path : "/echo/:message"
         response : "*"
       };
     }
   }
   
   message EchoRequest {
     string message = 1;
   }
   
   message EchoResponse {
     string message = 1;
   }
   ```

3. Generate the REST API:
   ```bash
   protoc --go_out=. --go-rest_out=. --proto_path=. example.proto
   ```

4. Run the generated server and explore the REST API.

## Installation

1. Install the Protobuf compiler:
   ```bash
   sudo apt install -y protobuf-compiler
   ```

2. Install the required Go plugins:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest@latest
   ```

3. Ensure the plugins are in your `PATH`:
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

## Usage

1. Define your Protobuf services and messages in `.proto` files using annotations from the `restapi` package.

2. Generate the REST API and OpenAPI specification:
   ```bash
   protoc --go_out=. --go-rest_out=. --proto_path=./path/to/protos your_service.proto
   ```

3. The generated OpenAPI specification can be found in the output directory.

## REST API Annotations

The project uses special annotations to describe the REST API:

- `restapi.service` - settings for the entire service:
  - `base_path` - base path for all service endpoints
  - `auth` - authentication type (NONE, BEARER)
  - `auth_scope` - authentication scope

- `restapi.method` - settings for individual methods:
  - `method` - HTTP method (GET, POST, PUT, PATCH, DELETE)
  - `path` - path pattern, can contain parameters with `:` prefix and query parameters
  - `request` - request field for mapping HTTP request body
  - `response` - response field for mapping HTTP response body
  - `success_code` - successful HTTP response code (default 200)

## Examples

### Basic Example

```proto
syntax = "proto3";
package example;

import "restapi/annotations.proto";

service ExampleService {
  option (restapi.service) = {
    base_path : "/api/v1"
  };

  rpc Echo(EchoRequest) returns (EchoResponse) {
    option (restapi.method) = {
      method : "GET"
      path : "/echo/:message"
      response : "*"
    };
  }
}

message EchoRequest {
  string message = 1;
}

message EchoResponse {
  string message = 1;
}
```

Generate the REST API:
```bash
protoc --go_out=. --go-rest_out=. --proto_path=./protos example.proto
```

### Advanced Features

The project supports various REST API usage scenarios, as shown in the `example/proto/example.proto` example:

- Path and query parameters (e.g., `/messages/:id?:page&:per_page`)
- POST, PUT, PATCH requests with body
- Streaming APIs with WebSocket support
- Authentication via Bearer tokens

## Make Commands

The project includes a Makefile with useful commands:

- `make build` - build the project, generate proto files and install the plugin
- `make test-runtime` - run tests and generate code coverage report
- `make example-run` - run the API example
- `make proto` - generate code from annotation files
- `make proto-example` - generate code from examples
- `make lint` - run linter and format code

## Streaming API Support

The project supports streaming APIs via WebSocket for gRPC methods with the `stream` keyword.
Example from `example/proto/example.proto`:

```proto
rpc Echo(stream EchoRequest) returns (stream EchoResponse) {
  option (restapi.method) = {
    method : "GET"
    path : "/echo/:channel"
    request : "*"
    response : "*"
  };
}

rpc Ticker(TickerRequest) returns (stream TickerResponse) {
  option (restapi.method) = {
    method : "GET"
    path : "/ticker/:count"
    request : "*"
    response : "*"
  };
}
```

## Authentication

The project supports authentication via Bearer tokens:

```proto
service ExampleService {
  option (restapi.service) = {
    base_path : "/api/v1/example"
    auth : AUTH_TYPE_BEARER
    auth_scope : "JWT"
  };
  
  // API methods...
}
```

## Contributing

Contributions to the project are welcome! Please create issues or submit pull requests with improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
