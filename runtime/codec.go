package runtime

import (
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var ErrMessageType = errors.New("message should be proto.Message")

func ProtoUnmarshal(b []byte, m any) error {
	pm, ok := m.(proto.Message)
	if !ok {
		return ErrMessageType
	}

	//nolint:wrapcheck // native protojson error
	return protojson.Unmarshal(b, pm)
}

func ProtoMarshal(m any) ([]byte, error) {
	pm, ok := m.(proto.Message)
	if !ok {
		return nil, ErrMessageType
	}

	//nolint:wrapcheck // native protojson error
	return protojson.Marshal(pm)
}
