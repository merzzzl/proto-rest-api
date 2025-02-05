package gen

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func UnaryHandlerRequest(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule, genQueue map[string]func()) error {
	g.P("var protoReq ", method.Input.GoIdent)
	g.P()

	if restRule.GetRequest() != "" {
		g.P("defer r.Body.Close()")
		g.P()
		g.P("data, err := ", ioPackage.Ident("ReadAll"), "(r.Body)")
		g.P("if err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
		g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
		g.P("}")
		g.P()
		g.P("return")
		g.P("}")
		g.P()

		g.P("if len(data) != 0 {")
		g.P("fm, err := ", runtimePackage.Ident("GetFieldMaskJS"), "(data)")
		g.P("if err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
		g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
		g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
		g.P("}")
		g.P()
		g.P("return")
		g.P("}")
		g.P()
		g.P("ctx = ", runtimePackage.Ident("ContextWithFieldMask"), "(ctx, fm)")

		fileds := strings.Split(restRule.GetRequest(), ".")

		if fileds[0] == "" {
			fileds = fileds[1:]
		}

		if fileds[0] == "*" {
			fileds = []string{}
		}

		if len(fileds) != 0 {
			g.P()

			field, goFieldsPath, err := tools.FieldsPath(method.Input, fileds)
			if err != nil {
				return err
			}

			if field.Desc.Kind() != protoreflect.MessageKind {
				return fmt.Errorf("field %s in %s is not a message", strings.Join(fileds, "."), method.Input.GoIdent)
			}

			if field.Desc.IsList() {
				if field.Message == nil {
					return fmt.Errorf("field %s in %s is not a message", strings.Join(fileds, "."), method.Input.GoIdent)
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

				g.P("var list []*", field.Message.GoIdent.GoName)
				g.P()

				g.P("err = ", jsonPackage.Ident("Unmarshal"), "(data, &list)")
				g.P("if err != nil {")
				g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
				g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
				g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
				g.P("}")
				g.P()
				g.P("return")
				g.P("}")
				g.P()

				g.P("protoReq.", strings.Join(goFieldsPath, "."), " = list")
			} else {
				if field.Message == nil {
					return fmt.Errorf("field %s in %s is not a message", strings.Join(fileds, "."), method.Input.GoIdent)
				}

				g.P("var sub ", field.Message.GoIdent.GoName)
				g.P()

				g.P("err = ", runtimePackage.Ident("ProtoUnmarshal"), "(data, &sub)")
				g.P("if err != nil {")
				g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
				g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
				g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
				g.P("}")
				g.P()
				g.P("return")
				g.P("}")
				g.P()

				g.P("protoReq.", strings.Join(goFieldsPath, "."), " = &sub")
			}
		} else {
			g.P("err = ", runtimePackage.Ident("ProtoUnmarshal"), "(data, &protoReq)")
			g.P("if err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
			g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("}")
			g.P()
			g.P("return")
			g.P("}")
		}

		g.P("}")
	}

	return nil
}
