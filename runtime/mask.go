package runtime

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldMask [][]string

func MergeByMask(src, dst protoreflect.ProtoMessage, mask FieldMask) {
	maskToMaskFilter(mask).overwrite(src.ProtoReflect(), dst.ProtoReflect())
}

type maskFilter map[string]maskFilter

func maskToMaskFilter(mask FieldMask) maskFilter {
	m := make(maskFilter)

	for _, paths := range mask {
		tmp := m
		for _, name := range paths {
			if tmp[name] == nil {
				tmp[name] = make(maskFilter)
			}

			tmp = tmp[name]
		}
	}

	return m
}

func (mf maskFilter) overwrite(src, dst protoreflect.Message) {
	for k, v := range mf {
		srcFD := src.Descriptor().Fields().ByJSONName(k)
		dstFD := dst.Descriptor().Fields().ByJSONName(k)

		if srcFD == nil || dstFD == nil {
			continue
		}

		if len(v) == 0 {
			if srcFD.Kind() == dstFD.Kind() {
				val := src.Get(srcFD)
				if isValid(srcFD, val) {
					dst.Set(dstFD, val)
				} else {
					dst.Clear(dstFD)
				}
			}
		} else if srcFD.Kind() == protoreflect.MessageKind {
			if !dst.Get(dstFD).Message().IsValid() {
				dst.Set(dstFD, protoreflect.ValueOf(dst.Get(dstFD).Message().New()))
			}

			v.overwrite(src.Get(srcFD).Message(), dst.Get(dstFD).Message())
		}
	}
}

func isValid(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
	switch {
	case fd.IsList():
		return val.List().IsValid()
	case fd.IsMap():
		return val.Map().IsValid()
	case fd.Message() != nil:
		return val.Message().IsValid()
	default:
		return true
	}
}
