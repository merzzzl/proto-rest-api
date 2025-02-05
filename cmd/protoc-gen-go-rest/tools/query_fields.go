package tools

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func QueryFields(method *protogen.Method) (map[string]*protogen.Field, error) {
	fields := make(map[string]*protogen.Field, 10)

	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return fields, nil
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return fields, nil
	}

	path := restRule.GetPath()

	if sep := strings.LastIndex(path, "?"); sep == -1 {
		return fields, nil
	} else {
		path = path[sep+1:]
	}

	for _, param := range strings.Split(path, "&") {
		if !strings.HasPrefix(param, ":") {
			continue
		}

		param = param[1:]

		paramPaths := strings.Split(param, ".")
		pFields := method.Input.Fields

		for i, p := range paramPaths {
			for _, field := range pFields {
				if field.Desc.TextName() == p {
					if i+1 == len(paramPaths) {
						fields[param] = field
					} else if field.Message != nil {
						pFields = field.Message.Fields
					} else {
						pFields = nil
					}

					break
				}
			}
		}

		if _, ok := fields[param]; !ok {
			return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, param)
		}
	}

	return fields, nil
}
