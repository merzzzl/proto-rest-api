package main

import (
	"flag"
	"fmt"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/gen"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/openapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const Version = "0.0.0-alpha.0"

var requireUnimplemented *bool

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *showVersion {
		_, _ = fmt.Printf("protoc-gen-go-rest %v\n", Version)

		return
	}

	var flags flag.FlagSet
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			if len(f.Services) == 0 {
				continue
			}

			generateFile(gen, f)
		}

		return nil
	})
}

func generateFile(plug *protogen.Plugin, file *protogen.File) {
	g := plug.NewGeneratedFile(file.GeneratedFilenamePrefix+"_rest.pb.go", file.GoImportPath)

	if err := gen.FileHeader(plug, g, file, Version); err != nil {
		plug.Error(err)
	}

	g.P()

	if err := openapi.NewSwagger(file); err != nil {
		plug.Error(err)
	}

	g.P()

	for _, service := range file.Services {
		gen.WebService(g, service, *requireUnimplemented)
		g.P()

		if err := gen.WebSocket(g, service); err != nil {
			plug.Error(err)
		}

		g.P()
	}

	for _, service := range file.Services {
		if err := gen.RegisterHandler(g, service); err != nil {
			plug.Error(err)
		}

		g.P()
	}

	for _, service := range file.Services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
				if err := gen.StreamHandler(g, service, method); err != nil {
					plug.Error(err)
				}
			} else {
				if err := gen.UnaryHandler(g, file, service, method); err != nil {
					plug.Error(err)
				}
			}

			g.P()
		}
	}

	swaggerJSON := plug.NewGeneratedFile(file.GeneratedFilenamePrefix+"_swagger.json", file.GoImportPath)

	jsonDoc, err := openapi.GetJSON()
	if err != nil {
		plug.Error(err)
	}

	if _, err := swaggerJSON.Write(jsonDoc); err != nil {
		plug.Error(err)
	}

	swaggerDoc := plug.NewGeneratedFile(file.GeneratedFilenamePrefix+"_swagger.go", file.GoImportPath)

	if err := gen.FileHeader(plug, swaggerDoc, file, Version); err != nil {
		plug.Error(err)
	}

	openapi.GenDoc(swaggerDoc, file)
}
