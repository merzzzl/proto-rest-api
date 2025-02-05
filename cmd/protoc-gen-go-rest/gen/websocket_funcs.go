package gen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
)

func WebSocketFuncs(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	qfileds, err := tools.QueryFields(method)
	if err != nil {
		return fmt.Errorf("failed to get query fields for %s: %w", method.GoName, err)
	}

	pfileds, err := tools.PathFields(method)
	if err != nil {
		return fmt.Errorf("failed to get path fields for %s: %w", method.GoName, err)
	}

	if method.Desc.IsStreamingServer() {
		g.P("func (x *x_", service.GoName, method.GoName, "WebSocket) Send(m *", method.Output.GoIdent, ") error {")
		g.P("return x.ServerStream.SendMsg(m)")
		g.P("}")
		g.P()
	} else {
		g.P("func (x *x_", service.GoName, method.GoName, "WebSocket) SendAndClose(m *", method.Output.GoIdent, ") error {")
		g.P("return x.ServerStream.SendMsg(m)")
		g.P("}")
		g.P()
	}

	g.P("func (x *x_", service.GoName, method.GoName, "WebSocket) Recv() (*", method.Input.GoIdent, ", error) {")
	g.P("m := new(", method.Input.GoIdent, ")")
	g.P()
	g.P("if err := x.ServerStream.RecvMsg(m); err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P()

	for _, filed := range qfileds {
		g.P("m.", filed.GoName, " = ", "x.", filed.GoName)
	}

	for _, filed := range pfileds {
		g.P("m.", filed.GoName, " = ", "x.", filed.GoName)
	}

	g.P()
	g.P("return m, nil")
	g.P("}")

	return nil
}
