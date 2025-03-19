package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/restapi"
)

var version = "0.0.0-alpha.0"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("protoc-gen-go-rest %v\n", version)
		return
	}

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		files := make(map[string]*protogen.File)

		for _, f := range gen.Files {
			if !f.Generate || len(f.Services) == 0 {
				continue
			}
			files[string(f.GoPackageName)] = f
		}

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			lf, ok := files[string(f.GoPackageName)]
			if !ok {
				files[string(f.GoPackageName)] = f
				continue
			}
			if lf.GeneratedFilenamePrefix != f.GeneratedFilenamePrefix {
				lf.Services = append(lf.Services, f.Services...)
				lf.Messages = append(lf.Messages, f.Messages...)
			}
		}

		for _, lf := range files {
			if err := generateFile(gen, lf); err != nil {
				return err
			}
		}
		return nil
	})
}

const (
	runtimePackage = protogen.GoImportPath("github.com/merzzzl/proto-rest-api/runtime")
	swaggerPackage = protogen.GoImportPath("github.com/merzzzl/proto-rest-api/swagger")
	httpPackage    = protogen.GoImportPath("net/http")
	stringsPackage = protogen.GoImportPath("strings")
)

func generateFile(gen *protogen.Plugin, file *protogen.File) error {
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_oapi.pb.go", file.GoImportPath)

	g.P("// Code generated by protoc-gen-go-oapi. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-oapi v", version)
	g.P("// - protoc             ", protocVersion(gen))

	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}

	g.P()

	g.P("package ", file.GoPackageName)
	g.P()

	// Add blank embed import
	g.P("import _ \"embed\"")
	g.P()

	doc := &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:   toTitleString(string(file.GoPackageName)),
			Version: time.Now().Format(time.DateOnly),
		},
		Paths:      openapi3.NewPaths(),
		Components: &openapi3.Components{Schemas: openapi3.Schemas{}},
	}

	for _, service := range file.Services {
		doc.Tags = append(doc.Tags, &openapi3.Tag{
			Name:        service.GoName,
			Description: formatComment(service.Comments.Leading.String()),
		})

		serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
		if !ok {
			return fmt.Errorf("unknown service options in %s", service.GoName)
		}

		extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)
		serviceRule, ok := extValSrv.(*restapi.ServiceRule)
		if !ok {
			return fmt.Errorf("unknown http options in %s", service.GoName)
		}

		basePath := normalizeBasePath(serviceRule.GetBasePath())
		if basePath == "/" {
			return fmt.Errorf("base path %s is not allowed", basePath)
		}

		for _, method := range service.Methods {
			if err := processMethod(doc, service, method, basePath, serviceRule.GetAuth()); err != nil {
				return err
			}
		}
	}

	jsonBytes, err := doc.MarshalJSON()
	if err != nil {
		return err
	}

	gspec := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_oapi.json", file.GoImportPath)
	gspec.P(string(jsonBytes))
	swaggerPath := strings.Split(file.GeneratedFilenamePrefix, "/")

	g.P("//go:embed ", swaggerPath[len(swaggerPath)-1]+"_oapi.json")
	g.P("var swaggerJSON []byte")
	g.P()

	g.P("// RegisterSwaggerUIHandler registers swagger ui handler.")
	g.P("func RegisterSwaggerUIHandler(mux ", runtimePackage.Ident("ServeMuxer"), ", path string) error {")
	g.P("if !", stringsPackage.Ident("HasPrefix"), "(path, \"/\") {")
	g.P("path = \"/\" + path")
	g.P("}")
	g.P()
	g.P("if !", stringsPackage.Ident("HasSuffix"), "(path, \"/\") {")
	g.P("path += \"/\"")
	g.P("}")
	g.P()
	g.P("fs, err := ", swaggerPackage.Ident("GetSwaggerUI"), "(swaggerJSON, path)")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P()
	g.P("mux.Handle(path, ", httpPackage.Ident("FileServer"), "(fs))")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("// RegisterReDocUIHandler registers redoc ui handler.")
	g.P("func RegisterReDocUIHandler(mux ", runtimePackage.Ident("ServeMuxer"), ", path string) error {")
	g.P("if !", stringsPackage.Ident("HasPrefix"), "(path, \"/\") {")
	g.P("path = \"/\" + path")
	g.P("}")
	g.P()
	g.P("if !", stringsPackage.Ident("HasSuffix"), "(path, \"/\") {")
	g.P("path += \"/\"")
	g.P("}")
	g.P()
	g.P("fs, err := ", swaggerPackage.Ident("GetReDocUI"), "(swaggerJSON, path)")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P()
	g.P("mux.Handle(path, ", httpPackage.Ident("FileServer"), "(fs))")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	return nil
}

func addAuthorization(doc *openapi3.T, operation *openapi3.Operation, authType restapi.AuthType) {
	if doc.Components.SecuritySchemes == nil {
		doc.Components.SecuritySchemes = openapi3.SecuritySchemes{}
	}

	if _, exists := doc.Components.SecuritySchemes["BearerToken"]; !exists {
		doc.Components.SecuritySchemes["BearerToken"] = &openapi3.SecuritySchemeRef{
			Value: openapi3.NewSecurityScheme().
				WithType("http").
				WithScheme("bearer").
				WithBearerFormat("JWT"),
		}
	}

	if authType == restapi.AuthType_AUTH_TYPE_BEARER {
		operation.Security = &openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{"BearerToken": []string{}},
		}
	}
}

func processMethod(doc *openapi3.T, service *protogen.Service, method *protogen.Method, basePath string, authType restapi.AuthType) error {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)
	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	path := normalizePath(basePath + restRule.GetPath())
	if path == "/" || path == "" {
		return fmt.Errorf("invalid path %s for method %s", path, method.GoName)
	}

	oapiPath, pathParams, queryParams := getPathForOAPI(path)
	operation := openapi3.NewOperation()
	operation.Tags = []string{service.GoName}
	operation.Summary = formatComment(method.Comments.Leading.String())
	operation.OperationID = service.GoName + "_" + method.GoName

	addAuthorization(doc, operation, authType)

	for _, param := range pathParams {
		field := findFieldByName(method.Input, param)
		var schema *openapi3.Schema
		if field == nil {
			// If the field isn’t found, assume it’s a string parameter
			schema = &openapi3.Schema{Type: &openapi3.Types{"string"}}
		} else {
			schema = getSchemaForField(field)
			if field.Desc.Kind() == protoreflect.MessageKind {
				msgSchema := genMessageSchema(field.Message)
				doc.Components.Schemas[field.Message.GoIdent.GoName] = msgSchema.NewRef()
				schema = openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, nil).Value
			}
		}
		operation.AddParameter(openapi3.NewPathParameter(param).
			WithRequired(true).
			WithDescription(getDescriptionByName(method.Input, param)).
			WithSchema(schema))
	}

	for _, param := range queryParams {
		field := findFieldByName(method.Input, param)
		var schema *openapi3.Schema
		if field == nil {
			// If the field isn’t found, assume it’s a string parameter
			schema = &openapi3.Schema{Type: &openapi3.Types{"string"}}
		} else {
			schema = getSchemaForField(field)
			if field.Desc.Kind() == protoreflect.MessageKind {
				msgSchema := genMessageSchema(field.Message)
				doc.Components.Schemas[field.Message.GoIdent.GoName] = msgSchema.NewRef()
				schema = openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, nil).Value
			}
		}
		operation.AddParameter(openapi3.NewQueryParameter(param).
			WithRequired(true).
			WithDescription(getDescriptionByName(method.Input, param)).
			WithSchema(schema))
	}

	// Rest of the function remains unchanged...
	if req := restRule.GetRequest(); req != "" {
		fields := strings.Split(req, ".")
		if fields[0] == "" {
			fields = fields[1:]
		}

		var schemaRef *openapi3.SchemaRef
		if fields[0] == "*" {
			schema := genMessageSchema(method.Input)
			doc.Components.Schemas[method.Input.GoIdent.GoName] = schema.NewRef()
			schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+method.Input.GoIdent.GoName, schema)
		} else {
			field := fieldByPath(method.Input, fields)
			schema := getSchemaForField(field)
			if field.Desc.Kind() == protoreflect.MessageKind {
				msgSchema := genMessageSchema(field.Message)
				doc.Components.Schemas[field.Message.GoIdent.GoName] = msgSchema.NewRef()
				schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, msgSchema)
			} else {
				doc.Components.Schemas[field.GoIdent.GoName] = schema.NewRef()
				schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+field.GoIdent.GoName, schema)
			}
		}

		operation.RequestBody = &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("A JSON object containing request parameters.").
				WithRequired(true).
				WithContent(openapi3.NewContentWithJSONSchemaRef(schemaRef)),
		}
	}

	successCode := restRule.GetSuccessCode()
	if successCode == 0 {
		successCode = 200
	}

	response := openapi3.NewResponse().WithDescription("A successful response.")
	if resp := restRule.GetResponse(); resp != "" {
		fields := strings.Split(resp, ".")
		if fields[0] == "" {
			fields = fields[1:]
		}

		var schemaRef *openapi3.SchemaRef
		if fields[0] == "*" {
			schema := genMessageSchema(method.Output)
			doc.Components.Schemas[method.Output.GoIdent.GoName] = schema.NewRef()
			schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+method.Output.GoIdent.GoName, schema)
		} else {
			field := fieldByPath(method.Output, fields)
			schema := getSchemaForField(field)
			if field.Desc.Kind() == protoreflect.MessageKind {
				msgSchema := genMessageSchema(field.Message)
				doc.Components.Schemas[field.Message.GoIdent.GoName] = msgSchema.NewRef()
				schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, msgSchema)
			} else {
				doc.Components.Schemas[field.GoIdent.GoName] = schema.NewRef()
				schemaRef = openapi3.NewSchemaRef("#/components/schemas/"+field.GoIdent.GoName, schema)
			}
		}

		response.WithContent(openapi3.NewContentWithJSONSchemaRef(schemaRef))
	}

	pathItem := doc.Paths.Find(oapiPath)
	if pathItem == nil {
		pathItem = &openapi3.PathItem{}
		doc.Paths.Set(oapiPath, pathItem)
	}

	switch strings.ToLower(restRule.GetMethod()) {
	case "get":
		pathItem.Get = operation
	case "post":
		pathItem.Post = operation
	case "put":
		pathItem.Put = operation
	case "delete":
		pathItem.Delete = operation
	case "patch":
		pathItem.Patch = operation
	default:
		return fmt.Errorf("unsupported HTTP method: %s", restRule.GetMethod())
	}

	return nil
}

func normalizeBasePath(path string) string {
	if path == "" {
		return ""
	}
	if path[0] != '/' {
		path = "/" + path
	}
	return strings.TrimSuffix(path, "/")
}

func normalizePath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		path = "/" + path
	}
	return strings.TrimSuffix(path, "/") + "/"
}

func getPathForOAPI(s string) (string, []string, []string) {
	pathParams := []string{}
	queryParams := []string{}

	parts := strings.Split(s, "?")
	segs := strings.Split(parts[0], "/")
	for i, seg := range segs {
		if strings.HasPrefix(seg, ":") {
			param := seg[1:]
			pathParams = append(pathParams, param)
			segs[i] = "{" + param + "}"
		}
	}

	if len(parts) > 1 {
		for _, param := range strings.Split(parts[1], "&") {
			if strings.HasPrefix(param, ":") {
				queryParams = append(queryParams, param[1:])
			}
		}
	}

	return strings.Join(segs, "/"), pathParams, queryParams
}

func fieldByPath(msg *protogen.Message, fields []string) *protogen.Field {
	for _, field := range msg.Fields {
		if field.Desc.TextName() == fields[0] {
			if len(fields) == 1 {
				return field
			}
			return fieldByPath(field.Message, fields[1:])
		}
	}
	panic(fmt.Sprintf("unknown field %s in %s", fields[0], msg.GoIdent))
}

func genMessageSchema(msg *protogen.Message) *openapi3.Schema {
	schema := openapi3.NewObjectSchema()
	for _, field := range msg.Fields {
		if field.Message != nil {
			schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, nil)
			if field.Desc.IsList() {
				arraySchema := openapi3.NewArraySchema()
				arraySchema.Items = schemaRef
				schema.Properties[field.Desc.JSONName()] = openapi3.NewSchemaRef("", arraySchema)
			} else {
				schema.Properties[field.Desc.JSONName()] = schemaRef
			}
			continue
		}
		schema.Properties[field.Desc.JSONName()] = getSchemaForField(field).NewRef()
	}
	return schema
}

func getSchemaForField(field *protogen.Field) *openapi3.Schema {
	var schema *openapi3.Schema
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"boolean"}}
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int32"}
	case protoreflect.Uint32Kind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "uint32", Min: ptrFloat64(0)}
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}
	case protoreflect.Uint64Kind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "uint64", Min: ptrFloat64(0)}
	case protoreflect.FloatKind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"number"}, Format: "float"}
	case protoreflect.DoubleKind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"number"}, Format: "double"}
	case protoreflect.StringKind, protoreflect.BytesKind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"string"}}
	case protoreflect.EnumKind:
		schema = &openapi3.Schema{Type: &openapi3.Types{"string"}, Enum: enumValues(field.Enum)}
	case protoreflect.MessageKind:
		// Create a reference to the nested message schema
		schema = openapi3.NewObjectSchema()
		schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+field.Message.GoIdent.GoName, nil)
		return schemaRef.Value // Return the schema; actual definition is handled elsewhere
	default:
		panic(fmt.Sprintf("unsupported field type %s", field.Desc.Kind()))
	}

	if field.Desc.IsList() {
		schema = &openapi3.Schema{Type: &openapi3.Types{"array"}, Items: &openapi3.SchemaRef{Value: schema}}
	}

	return schema
}

func ptrFloat64(v float64) *float64 {
	return &v
}

func findFieldByName(msg *protogen.Message, name string) *protogen.Field {
	for _, field := range msg.Fields {
		if field.Desc.TextName() == name {
			return field
		}
	}
	return nil
}

func getDescriptionByName(msg *protogen.Message, name string) string {
	if field := findFieldByName(msg, name); field != nil {
		return formatComment(field.Comments.Trailing.String())
	}
	return ""
}

func enumValues(enum *protogen.Enum) []interface{} {
	values := make([]interface{}, len(enum.Values))
	for i, v := range enum.Values {
		values[i] = string(v.Desc.Name())
	}
	return values
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	suffix := ""
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func toTitleString(s string) string {
	if s == "" {
		return "Swagger"
	}
	seps := strings.Split(s, "_")
	for i, sep := range seps {
		seps[i] = string(unicode.ToUpper(rune(sep[0]))) + sep[1:]
	}
	return strings.Join(seps, "")
}

func formatComment(s string) string {
	s = strings.TrimPrefix(s, "//")
	s = strings.TrimSuffix(s, "\n")
	return strings.TrimSpace(s)
}
