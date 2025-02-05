package gen

import "google.golang.org/protobuf/compiler/protogen"

func WebSocketInreface(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	g.P("type ", service.GoName, method.GoName, "WebSocket interface {")

	if method.Desc.IsStreamingServer() {
		g.P("Send(*", method.Output.GoIdent, ") error")
	} else {
		g.P("SendAndClose(*", method.Output.GoIdent, ") error")
	}

	if method.Desc.IsStreamingClient() {
		g.P("Recv() (*", method.Input.GoIdent, ", error)")
	}

	g.P(protogen.GoImportPath("google.golang.org/grpc").Ident("ServerStream"))
	g.P("}")
}
