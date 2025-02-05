package tools

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func ProtocVersion(gen *protogen.Plugin) string {
	if gen == nil || gen.Request == nil {
		return "(unknown)"
	}

	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}

	var sb strings.Builder

	fmt.Fprintf(&sb, "v%d.%d.%d", v.GetMajor(), v.GetMinor(), v.GetPatch())

	if s := v.GetSuffix(); s != "" {
		fmt.Fprintf(&sb, "-%s", s)
	}

	return sb.String()
}
