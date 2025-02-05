package tools

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func FieldsPath(msg *protogen.Message, fileds []string) (*protogen.Field, []string, error) {
	var err error

	for _, field := range msg.Fields {
		if field.Desc.TextName() == fileds[0] {
			if len(fileds) == 1 {
				return field, []string{field.GoName}, nil
			}

			var goFieldsPath []string

			field, goFieldsPath, err = FieldsPath(field.Message, fileds[1:])
			if err != nil {
				return nil, nil, err
			}

			return field, append([]string{field.GoName}, goFieldsPath...), nil
		}
	}

	return nil, nil, fmt.Errorf("%w: %s", ErrFieldNotFound, fileds[0])
}
