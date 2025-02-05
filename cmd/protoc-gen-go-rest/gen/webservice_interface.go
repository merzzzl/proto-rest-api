package gen

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func WebServiceInterface(g *protogen.GeneratedFile, service *protogen.Service, requireUnimplemented bool) {
	mustOrShould := "must"
	if !requireUnimplemented {
		mustOrShould = "should"
	}

	g.P("// ", service.GoName, "WebServer is the server API for ", service.GoName, " service.")
	g.P("// All implementations ", mustOrShould, " embed Unimplemented", service.GoName, "WebServer for forward compatibility.")

	if serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions); ok && serviceOptions.GetDeprecated() {
		g.P("//")
		g.P("// Deprecated: Do not use.")
	}

	g.P("type ", service.GoName, "WebServer interface {")

	for _, method := range service.Methods {
		if methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions); ok && methodOptions.GetDeprecated() {
			g.P("// Deprecated: Do not use.")
		}

		reqArgs := []string{}
		ret := "error"

		if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
			reqArgs = append(reqArgs, g.QualifiedGoIdent(contextPackage.Ident("Context")))
			ret = "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
		}

		if !method.Desc.IsStreamingClient() {
			reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(method.Input.GoIdent))
		}

		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			reqArgs = append(reqArgs, method.Parent.GoName+method.GoName+"WebSocket")
		}

		g.P(method.Comments.Leading, method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret)
	}

	if requireUnimplemented {
		g.P("mustEmbedUnimplemented", service.GoName, "WebServer()")
	}

	g.P("}")
}
