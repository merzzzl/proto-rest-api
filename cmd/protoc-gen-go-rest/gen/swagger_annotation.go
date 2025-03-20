package gen

import (
	"fmt"
	"strings"
	"time"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func swaggerAnnotation(g *protogen.GeneratedFile, file *protogen.File, service *protogen.Service, method *protogen.Method) error {
	serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if !ok {
		return fmt.Errorf("unknown service options in %s", service.GoName)
	}

	extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)

	serviceRule, ok := extValSrv.(*restapi.ServiceRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", service.GoName)
	}

	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	var reqBody *string

	if restRule.GetRequest() != "" {
		reqFileds := strings.Split(restRule.GetRequest(), ".")

		if reqFileds[0] == "" {
			reqFileds = reqFileds[1:]
		}

		if reqFileds[0] == "*" {
			reqFileds = []string{}
		}

		body := method.Input.GoIdent.GoName

		if len(reqFileds) > 0 {
			_, goFieldsPath, err := tools.FieldsPath(method.Input, reqFileds)
			if err != nil {
				return fmt.Errorf("failed to get request fields for %s: %w", method.GoName, err)
			}

			body += "." + strings.Join(goFieldsPath, ".")
		}

		reqBody = &body
	}

	var rspBody *string

	if restRule.GetResponse() != "" {
		reqFileds := strings.Split(restRule.GetResponse(), ".")

		if reqFileds[0] == "" {
			reqFileds = reqFileds[1:]
		}

		if reqFileds[0] == "*" {
			reqFileds = []string{}
		}

		body := method.Output.GoIdent.GoName

		if len(reqFileds) > 0 {
			_, goFieldsPath, err := tools.FieldsPath(method.Output, reqFileds)
			if err != nil {
				return fmt.Errorf("failed to get response fields for %s: %w", method.GoName, err)
			}

			body += "." + strings.Join(goFieldsPath, ".")
		}

		rspBody = &body
	}

	pathFields, err := tools.PathFields(method)
	if err != nil {
		return fmt.Errorf("failed to get path fields for %s: %w", method.GoName, err)
	}

	queryFields, err := tools.QueryFields(method)
	if err != nil {
		return fmt.Errorf("failed to get query fields for %s: %w", method.GoName, err)
	}

	var successCode int32

	if restRule.GetSuccessCode() != 0 {
		successCode = restRule.GetSuccessCode()
	} else {
		successCode = 200
	}

	g.P("// ", method.GoName)
	g.P("// @Tags ", service.GoName)

	if serviceRule.GetAuth() != restapi.AuthType_AUTH_TYPE_NONE {
		g.P("// @Security Auth"+service.GoName)
	}

	g.P("// @Summary ", tools.LineComment(method.Comments.Leading))
	g.P("// @Description ", tools.LineComment(method.Comments.Trailing))

	if reqBody != nil {
		comment := tools.LineComments(method.Input.Comments)
		if comment == "" {
			comment = "body of the request"
		}

		g.P("// @Accept json")
		g.P("// @Param request body ", file.GoPackageName, ".", *reqBody, " true \"", comment, "\"")
	}

	for param, field := range pathFields {
		mandatory := "true"

		if field.Desc.HasOptionalKeyword() {
			mandatory = "false"
		}

		t, err := tools.TypeConverter(field)
		if err != nil {
			return fmt.Errorf("failed to convert type for %s: %w", field.GoIdent, err)
		}

		g.P("// @Param ", param, " path ", t, " ", mandatory, " \"", tools.LineComments(field.Comments), "\"")
	}

	for param, field := range queryFields {
		mandatory := "true"

		if field.Desc.HasOptionalKeyword() {
			mandatory = "false"
		}

		t, err := tools.TypeConverter(field)
		if err != nil {
			return fmt.Errorf("failed to convert type for %s: %w", field.GoIdent, err)
		}

		if field.Desc.IsList() {
			t = "[]" + t
		}

		g.P("// @Param ", param, " query ", t, " ", mandatory, " \"", tools.LineComments(field.Comments), "\"")
	}

	if rspBody != nil {
		g.P("// @Produce json")
		g.P("// @Success ", successCode, " {object} ", file.GoPackageName, ".", *rspBody)
	}

	pathMethodSegs := strings.Split(strings.Split(restRule.GetPath(), "?")[0], "/")

	for i, s := range pathMethodSegs {
		if strings.HasPrefix(s, ":") {
			pathMethodSegs[i] = "{" + s[1:] + "}"
		}
	}

	pathService := strings.TrimSuffix(serviceRule.GetBasePath(), "/")
	pathMethod := strings.TrimPrefix(strings.Join(pathMethodSegs, "/"), "/")
	pathMethod = strings.TrimSuffix(pathMethod, "/")
	pathMethod = pathService + "/" + pathMethod

	g.P("// @Router ", pathMethod, " [", strings.ToUpper(restRule.GetMethod()), "]")
	g.P("//")

	return nil
}

func SwaggerFile(g *protogen.GeneratedFile, file *protogen.File) error {
	g.P("// @title Swagger API (", file.GoPackageName, ")")
	g.P("// @version ", time.Now().Format("20060102"))

	for _, s := range file.Services {
		line := tools.LineComments(s.Comments)

		g.P("// @tag.name ", s.GoName)

		if line != "" {
			g.P("// @tag.description ", line)
		}
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

		g.P("// @securityDefinitions.apikey Auth"+s.GoName)
		g.P("// @in header")
		g.P("// @name Authorization")
		g.P("// @description Bearer token")
	}

	return nil
}
