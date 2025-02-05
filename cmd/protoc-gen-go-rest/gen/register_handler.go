package gen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func RegisterHandler(g *protogen.GeneratedFile, service *protogen.Service) error {
	g.P("// Register", service.GoName, "Handler registers the http handlers for service ", service.GoName, " to \"mux\".")
	g.P("func Register", service.GoName, "Handler(mux ", runtimePackage.Ident("ServeMuxer"), ", server ", service.GoName, "WebServer, interceptors ...", runtimePackage.Ident("Interceptor"), ") {")
	g.P("router := ", runtimePackage.Ident("NewRouter"), "()")
	g.P()

	serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if !ok {
		return fmt.Errorf("unknown service options in %s", service.GoName)
	}

	extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)

	serviceRule, ok := extValSrv.(*restapi.ServiceRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", service.GoName)
	}

	basePath := serviceRule.GetBasePath()

	if basePath == "/" {
		return fmt.Errorf("base path %s is not allowed", basePath)
	}

	basePath = strings.TrimSuffix(basePath, "/")

	if !strings.HasPrefix(basePath, "/") {
		basePath = "/" + basePath
	}

	paths, err := tools.MethodsPaths(service)
	if err != nil {
		return fmt.Errorf("failed to get methods paths for %s: %w", service.GoName, err)
	}

	for method, restRule := range paths {
		subPath := restRule.GetPath()

		if sep := strings.LastIndex(subPath, "?"); sep != -1 {
			subPath = subPath[:sep]
		}

		if subPath == "/" || subPath == "" {
			return errors.New("empty path is not allowed")
		}

		if !strings.HasPrefix(subPath, "/") {
			subPath = "/" + subPath
		}

		subPath = strings.TrimSuffix(subPath, "/")

		if (method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer()) && restRule.GetMethod() != "GET" {
			return fmt.Errorf("streaming methods are not allowed for %s", restRule.GetMethod())
		}

		g.P("router.Handle(\"", strings.ToUpper(restRule.GetMethod()), "\", \"", basePath+subPath, "\", func(w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", p ", runtimePackage.Ident("Params"), ") {")
		g.P("handler", service.GoName, "WebServer", method.GoName, "(server, w, r, p, interceptors)")
		g.P("})")
		g.P()
	}

	g.P("mux.Handle(\"", basePath, "/\", router)")
	g.P("}")

	return nil
}
