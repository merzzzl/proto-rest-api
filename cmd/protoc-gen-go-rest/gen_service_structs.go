package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/merzzzl/proto-rest-api/gen/go/rest/api"
)

func genServiceStructs(g *protogen.GeneratedFile, service *protogen.Service) {
	mustOrShould := "must"
	if !*requireUnimplemented {
		mustOrShould = "should"
	}

	serverType := service.GoName + serviceSufix

	g.P("// ", serverType, " is the server API for ", service.GoName, " service.")
	g.P("// All implementations ", mustOrShould, " embed Unimplemented", serverType, " for forward compatibility.")

	if serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions); ok && serviceOptions.GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}

	g.P("type ", serverType, " interface {")

	for _, method := range service.Methods {
		if methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions); ok && methodOptions.GetDeprecated() {
			g.P(deprecationComment)
		}

		g.P(method.Comments.Leading, serverSignature(g, method))
	}

	if *requireUnimplemented {
		g.P("mustEmbedUnimplemented", serverType, "()")
	}

	g.P("}")
	g.P()

	mustOrShould = "must"

	if !*requireUnimplemented {
		mustOrShould = "should"
	}

	for _, method := range service.Methods {
		genSend := method.Desc.IsStreamingServer()
		genSendAndClose := !method.Desc.IsStreamingServer()
		genRecv := method.Desc.IsStreamingClient()

		if !genSend && !genRecv {
			continue
		}

		g.P("type ", service.GoName, method.GoName, "WebSocket interface {")

		if genSend {
			g.P("Send(*", method.Output.GoIdent, ") error")
		}

		if genSendAndClose {
			g.P("SendAndClose(*", method.Output.GoIdent, ") error")
		}

		if genRecv {
			g.P("Recv() (*", method.Input.GoIdent, ", error)")
		}

		g.P(grpcPackage.Ident("ServerStream"))
		g.P("}")
		g.P()

		g.P("type ", lowerFirst(service.GoName), method.GoName, "WebSocket struct {")

		fileds := listOfPathFields(method)

		for name, kind := range fileds {
			g.P(name, " ", kind)
		}

		g.P(grpcPackage.Ident("ServerStream"))
		g.P("}")
		g.P()

		if genSend {
			g.P("func (x *", lowerFirst(service.GoName), method.GoName, "WebSocket) Send(m *", method.Output.GoIdent, ") error {")
			g.P("return x.ServerStream.SendMsg(m)")
			g.P("}")
			g.P()
		}

		if genSendAndClose {
			g.P("func (x *", lowerFirst(service.GoName), method.GoName, "WebSocket) SendAndClose(m *", method.Output.GoIdent, ") error {")
			g.P("return x.ServerStream.SendMsg(m)")
			g.P("}")
			g.P()
		}

		g.P("func (x *", lowerFirst(service.GoName), method.GoName, "WebSocket) Recv() (*", method.Input.GoIdent, ", error) {")
		g.P("m := new(", method.Input.GoIdent, ")")
		g.P()
		g.P("if err := x.ServerStream.RecvMsg(m); err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P()

		for name := range fileds {
			g.P("m.", name, " = ", "x.", name)
		}

		g.P()
		g.P("return m, nil")
		g.P("}")
		g.P()
	}

	g.P("// Unimplemented", serverType, " ", mustOrShould, " be embedded to have forward compatible implementations.")
	g.P("type Unimplemented", serverType, " struct {}")
	g.P()

	for _, method := range service.Methods {
		nilArg := ""

		if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
			nilArg = "nil,"
		}

		g.P("func (Unimplemented", serverType, ") ", serverSignature(g, method), "{")
		g.P("return ", nilArg, runtimePackage.Ident("Errorf"), "(", httpPackage.Ident("StatusNotImplemented"), `, "method not implemented")`)
		g.P("}")
		g.P()
	}

	if *requireUnimplemented {
		g.P("func (Unimplemented", serverType, ") mustEmbedUnimplemented", serverType, "() {}")
	}

	g.P()
}

func listOfPathFields(method *protogen.Method) map[string]string {
	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		exitWithError(fmt.Sprintf("unknown method options in %s", method.GoName))
	}

	extVal := proto.GetExtension(methodOptions, api.E_Method)

	restRule, ok := extVal.(*api.MethodRule)
	if !ok {
		exitWithError(fmt.Sprintf("unknown http options in %s", method.GoName))
	}

	subPath := restRule.GetPath()

	fields := make(map[string]string)

	if sep := strings.LastIndex(subPath, "?"); sep != -1 {
		subPath = subPath[:sep]

		for _, param := range strings.Split(restRule.GetPath()[sep+1:], "&") {
			var found bool

			for _, field := range method.Input.Fields {
				if field.Desc.TextName() == param[1:] {
					found = true

					switch field.Desc.Kind() {
					case protoreflect.StringKind:
						fields[field.GoName] = "string"
					case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
						fields[field.GoName] = "int32"
					case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
						fields[field.GoName] = "int64"
					case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
						fields[field.GoName] = "uint32"
					case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
						fields[field.GoName] = "uint64"
					case protoreflect.EnumKind:
						fields[field.GoName] = "int32"
					case protoreflect.BoolKind, protoreflect.FloatKind, protoreflect.DoubleKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
						exitWithError(fmt.Sprintf("unknown field %s in %s", param[1:], method.Input.GoIdent))
					}
				}
			}

			if !found {
				exitWithError(fmt.Sprintf("unknown field %s in %s", param[1:], method.Input.GoIdent))
			}
		}
	}

	for _, segment := range strings.Split(subPath, "/") {
		if strings.HasPrefix(segment, ":") {
			var found bool

			for _, field := range method.Input.Fields {
				if field.Desc.TextName() == segment[1:] {
					found = true

					switch field.Desc.Kind() {
					case protoreflect.StringKind:
						fields[field.GoName] = "string"
					case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
						fields[field.GoName] = "int32"
					case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
						fields[field.GoName] = "int64"
					case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
						fields[field.GoName] = "uint32"
					case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
						fields[field.GoName] = "uint64"
					case protoreflect.EnumKind:
						fields[field.GoName] = "int32"
					case protoreflect.BoolKind, protoreflect.FloatKind, protoreflect.DoubleKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
						exitWithError(fmt.Sprintf("unknown field %s in %s", segment[1:], method.Input.GoIdent))
					}
				}
			}

			if !found {
				exitWithError(fmt.Sprintf("unknown field %s in %s", segment[1:], method.Input.GoIdent))
			}
		}
	}

	return fields
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
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

	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}
