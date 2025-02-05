package tools

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TypeConverter(field *protogen.Field) (string, error) {
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		return "string", nil
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind:
		return "int32", nil
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind:
		return "int64", nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32", nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64", nil
	case protoreflect.EnumKind:
		return "int32", nil
	case protoreflect.BoolKind:
		return "bool", nil
	case protoreflect.FloatKind:
		return "float32", nil
	case protoreflect.DoubleKind:
		return "float64", nil
	}

	return "", fmt.Errorf("%w: %v", ErrUnsupportedFieldType, field.Desc.Kind())
}
