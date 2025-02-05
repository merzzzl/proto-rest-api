package gen

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func FileHeader(gen *protogen.Plugin, g *protogen.GeneratedFile, file *protogen.File, version string) error {
	v := gen.Request.GetCompilerVersion()

	var sb strings.Builder

	_, err := fmt.Fprintf(&sb, "v%d.%d.%d", v.GetMajor(), v.GetMinor(), v.GetPatch())
	if err != nil {
		return fmt.Errorf("failed to write compiler version: %w", err)
	}

	if s := v.GetSuffix(); s != "" {
		_, err := fmt.Fprintf(&sb, "-%s", s)
		if err != nil {
			return fmt.Errorf("failed to write compiler version suffix: %w", err)
		}
	}

	g.P("// Code generated by protoc-gen-go-rest. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-rest v", version)
	g.P("// - protoc             ", sb.String())

	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}

	g.P()
	g.P("package ", file.GoPackageName)

	return nil
}
