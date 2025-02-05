package gen

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/merzzzl/proto-rest-api/restapi"
)

var genQueue map[string]func()

func UnaryHandler(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	genQueue = make(map[string]func(), 0)

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

	if err := UnaryHandlerRequest(g, method, restRule); err != nil {
		return err
	}

	if err := ReadPath(g, method, "protoReq"); err != nil {
		return err
	}

	g.P()

	if err := ReadQuery(g, method, "protoReq"); err != nil {
		return err
	}

	g.P()

	if restRule.GetResponse() != "" {
		if err := UnaryHandlerResponse(g, method, restRule); err != nil {
			return err
		}
	} else {
		UnaryHandlerEmptyResponse(g, method, restRule)
	}

	g.P()
	g.P("}")

	for _, f := range genQueue {
		f()
	}

	return nil
}

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

func UnaryHandlerResponse(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) error {
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
		field, goFieldsPath, err := FieldsPath(method.Output, fileds)
		if err != nil {
			return err
		}

		if field.Desc.IsList() {
			if field.Message == nil {
				return fmt.Errorf("field %s in %s is not a message", strings.Join(fileds, "."), method.Output.GoIdent)
			}

			if field.Message != nil {
				genQueue[field.Message.GoIdent.GoName] = func() {
					RequiredStructList(g, field.Message)
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

func UnaryHandlerRequest(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) error {
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

			field, goFieldsPath, err := FieldsPath(method.Input, fileds)
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
						RequiredStructList(g, field.Message)
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

func FieldsPath(msg *protogen.Message, fileds []string) (*protogen.Field, []string, error) {
	var err error

	for _, field := range msg.Fields {
		if field.Desc.TextName() == fileds[0] {
			if len(fileds) == 1 {
				return field, []string{field.GoName}, nil
			}

			var goFieldsPath []string

			field, goFieldsPath, err = FieldsPath(field.Message, fileds[1:])
			if err != nil {
				return nil, nil, err
			}

			return field, append([]string{field.GoName}, goFieldsPath...), nil
		}
	}

	return nil, nil, fmt.Errorf("unknown field %s in %s", fileds[0], msg.GoIdent)
}

func RequiredStructList(g *protogen.GeneratedFile, msg *protogen.Message) {
	g.P("func (x *", msg.GoIdent.GoName, ") UnmarshalJSON(data []byte) error {")
	g.P("return ", runtimePackage.Ident("ProtoUnmarshal"), "(data, x)")
	g.P("}")
	g.P()

	g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	g.P("return ", runtimePackage.Ident("ProtoMarshal"), "(x)")
	g.P("}")
	g.P()
}
