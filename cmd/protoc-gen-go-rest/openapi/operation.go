package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
)

func Operation(service *protogen.Service, method *protogen.Method) *openapi3.Operation {
	operation := openapi3.NewOperation()
	operation.Tags = []string{service.GoName}
	operation.Description = tools.LineComment(method.Comments.Trailing)
	operation.Summary = tools.LineComment(method.Comments.Leading)

	return operation
}
