package gen

import (
	"fmt"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ReadQuery(g *protogen.GeneratedFile, method *protogen.Method, varName string) error {
	fields, err := tools.QueryFields(method)
	if err != nil {
		return fmt.Errorf("failed to get query fields for %s: %w", method.GoName, err)
	}

	for param, field := range fields {
		fullGoName := tools.FieldFullNmae(method.Input, param)

		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			g.P(varName, ".", fullGoName, " = r.URL.Query().Get(\"", param, "\")")
		case protoreflect.EnumKind:
			g.P(varName, ".", fullGoName, " = ", field.Enum.GoIdent, "(r.URL.Query().Get(\"", param, "\"))")

			g.P("if l", ", ok := r.URL.Query()[\"", param, "\"]; ok {")
			g.P("for _, s := range l", " {")
			g.P()

			if field.Desc.IsList() {
				g.P(varName, ".", fullGoName, " = append(", varName, ".", fullGoName, ", ", field.Enum.GoIdent, "(", field.Enum.GoIdent, "_value[s])")
			} else {
				g.P(varName, ".", fullGoName, " = ", field.Enum.GoIdent, "(", field.Enum.GoIdent, "_value[s])")
				g.P()
				g.P("continue")
			}

			g.P("}")
			g.P("}")
		case protoreflect.DoubleKind, protoreflect.FloatKind, protoreflect.BoolKind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind, protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Int32Kind, protoreflect.Sfixed32Kind:
			parser, err := tools.ValueParser(field)
			if err != nil {
				return fmt.Errorf("failed to get value parser for %s: %w", field.GoName, err)
			}

			g.P("if l", ", ok := r.URL.Query()[\"", param, "\"]; ok {")
			g.P("for _, s := range l", " {")
			g.P("v, err := ", runtimePackage.Ident(parser), "(s)")
			g.P("if err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
			g.P()
			g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("}")
			g.P()
			g.P("return")
			g.P("}")
			g.P()

			if field.Desc.IsList() {
				g.P(varName, ".", fullGoName, " = append(", varName, ".", fullGoName, ", v)")
			} else {
				g.P(varName, ".", fullGoName, " = v")
				g.P()
				g.P("continue")
			}

			g.P("}")
			g.P("}")
		default:
			return fmt.Errorf("unsupported type %s", field.Desc.Kind())
		}
	}

	return nil
}
