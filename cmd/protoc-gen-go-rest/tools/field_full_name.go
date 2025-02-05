package tools

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func FieldFullNmae(input *protogen.Message, key string) string {
	fullGoName := make([]string, 0)

	segmentPaths := strings.Split(key, ".")
	pFields := input.Fields

	for i, p := range segmentPaths {
		for _, field := range pFields {
			if field.Desc.TextName() == p {
				if i+1 == len(segmentPaths) {
					fullGoName = append(fullGoName, field.GoName)
				} else if field.Message != nil {
					pFields = field.Message.Fields
					fullGoName = append(fullGoName, field.GoName)
				} else {
					pFields = nil
				}

				break
			}
		}
	}

	return strings.Join(fullGoName, ".")
}
