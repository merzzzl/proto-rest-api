package openapi

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func RequestPathRef(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) ([]string, error) {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return nil, fmt.Errorf("unknown http options in %s", method.GoName)
	}

	fields, err := tools.PathFields(method)
	if err != nil {
		return nil, fmt.Errorf("failed to get request fields for %s: %w", method.GoName, err)
	}

	if len(fields) == 0 {
		return []string{}, nil
	}

	if swagger.Paths == nil {
		swagger.Paths = openapi3.NewPaths()
	}

	path, err := FormatedPath(service, method)
	if err != nil {
		return nil, fmt.Errorf("failed to format path for %s: %w", method.GoName, err)
	}

	httpMethod := strings.ToUpper(restRule.GetMethod())

	used := make([]string, 0, len(fields))

	if path != "" && httpMethod != "" {
		pathItem := swagger.Paths.Find(path)
		if pathItem == nil {
			pathItem = &openapi3.PathItem{}
			swagger.Paths.Set(path, pathItem)
		}

		operation := pathItem.GetOperation(httpMethod)
		if operation == nil {
			operation = Operation(service, method)

			pathItem.SetOperation(httpMethod, operation)
		}

		for name, field := range fields {
			schema := Field(field, []string{})

			parameter := &openapi3.Parameter{
				Name:        name,
				In:          "path",
				Description: tools.LineComments(field.Comments),
				Required:    !field.Desc.HasOptionalKeyword(),
				Schema: &openapi3.SchemaRef{
					Value: schema,
				},
			}

			operation.Parameters = append(operation.Parameters, &openapi3.ParameterRef{
				Value: parameter,
			})

			used = append(used, field.GoIdent.GoName)
		}
	}

	return used, nil
}
