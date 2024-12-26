package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/merzzzl/proto-rest-api/example/api"
	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestProtoUnmarshal(t *testing.T) {
	t.Run("valid proto message", func(t *testing.T) {
		t.Parallel()

		js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

		var in pb.Message
		err := runtime.ProtoUnmarshal([]byte(js), &in)
		require.NoError(t, err, "ProtoUnmarshal should successfully parse valid input")
	})

	t.Run("non-proto message", func(t *testing.T) {
		t.Parallel()

		js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

		var in string
		err := runtime.ProtoUnmarshal([]byte(js), in)
		require.ErrorIs(t, err, runtime.ErrMessageType, "ProtoUnmarshal should return ErrMessageType for non-proto message")
	})
}

func TestProtoMarshal(t *testing.T) {
	t.Run("valid proto message", func(t *testing.T) {
		t.Parallel()

		var in pb.Message
		_, err := runtime.ProtoMarshal(&in)
		require.NoError(t, err, "ProtoMarshal should successfully marshal valid proto message")
	})

	t.Run("non-proto message", func(t *testing.T) {
		t.Parallel()

		var in string
		_, err := runtime.ProtoMarshal(in)
		require.ErrorIs(t, err, runtime.ErrMessageType, "ProtoMarshal should return ErrMessageType for non-proto message")
	})
}
