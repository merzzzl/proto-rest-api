package tools

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValueParser(field *protogen.Field) (string, error) {
	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind:
		return "ParseInt32", nil
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind:
		return "ParseInt64", nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "ParseUint32", nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "ParseUint64", nil
	case protoreflect.BoolKind:
		return "ParseBool", nil
	case protoreflect.DoubleKind:
		return "ParseFloat64", nil
	case protoreflect.FloatKind:
		return "ParseFloat32", nil
	}

	return "", fmt.Errorf("%w: %v", ErrUnsupportedFieldType, field.Desc.Kind())
}
