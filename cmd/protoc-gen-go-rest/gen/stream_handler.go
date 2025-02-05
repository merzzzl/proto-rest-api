package gen

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func StreamHandler(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	if strings.Contains(restRule.GetPath(), "/:") {
		g.P("func handler", service.GoName, "WebServer", method.GoName, "(server ", service.GoName, "WebServer, w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", p ", runtimePackage.Ident("Params"), ", il []", runtimePackage.Ident("Interceptor"), ") {")
	} else {
		g.P("func handler", service.GoName, "WebServer", method.GoName, "(server ", service.GoName, "WebServer, w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", _ ", runtimePackage.Ident("Params"), ", il []", runtimePackage.Ident("Interceptor"), ") {")
	}

	g.P("ctx, cancel := ", contextPackage.Ident("WithCancel"), "(r.Context())")
	g.P("defer cancel()")
	g.P()

	g.P("ctx, err := ", runtimePackage.Ident("ApplyInterceptors"), "(ctx, r, il...)")
	g.P("if err != nil {")
	g.P("errstatus := ", runtimePackage.Ident("GetHTTPStatusFromError"), "(err)")
	g.P()
	g.P("w.WriteHeader(errstatus.Code())")
	g.P("if _, err := w.Write([]byte(errstatus.Message())); err != nil {")
	g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
	g.P("}")
	g.P()
	g.P("return")
	g.P("}")
	g.P()

	g.P("ctx = ", runtimePackage.Ident("ContextWithHeaders"), "(ctx, r.Header)")
	g.P()

	g.P("stream, err := ", runtimePackage.Ident("NewWebSocketStream"), "(w, r)")
	g.P("if err != nil {")
	g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
	g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
	g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
	g.P("}")
	g.P()
	g.P("return")
	g.P("}")
	g.P()

	g.P("defer stream.Close()")
	g.P()

	g.P("streamReq := x_", service.GoName, method.GoName, "WebSocket{")
	g.P("ServerStream: stream,")
	g.P("}")
	g.P()

	if err := ReadPath(g, method, "streamReq"); err != nil {
		return err
	}

	g.P()

	if err := ReadQuery(g, method, "streamReq"); err != nil {
		return err
	}

	g.P()

	if method.Desc.IsStreamingClient() {
		g.P("if err := server.", method.GoName, "(&streamReq); err != nil {")
		g.P("stream.WriteError(err)")
		g.P()
		g.P("return")
		g.P("}")
		g.P()
	} else {
		g.P("protoReq, err := streamReq.Recv()")
		g.P("if err != nil {")
		g.P("stream.WriteError(err)")
		g.P()
		g.P("return")
		g.P("}")
		g.P()

		g.P("if err := server.", method.GoName, "(protoReq, &streamReq); err != nil {")
		g.P("stream.WriteError(err)")
		g.P()
		g.P("return")
		g.P("}")
		g.P()
	}

	g.P("}")

	return nil
}
