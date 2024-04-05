package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/merzzzl/proto-rest-api/gen/go/rest/api"
)

type (
	methodMap map[string]methodMap
	baseptMap map[string]string
)

var (
	methods = make(methodMap)
	basepts = make(baseptMap)
)

func genRegisterHandler(g *protogen.GeneratedFile, service *protogen.Service) {
	serverType := service.GoName + serviceSufix

	g.P("// Register", service.GoName, "Handler registers the http handlers for service ", service.GoName, " to \"mux\".")
	g.P("func Register", service.GoName, "Handler(mux ", runtimePackage.Ident("ServeMuxer"), ", server ", serverType, ", interceptors ...", runtimePackage.Ident("Interceptor"), ") {")
	g.P("router := ", runtimePackage.Ident("NewRouter"), "()")
	g.P()

	serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if !ok {
		exitWithError(fmt.Sprintf("unknown service options in %s", service.GoName))
	}

	extValSrv := proto.GetExtension(serviceOptions, api.E_Service)

	serviceRule, ok := extValSrv.(*api.ServiceRule)
	if !ok {
		exitWithError(fmt.Sprintf("unknown http options in %s", service.GoName))
	}

	basePath := serviceRule.GetBasePath()

	if basePath == "/" {
		exitWithError(fmt.Sprintf("base path %s is not allowed", basePath))
	}

	if basePath != "" {
		if basePath[0] != '/' {
			basePath = "/" + basePath
		}

		if basePath[len(basePath)-1] == '/' {
			basePath = basePath[:len(basePath)-1]
		}
	}

	for _, method := range service.Methods {
		methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
		if !ok {
			exitWithError(fmt.Sprintf("unknown method options in %s", method.GoName))
		}

		extVal := proto.GetExtension(methodOptions, api.E_Method)

		restRule, ok := extVal.(*api.MethodRule)
		if !ok {
			exitWithError(fmt.Sprintf("unknown http options in %s", method.GoName))
		}

		subPath := restRule.GetPath()

		if sep := strings.LastIndex(subPath, "?"); sep != -1 {
			subPath = subPath[:sep]
		}

		if subPath == "/" {
			exitWithError(fmt.Sprintf("path %s is not allowed", subPath))
		}

		if subPath == "" {
			exitWithError("empty path is not allowed")
		}

		if subPath[0] != '/' {
			subPath = "/" + subPath
		}

		if subPath[len(subPath)-1] == '/' {
			subPath = subPath[:len(subPath)-1]
		}

		if (method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer()) && restRule.GetMethod() != "GET" {
			exitWithError(fmt.Sprintf("streaming methods are not allowed for %s", restRule.GetMethod()))
		}

		methodExistsError(service.GoName, strings.ToUpper(restRule.GetMethod()), basePath, subPath)

		g.P("router.Handle(\"", strings.ToUpper(restRule.GetMethod()), "\", \"", basePath+subPath, "\", func(w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", p ", runtimePackage.Ident("Params"), ") {")
		g.P("handler", serverType, method.GoName, "(server, w, r, p, interceptors)")
		g.P("})")
		g.P()
	}

	g.P("mux.Handle(\"", basePath, "/\", router)")
	g.P("}")
	g.P()
}

func methodExistsError(service, method, basePath, url string) {
	segs := strings.Split(url, "/")
	segs = segs[1:]

	if basepts[basePath] != "" && basepts[basePath] != service {
		exitWithError(fmt.Sprintf("base path %s already exists in %s", basePath, basepts[basePath]))
	}

	basepts[basePath] = service

	if methods[service] == nil {
		methods[service] = make(methodMap)
	}

	tmp := methods[service]

	for i := 0; i < len(segs); i++ {
		if segs[i][0] == ':' {
			segs[i] = ":"
		}

		if next, ok := tmp[segs[i]]; ok {
			tmp = next
		} else {
			if _, ok := tmp[":"]; ok {
				if segs[i] != ":" {
					exitWithError(fmt.Sprintf("dunamically and statically defined methods %s in %s", url, service))
				}
			}

			tmp[segs[i]] = make(methodMap)
			tmp = tmp[segs[i]]
		}
	}

	if _, ok := tmp[method]; ok {
		exitWithError(fmt.Sprintf("method %s already exists in %s", url, service))
	}

	tmp[method] = make(methodMap)
}
