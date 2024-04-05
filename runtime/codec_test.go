package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merzzzl/proto-rest-api/example/gen/go/example"
	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestProtoUnmarshal_0(t *testing.T) {
	t.Parallel()

	js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

	var in example.Message

	err := runtime.ProtoUnmarshal([]byte(js), &in)
	require.NoError(t, err)
}

func TestProtoUnmarshal_1(t *testing.T) {
	t.Parallel()

	js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

	var in string

	err := runtime.ProtoUnmarshal([]byte(js), in)
	require.ErrorIs(t, err, runtime.ErrMessageType)
}

func TestProtoMarshal_0(t *testing.T) {
	t.Parallel()

	var in example.Message

	_, err := runtime.ProtoMarshal(&in)
	require.NoError(t, err)
}

func TestProtoMarshal_1(t *testing.T) {
	t.Parallel()

	var in string

	_, err := runtime.ProtoMarshal(in)
	require.ErrorIs(t, err, runtime.ErrMessageType)
}
