package gen

import (
	"fmt"
	"slices"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ReadQuery(g *protogen.GeneratedFile, method *protogen.Method, varName string) error {
	fields, err := tools.QueryFields(method)
	if err != nil {
		return fmt.Errorf("failed to get query fields for %s: %w", method.GoName, err)
	}

	params := make([]string, 0, len(fields))

	for param := range fields {
		params = append(params, param)
	}

	slices.Sort(params)

	for _, param := range params {
		field := fields[param]
		fullGoName := tools.FieldFullNmae(method.Input, param)

		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			if field.Desc.HasOptionalKeyword() {
				g.P("if s := r.URL.Query().Get(\"", param, "\"); s != \"\" {")
				g.P(varName, ".", fullGoName, " = &s")
				g.P("}")
			} else {
				g.P(varName, ".", fullGoName, " = r.URL.Query().Get(\"", param, "\")")
			}
		case protoreflect.EnumKind:
			g.P("if l", ", ok := r.URL.Query()[\"", param, "\"]; ok {")
			g.P("for _, s := range l", " {")

			if field.Desc.IsList() {
				g.P(varName, ".", fullGoName, " = append(", varName, ".", fullGoName, ", ", field.Enum.GoIdent, "(", field.Enum.GoIdent, "_value[s]))")
			} else {
				if field.Desc.HasOptionalKeyword() {
					g.P("if v, ok := ", field.Enum.GoIdent, "_value[s]; ok {")
					g.P(varName, ".", fullGoName, " = ", field.Enum.GoIdent, "(v).Enum()")
					g.P("}")
				} else {
					g.P(varName, ".", fullGoName, " = ", field.Enum.GoIdent, "(", field.Enum.GoIdent, "_value[s])")
				}
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

			if field.Desc.IsList() {
				g.P("if len(l) == 1 && ", stringsPackage.Ident("Contains"), "(l[0], \",\")", " {")
				g.P("l = ", stringsPackage.Ident("Split"), "(l[0], \",\")")
				g.P("}")
				g.P()
			}

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
				if field.Desc.HasOptionalKeyword() {
					g.P(varName, ".", fullGoName, " = &v")
				} else {
					g.P(varName, ".", fullGoName, " = v")
				}
				g.P()
				g.P("continue")
			}

			g.P("}")
			g.P("}")
		default:
			return fmt.Errorf("unsupported type %s", field.Desc.Kind())
		}

		g.P()
	}

	return nil
}
