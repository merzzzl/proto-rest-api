package gen

import (
	"fmt"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ReadPath(g *protogen.GeneratedFile, method *protogen.Method, varName string) error {
	fields, err := tools.PathFields(method)
	if err != nil {
		return fmt.Errorf("failed to get path fields for %s: %w", method.GoName, err)
	}

	for param, field := range fields {
		fullGoName := tools.FieldFullNmae(method.Input, param)

		if field.Desc.HasOptionalKeyword() {
			return fmt.Errorf("optional fields are not supported in path")
		}

		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			g.P(varName, ".", fullGoName, " = p.ByName(\"", param, "\")")
		case protoreflect.EnumKind:
			g.P(varName, ".", fullGoName, " = ", field.Enum.GoIdent, "(", field.Enum.GoIdent, "_value[p.ByName(\"", param, "\")])")
		case protoreflect.DoubleKind, protoreflect.FloatKind, protoreflect.BoolKind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind, protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Int32Kind, protoreflect.Sfixed32Kind:
			parser, err := tools.ValueParser(field)
			if err != nil {
				return fmt.Errorf("failed to get value parser for %s: %w", field.GoName, err)
			}

			g.P("if v, err := ", runtimePackage.Ident(parser), "(p.ByName(\"", param, "\")); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusBadRequest"), ")")
			g.P()
			g.P("if _, err := w.Write([]byte(err.Error())); err != nil {")
			g.P("w.WriteHeader(", httpPackage.Ident("StatusInternalServerError"), ")")
			g.P("}")
			g.P()
			g.P("return")
			g.P("} else {")
			if field.Desc.HasOptionalKeyword() {
				g.P(varName, ".", fullGoName, " = &v")
			} else {
				g.P(varName, ".", fullGoName, " = v")
			}
			g.P("}")
		default:
			return fmt.Errorf("unsupported type %s", field.Desc.Kind())
		}

		g.P()
	}

	return nil
}
