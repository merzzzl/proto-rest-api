package tools

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func PathFields(method *protogen.Method) (map[string]*protogen.Field, error) {
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

	for _, segment := range strings.Split(path, "/") {
		if !strings.HasPrefix(segment, ":") {
			continue
		}

		segment = segment[1:]

		segmentPaths := strings.Split(segment, ".")
		pFields := method.Input.Fields

		for i, p := range segmentPaths {
			for _, field := range pFields {
				if field.Desc.TextName() == p {
					if i+1 == len(segmentPaths) {
						fields[segment] = field
					} else if field.Message != nil {
						pFields = field.Message.Fields
					} else {
						pFields = nil
					}

					break
				}
			}
		}

		if _, ok := fields[segment]; !ok {
			return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, segment)
		}
	}

	return fields, nil
}
