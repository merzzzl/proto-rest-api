package gen

import (
	"fmt"

	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
)

func WebSocketStruct(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) error {
	g.P("type x_", service.GoName, method.GoName, "WebSocket struct {")

	pfileds, err := tools.PathFields(method)
	if err != nil {
		return fmt.Errorf("failed to get path fields for %s: %w", method.GoName, err)
	}

	for _, filed := range pfileds {
		var ftype string

		ftype, err = tools.TypeConverter(filed)
		if err != nil {
			return fmt.Errorf("failed to get field type for %s: %w", filed.GoName, err)
		}

		g.P(filed.GoName, " ", ftype)
	}

	qfileds, err := tools.QueryFields(method)
	if err != nil {
		return fmt.Errorf("failed to get query fields for %s: %w", method.GoName, err)
	}

	for _, filed := range qfileds {
		var ftype string

		ftype, err = tools.TypeConverter(filed)
		if err != nil {
			return fmt.Errorf("failed to get field type for %s: %w", filed.GoName, err)
		}

		g.P(filed.GoName, " ", ftype)
	}

	g.P(protogen.GoImportPath("google.golang.org/grpc").Ident("ServerStream"))
	g.P("}")

	return nil
}
