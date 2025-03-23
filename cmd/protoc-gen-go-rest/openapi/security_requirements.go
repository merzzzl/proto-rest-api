package openapi

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func SecurityRequirements(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if !ok {
		return fmt.Errorf("unknown service options in %s", service.GoName)
	}

	extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)

	serviceRule, ok := extValSrv.(*restapi.ServiceRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", service.GoName)
	}

	if serviceRule.GetAuth() == restapi.AuthType_AUTH_TYPE_NONE {
		return nil
	}

	nameSecurity := service.GoName + "Auth"

	if serviceRule.GetAuthScope() != "" {
		nameSecurity = serviceRule.GetAuthScope()
	}

	if swagger.Paths == nil {
		swagger.Paths = openapi3.NewPaths()
	}

	path, err := FormatedPath(service, method)
	if err != nil {
		return fmt.Errorf("failed to format path for %s: %w", method.GoName, err)
	}

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

		if operation.Security == nil {
			operation.Security = openapi3.NewSecurityRequirements()
		}

		operation.Security.With(map[string][]string{
			nameSecurity: {},
		})
	}

	return nil
}
