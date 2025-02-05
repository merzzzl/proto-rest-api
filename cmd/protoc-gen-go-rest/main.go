package main

import (
	"flag"
	"fmt"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/gen"
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

func generateFile(plug *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_rest.pb.go"
	g := plug.NewGeneratedFile(filename, file.GoImportPath)

	gen.FileHeader(plug, g, file, Version)
	g.P()

	for _, service := range file.Services {
		gen.WebService(g, service, *requireUnimplemented)
		g.P()
		gen.WebSocket(g, service)
		g.P()
	}

	for _, service := range file.Services {
		gen.RegisterHandler(g, service)
		g.P()
	}

	for _, service := range file.Services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
				gen.StreamHandler(g, service, method)
			} else {
				gen.UnaryHandler(g, service, method)
			}

			g.P()
		}
	}

	return g
}
