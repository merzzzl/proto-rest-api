package openapi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func ResponseRef(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	successCode := restRule.GetSuccessCode()
	response := openapi3.NewResponse()

	var cName string

	if restRule.GetResponse() != "" {
		if successCode == 0 {
			successCode = 200
		}

		reqFields := strings.Split(restRule.GetResponse(), ".")
		if reqFields[0] == "" {
			reqFields = reqFields[1:]
		}

		if reqFields[0] == "*" {
			reqFields = []string{}
		}

		cName = g.QualifiedGoIdent(method.Output.GoIdent)
		workingMessage := method.Output

		if len(reqFields) > 0 {
			field, goFieldsPath, err := tools.FieldsPath(method.Output, reqFields)
			if err != nil {
				return fmt.Errorf("failed to get response fields for %s: %w", method.GoName, err)
			}

			cName += "." + strings.Join(goFieldsPath, ".")
			workingMessage = field.Message
		}

		schema := Scheme(workingMessage, false, []string{})

		if swagger.Components.Schemas == nil {
			swagger.Components.Schemas = make(openapi3.Schemas)
		}

		swagger.Components.Schemas[cName] = schema.NewRef()

		responseDesc := tools.LineComments(workingMessage.Comments)
		response.Description = &responseDesc
		response.Content = openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Ref: fmt.Sprintf("#/components/schemas/%s", cName),
				},
			},
		}

		if swagger.Components.Responses == nil {
			swagger.Components.Responses = make(openapi3.ResponseBodies)
		}

		if _, exists := swagger.Components.Responses[cName]; !exists {
			swagger.Components.Responses[cName] = &openapi3.ResponseRef{
				Value: response,
			}
		}
	} else {
		if successCode == 0 {
			successCode = 204
		}
	}

	if swagger.Paths == nil {
		swagger.Paths = openapi3.NewPaths()
	}

	path := FormatedPath(restRule.GetPath())
	httpMethod := strings.ToUpper(restRule.GetMethod())

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

		if operation.Responses == nil {
			operation.Responses = openapi3.NewResponses()
		}

		if restRule.GetResponse() == "" {
			operation.Responses.Set(strconv.Itoa(int(successCode)), &openapi3.ResponseRef{
				Value: response,
			})
		} else {
			operation.Responses.Set(strconv.Itoa(int(successCode)), &openapi3.ResponseRef{
				Ref: fmt.Sprintf("#/components/responses/%s", cName),
			})
		}
	}

	return nil
}
