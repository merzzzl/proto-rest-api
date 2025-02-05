package gen

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func UnimplementedStruct(g *protogen.GeneratedFile, service *protogen.Service, requireUnimplemented bool) {
	mustOrShould := "must"
	if !requireUnimplemented {
		mustOrShould = "should"
	}

	g.P("// Unimplemented", service.GoName, "WebServer ", mustOrShould, " be embedded to have forward compatible implementations.")
	g.P("type Unimplemented", service.GoName, "WebServer struct {}")
	g.P()

	for _, method := range service.Methods {
		nilArg := ""

		if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
			nilArg = "nil,"
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

		g.P("func (Unimplemented", service.GoName, "WebServer) ", method.GoName, "(", strings.Join(reqArgs, ", "), ") ", ret, "{")
		g.P("return ", nilArg, runtimePackage.Ident("Errorf"), "(", httpPackage.Ident("StatusNotImplemented"), `, "method not implemented")`)
		g.P("}")
		g.P()
	}

	if requireUnimplemented {
		g.P("func (Unimplemented", service.GoName, "WebServer) mustEmbedUnimplemented", service.GoName, "WebServer() {}")
	}
}
