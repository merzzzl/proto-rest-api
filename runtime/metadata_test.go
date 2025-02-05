package runtime_test

import (
	"testing"

	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestValidateMD(t *testing.T) {
	t.Parallel()

	t.Run("valid metadata", func(t *testing.T) {
		t.Parallel()

		md := metadata.MD{
			"key":         []string{"value1", "value2"},
			"another-key": []string{"value"},
		}

		err := runtime.ValidateMD(md)
		require.NoError(t, err, "ValidateMD should not return an error for valid metadata")
	})

	t.Run("empty header key", func(t *testing.T) {
		t.Parallel()

		md := metadata.MD{
			"": []string{"value"},
		}

		err := runtime.ValidateMD(md)
		require.ErrorIs(t, err, runtime.ErrEmptyHeaderKey, "ValidateMD should return ErrEmptyHeaderKey for empty key")
	})

	t.Run("non-printable characters in value", func(t *testing.T) {
		t.Parallel()

		md := metadata.MD{
			"key": []string{"value1", "val\x01ue2"},
		}

		err := runtime.ValidateMD(md)
		require.ErrorIs(t, err, runtime.ErrContainsNonPrintables, "ValidateMD should return ErrContainsNonPrintables for non-printable characters")
	})

	t.Run("illegal characters in key", func(t *testing.T) {
		t.Parallel()

		md := metadata.MD{
			"invalid:key": []string{"value"},
		}

		err := runtime.ValidateMD(md)
		require.Error(t, err, "ValidateMD should return an error for illegal characters in key")
		require.Contains(t, err.Error(), runtime.ErrContainsIllegal.Error(), "error should contain ErrContainsIllegal message")
	})

	t.Run("binary key suffix", func(t *testing.T) {
		t.Parallel()

		md := metadata.MD{
			"key-bin": []string{"binarydata"},
		}

		err := runtime.ValidateMD(md)
		require.NoError(t, err, "ValidateMD should not return an error for binary keys")
	})
}
