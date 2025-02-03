package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var version = "0.0.0-alpha.0"

var requireUnimplemented *bool

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *showVersion {
		_, _ = fmt.Printf("protoc-gen-go-rest %v\n", version)

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

const (
	errorsPackage  = protogen.GoImportPath("errors")
	contextPackage = protogen.GoImportPath("context")
	ioPackage      = protogen.GoImportPath("io")
	httpPackage    = protogen.GoImportPath("net/http")
	jsonPackage    = protogen.GoImportPath("encoding/json")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")
	runtimePackage = protogen.GoImportPath("github.com/merzzzl/proto-rest-api/runtime")
)

const (
	serviceSufix = "WebServer"
)

const deprecationComment = "// Deprecated: Do not use."

var genQueue = make([]func(), 0)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_rest.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	genHeader(gen, g, file)

	for _, service := range file.Services {
		genServiceStructs(g, service)
	}

	for _, service := range file.Services {
		genRegisterHandler(g, service)
	}

	for _, service := range file.Services {
		for _, method := range service.Methods {
			genHandler(g, service, method)
		}
	}

	for _, f := range genQueue {
		f()
	}

	return g
}

func exitWithError(message string) {
	_, _ = fmt.Fprintln(os.Stderr, message)

	os.Exit(1)
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
