package gen

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/openapi"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func UnaryHandler(g *protogen.GeneratedFile, file *protogen.File, service *protogen.Service, method *protogen.Method) error {
	genQueue := make(map[string]func(), 0)

	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fmt.Errorf("unknown http options in %s", method.GoName)
	}

	if err := openapi.AddMethod(g, service, method); err != nil {
		return err
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

	if err := UnaryHandlerRequest(g, method, restRule, genQueue); err != nil {
		return err
	}

	if err := ReadPath(g, method, "protoReq"); err != nil {
		return err
	}

	if err := ReadQuery(g, method, "protoReq"); err != nil {
		return err
	}

	g.P()

	if restRule.GetResponse() != "" {
		if err := UnaryHandlerResponse(g, method, restRule, genQueue); err != nil {
			return err
		}
	} else {
		UnaryHandlerEmptyResponse(g, method, restRule)
	}

	g.P()
	g.P("}")

	for _, f := range genQueue {
		g.P()
		f()
	}

	return nil
}
