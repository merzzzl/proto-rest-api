package openapi

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var swagger openapi3.T

func NewSwagger(file *protogen.File) error {
	swagger = openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:   fmt.Sprintf("Swagger API (%s)", file.GoPackageName),
			Version: time.Now().Format("20060102"),
		},
		Components: &openapi3.Components{},
		Paths:      &openapi3.Paths{},
		Tags:       openapi3.Tags{},
	}

	for _, s := range file.Services {
		line := tools.LineComments(s.Comments)

		swagger.Tags = append(swagger.Tags, &openapi3.Tag{
			Name:        s.GoName,
			Description: line,
		})
	}

	for _, s := range file.Services {
		serviceOptions, ok := s.Desc.Options().(*descriptorpb.ServiceOptions)
		if !ok {
			return fmt.Errorf("unknown service options in %s", s.GoName)
		}

		extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)

		serviceRule, ok := extValSrv.(*restapi.ServiceRule)
		if !ok {
			return fmt.Errorf("unknown http options in %s", s.GoName)
		}

		if serviceRule.GetAuth() == restapi.AuthType_AUTH_TYPE_NONE {
			continue
		}

		if swagger.Components.SecuritySchemes == nil {
			swagger.Components.SecuritySchemes = make(openapi3.SecuritySchemes)
		}

		swagger.Components.SecuritySchemes[s.GoName+"Auth"] = &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
				Name:         s.GoName + " Authorization",
				Description:  "JWT Authorization in " + s.GoName,
			},
		}
	}

	return nil
}

func AddMethod(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	usedInQuery, err := RequestQueryRef(g, service, method)
	if err != nil {
		return fmt.Errorf("failed to add request query for %s: %w", method.GoName, err)
	}

	usedInPath, err := RequestPathRef(g, service, method)
	if err != nil {
		return fmt.Errorf("failed to add request path for %s: %w", method.GoName, err)
	}

	if err := RequestBodyRef(g, service, method, append(usedInPath, usedInQuery...)); err != nil {
		return fmt.Errorf("failed to add request body for %s: %w", method.GoName, err)
	}

	if err := ResponseRef(g, service, method); err != nil {
		return fmt.Errorf("failed to add response for %s: %w", method.GoName, err)
	}

	return nil
}

func GetJSON() ([]byte, error) {
	return swagger.MarshalJSON()
}

func GenDoc(g *protogen.GeneratedFile, file *protogen.File) {
	embedPackage := protogen.GoImportPath("embed")

	fileName := path.Base(file.GeneratedFilenamePrefix)

	g.P("//go:embed ", fileName+"_swagger.json")
	g.P("var swaggerFile ", embedPackage.Ident("FS"))
	g.P()
	g.P("func Get", strings.Title(fileName), "Swagger() []byte {")
	g.P("js, err := swaggerFile.ReadFile(\"", fileName+"_swagger.json\")")
	g.P("if err != nil {")
	g.P("return []byte(\"{}\")")
	g.P("}")
	g.P()
	g.P("return js")
	g.P("}")
	g.P()
}
