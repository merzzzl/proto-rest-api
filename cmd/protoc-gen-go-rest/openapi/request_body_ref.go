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

func RequestBodyRef(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method, usedFileds []string) error {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	if restRule.GetRequest() == "" {
		return nil
	}

	reqFields := strings.Split(restRule.GetRequest(), ".")
	if reqFields[0] == "" {
		reqFields = reqFields[1:]
	}

	if reqFields[0] == "*" {
		reqFields = []string{}
	}

	cName := g.QualifiedGoIdent(method.Input.GoIdent)
	workingMessage := method.Input
	isListMessage := false

	if len(reqFields) > 0 {
		field, goFieldsPath, err := tools.FieldsPath(method.Input, reqFields)
		if err != nil {
			return fmt.Errorf("failed to get request fields for %s: %w", method.GoName, err)
		}

		cName += "." + strings.Join(goFieldsPath, ".")
		workingMessage = field.Message

		if field.Desc.IsList() {
			isListMessage = true
		}
	}

	schema := Scheme(workingMessage, isListMessage, usedFileds)

	if swagger.Components.Schemas == nil {
		swagger.Components.Schemas = make(openapi3.Schemas)
	}

	swagger.Components.Schemas[cName] = schema.NewRef()

	requestDesc := tools.LineComments(workingMessage.Comments)
	requestBody := &openapi3.RequestBody{
		Description: requestDesc,
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Ref: fmt.Sprintf("#/components/schemas/%s", cName),
				},
			},
		},
		Required: true,
	}

	if swagger.Components.RequestBodies == nil {
		swagger.Components.RequestBodies = make(openapi3.RequestBodies)
	}

	if _, exists := swagger.Components.RequestBodies[cName]; !exists {
		swagger.Components.RequestBodies[cName] = &openapi3.RequestBodyRef{
			Value: requestBody,
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

		operation.RequestBody = &openapi3.RequestBodyRef{
			Ref: fmt.Sprintf("#/components/requestBodies/%s", cName),
		}
	}

	return nil
}
