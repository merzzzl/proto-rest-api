package gen

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
)

func UnaryHandlerEmptyResponse(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) {
	g.P("_, err = server.", method.GoName, "(ctx, &protoReq)")
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

	statusCode := 200
	if restRule.GetSuccessCode() != 0 {
		statusCode = int(restRule.GetSuccessCode())
	}

	g.P("w.WriteHeader(", statusCode, ")")
	g.P("if _, err := w.Write(nil); err != nil {")
	g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
	g.P("}")
	g.P("")
	g.P("return")
}

func UnaryHandlerResponse(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule, genQueue map[string]func()) error {
	g.P("msg, err := server.", method.GoName, "(ctx, &protoReq)")
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

	fileds := strings.Split(restRule.GetResponse(), ".")

	if fileds[0] == "" {
		fileds = fileds[1:]
	}

	if fileds[0] == "*" {
		fileds = []string{}
	}

	if len(fileds) != 0 {
		field, goFieldsPath, err := tools.FieldsPath(method.Output, fileds)
		if err != nil {
			return err
		}

		if field.Desc.IsList() {
			if field.Message == nil {
				return fmt.Errorf("field %s in %s is not a message", strings.Join(fileds, "."), method.Output.GoIdent)
			}

			if field.Message != nil {
				genQueue[field.Message.GoIdent.GoName] = func() {
					g.P("func (x *", field.Message.GoIdent.GoName, ") UnmarshalJSON(data []byte) error {")
					g.P("return ", runtimePackage.Ident("ProtoUnmarshal"), "(data, x)")
					g.P("}")
					g.P()

					g.P("func (x *", field.Message.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
					g.P("return ", runtimePackage.Ident("ProtoMarshal"), "(x)")
					g.P("}")
				}
			}

			g.P()
			g.P("raw, err := ", jsonPackage.Ident("Marshal"), "(msg.Get", strings.Join(goFieldsPath, "().Get"), "())")
			g.P("if err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("}")
			g.P("")
			g.P("return")
			g.P("}")
		} else {
			g.P("raw, err := ", runtimePackage.Ident("ProtoMarshal"), "(msg.Get", strings.Join(goFieldsPath, "().Get"), "())")
			g.P("if err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("}")
			g.P("")
			g.P("return")
			g.P("}")
		}
	} else {
		g.P("raw, err := ", runtimePackage.Ident("ProtoMarshal"), "(msg)")
		g.P("if err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
		g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
		g.P("}")
		g.P("")
		g.P("return")
		g.P("}")
	}

	g.P()

	statusCode := 200
	if restRule.GetSuccessCode() != 0 {
		statusCode = int(restRule.GetSuccessCode())
	}

	g.P("w.WriteHeader(", statusCode, ")")
	g.P("if _, err := w.Write(raw); err != nil {")
	g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
	g.P("}")
	g.P("")
	g.P("return")

	return nil
}
