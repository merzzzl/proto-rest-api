package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/merzzzl/proto-rest-api/cmd/protoc-gen-go-rest/tools"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Scheme(msg *protogen.Message, isList bool, exclude []string) *openapi3.Schema {
	schema := openapi3.NewObjectSchema()

	schema.Description = tools.LineComments(msg.Comments)
	requireFields := make([]string, 0)

	for _, field := range msg.Fields {
		if tools.Contains(exclude, field.GoIdent.GoName) {
			continue
		}

		fieldSchema := Field(field, exclude)

		if schema.Type.Is(openapi3.TypeObject) {
			if schema.Properties == nil {
				schema.Properties = make(openapi3.Schemas)
			}

			if !field.Desc.HasOptionalKeyword() && field.Oneof == nil {
				requireFields = append(requireFields, string(field.Desc.JSONName()))
			}

			schema.Properties[string(field.Desc.JSONName())] = openapi3.NewSchemaRef("", fieldSchema)
		} else if schema.Type.Is(openapi3.TypeArray) && schema.Items != nil {
			schema.Items = openapi3.NewSchemaRef("", fieldSchema)
		}
	}

	schema.Required = requireFields

	if isList {
		arraySchema := openapi3.NewArraySchema()
		arraySchema.Items = openapi3.NewSchemaRef("", schema)
		schema = arraySchema
	}

	return schema
}

func Field(field *protogen.Field, exclude []string) *openapi3.Schema {
	fieldSchema := openapi3.NewObjectSchema()

	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		fieldSchema = openapi3.NewBoolSchema()
	case protoreflect.Int32Kind, protoreflect.Uint32Kind:
		fieldSchema = openapi3.NewInt32Schema()
	case protoreflect.Int64Kind, protoreflect.Uint64Kind:
		fieldSchema = &openapi3.Schema{
			Type:   &openapi3.Types{openapi3.TypeString},
			Format: "int64",
		}
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		fieldSchema = openapi3.NewFloat64Schema()
	case protoreflect.StringKind:
		fieldSchema = openapi3.NewStringSchema()
	case protoreflect.BytesKind:
		fieldSchema = openapi3.NewBytesSchema()
	case protoreflect.MessageKind:
		if field.Message.Desc.FullName() == "google.protobuf.Timestamp" {
			fieldSchema = &openapi3.Schema{
				Type:   &openapi3.Types{openapi3.TypeString},
				Format: "date-time",
			}
		} else if field.Message.Desc.FullName() == "google.protobuf.Struct" {
			fieldSchema = openapi3.NewObjectSchema()
		} else {
			fieldSchema = Scheme(field.Message, false, exclude)
		}
	case protoreflect.EnumKind:
		fieldSchema = openapi3.NewStringSchema()

		for _, val := range field.Enum.Values {
			fieldSchema.Enum = append(fieldSchema.Enum, string(val.Desc.Name()))
		}
	}

	fieldSchema.Description = tools.LineComments(field.Comments)

	if field.Desc.IsList() {
		arraySchema := openapi3.NewArraySchema()
		arraySchema.Items = openapi3.NewSchemaRef("", fieldSchema)
		fieldSchema = arraySchema
	}

	return fieldSchema
}
