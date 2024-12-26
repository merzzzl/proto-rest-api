package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/merzzzl/proto-rest-api/restapi"
)

func genHandler(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	serverType := service.GoName + serviceSufix

	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		exitWithError(fmt.Sprintf("unknown method options in %s", method.GoName))
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		exitWithError(fmt.Sprintf("unknown http options in %s", method.GoName))
	}

	if strings.Contains(restRule.GetPath(), "/:") {
		g.P("func handler", serverType, method.GoName, "(server ", serverType, ", w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", p ", runtimePackage.Ident("Params"), ", il []", runtimePackage.Ident("Interceptor"), ") {")
	} else {
		g.P("func handler", serverType, method.GoName, "(server ", serverType, ", w ", httpPackage.Ident("ResponseWriter"), ", r *", httpPackage.Ident("Request"), ", _ ", runtimePackage.Ident("Params"), ", il []", runtimePackage.Ident("Interceptor"), ") {")
	}

	genHandlerContext(g)

	if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
		genHandlerStream(g, service, method)
		genHandlerParsePath(g, method, restRule, "streamReq")
		genHandlerParseQuery(g, method, restRule, "streamReq")

		if method.Desc.IsStreamingClient() {
			genHandlerClientStream(g, method)
		} else {
			genHandlerNonClientStream(g, method)
		}
	} else {
		genHandlerRequest(g, method, restRule)
		genHandlerParsePath(g, method, restRule, "protoReq")
		genHandlerParseQuery(g, method, restRule, "protoReq")

		if restRule.GetResponse() != "" {
			genHandlerResponse(g, method, restRule)
		} else {
			genHandlerEmptyResponse(g, method, restRule)
		}
	}

	g.P("}")
	g.P()
}

func genHandlerClientStream(g *protogen.GeneratedFile, method *protogen.Method) {
	g.P("if err := server.", method.GoName, "(&streamReq); err != nil {")
	g.P("stream.WriteError(err)")
	g.P()
	g.P("return")
	g.P("}")
	g.P()
}

func genHandlerNonClientStream(g *protogen.GeneratedFile, method *protogen.Method) {
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

func genHandlerEmptyResponse(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) {
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
	g.P()
}

func genHandlerResponse(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) {
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
		field, goFieldsPath := fieldsPath(method.Output, fileds)

		if field.Desc.IsList() {
			if field.Message == nil {
				exitWithError(fmt.Sprintf("output field %s in %s is not a message", strings.Join(fileds, "."), method.Output.GoIdent))
			}

			genQueue = append(genQueue, func() {
				genStructList(g, field.Message)
			})

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
	g.P()
}

func genHandlerParseQuery(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule, varName string) {
	subPath := restRule.GetPath()

	sep := strings.LastIndex(subPath, "?")
	if sep == -1 {
		return
	}

	for _, param := range strings.Split(restRule.GetPath()[sep+1:], "&") {
		var found bool

		for _, field := range method.Input.Fields {
			if field.Desc.TextName() == param[1:] {
				found = true

				switch field.Desc.Kind() {
				case protoreflect.StringKind:
					g.P(varName, ".", field.GoName, " = r.URL.Query().Get(\"", param[1:], "\")")
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
					g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseInt32"), "(r.URL.Query().Get(\"", param[1:], "\"))")
					g.P("if err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
					g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
					g.P("}")
					g.P()
					g.P("return")
					g.P("}")
				case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
					g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseInt64"), "(r.URL.Query().Get(\"", param[1:], "\"))")
					g.P("if err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
					g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
					g.P("}")
					g.P()
					g.P("return")
					g.P("}")
				case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
					g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseUint32"), "(r.URL.Query().Get(\"", param[1:], "\"))")
					g.P("if err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
					g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
					g.P("}")
					g.P()
					g.P("return")
					g.P("}")
				case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
					g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseUint64"), "(r.URL.Query().Get(\"", param[1:], "\"))")
					g.P("if err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
					g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
					g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
					g.P("}")
					g.P()
					g.P("return")
					g.P("}")
				case protoreflect.EnumKind:
					g.P(varName, ".", field.GoName, " = ", field.Enum.GoIdent, "(r.URL.Query().Get(\"", param[1:], "\"))")
				case protoreflect.BoolKind, protoreflect.FloatKind, protoreflect.DoubleKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
					exitWithError(fmt.Sprintf("unknown field %s in %s", param[1:], method.Input.GoIdent))
				}

				g.P()
			}
		}

		if !found {
			exitWithError(fmt.Sprintf("unknown field %s in %s", param[1:], method.Input.GoIdent))
		}
	}
}

func genHandlerParsePath(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule, varName string) {
	subPath := restRule.GetPath()

	for _, segment := range strings.Split(subPath, "/") {
		if strings.HasPrefix(segment, ":") {
			var found bool

			for _, field := range method.Input.Fields {
				if field.Desc.TextName() == segment[1:] {
					found = true

					switch field.Desc.Kind() {
					case protoreflect.StringKind:
						g.P(varName, ".", field.GoName, " = p.ByName(\"", segment[1:], "\")")
					case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
						g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseInt32"), "(p.ByName(\"", segment[1:], "\"))")
						g.P("if err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
						g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
						g.P("}")
						g.P()
						g.P("return")
						g.P("}")
					case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
						g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseInt64"), "(p.ByName(\"", segment[1:], "\"))")
						g.P("if err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
						g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
						g.P("}")
						g.P()
						g.P("return")
						g.P("}")
					case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
						g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseUint32"), "(p.ByName(\"", segment[1:], "\"))")
						g.P("if err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
						g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
						g.P("}")
						g.P()
						g.P("return")
						g.P("}")
					case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
						g.P(varName, ".", field.GoName, ", err = ", runtimePackage.Ident("ParseUint64"), "(p.ByName(\"", segment[1:], "\"))")
						g.P("if err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
						g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
						g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
						g.P("}")
						g.P()
						g.P("return")
						g.P("}")
					case protoreflect.EnumKind:
						g.P(varName, ".", field.GoName, " = ", field.Enum.GoIdent, "(p.ByName(\"", segment[1:], "\"))")
					case protoreflect.BoolKind, protoreflect.FloatKind, protoreflect.DoubleKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
						exitWithError(fmt.Sprintf("unknown field %s in %s", segment[1:], method.Input.GoIdent))
					}

					g.P()
				}
			}

			if !found {
				exitWithError(fmt.Sprintf("unknown field %s in %s", segment[1:], method.Input.GoIdent))
			}
		}
	}
}

func genHandlerRequest(g *protogen.GeneratedFile, method *protogen.Method, restRule *restapi.MethodRule) {
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
		g.P()

		fileds := strings.Split(restRule.GetRequest(), ".")

		if fileds[0] == "" {
			fileds = fileds[1:]
		}

		if fileds[0] == "*" {
			fileds = []string{}
		}

		if len(fileds) != 0 {
			field, goFieldsPath := fieldsPath(method.Input, fileds)

			if field.Desc.IsList() {
				if field.Message == nil {
					exitWithError(fmt.Sprintf("output field %s in %s is not a message", strings.Join(fileds, "."), method.Output.GoIdent))
				}

				genQueue = append(genQueue, func() {
					genStructList(g, field.Message)
				})

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
		g.P()
	}
}

func genHandlerContext(g *protogen.GeneratedFile) {
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
}

func genHandlerStream(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
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

	g.P("streamReq := ", lowerFirst(service.GoName), method.GoName, "WebSocket{")
	g.P("ServerStream: stream,")
	g.P("}")
	g.P()
}

func fieldsPath(msg *protogen.Message, fileds []string) (*protogen.Field, []string) {
	for _, field := range msg.Fields {
		if field.Desc.TextName() == fileds[0] {
			if len(fileds) == 1 {
				return field, []string{field.GoName}
			}

			var goFieldsPath []string

			field, goFieldsPath = fieldsPath(field.Message, fileds[1:])

			return field, append([]string{field.GoName}, goFieldsPath...)
		}
	}

	exitWithError(fmt.Sprintf("unknown field %s in %s", fileds[0], msg.GoIdent))

	return nil, nil
}

func genStructList(g *protogen.GeneratedFile, msg *protogen.Message) {
	g.P("func (x *", msg.GoIdent.GoName, ") UnmarshalJSON(data []byte) error {")
	g.P("return ", runtimePackage.Ident("ProtoUnmarshal"), "(data, x)")
	g.P("}")
	g.P()

	g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	g.P("return ", runtimePackage.Ident("ProtoMarshal"), "(x)")
	g.P("}")
	g.P()
}
