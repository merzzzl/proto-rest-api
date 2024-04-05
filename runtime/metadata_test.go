package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestValidateMD_0(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		"key": []string{"value"},
	}

	err := runtime.ValidateMD(md)
	require.NoError(t, err)
}

func TestValidateMD_1(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		"key with space": []string{"value"},
	}

	err := runtime.ValidateMD(md)
	require.ErrorIs(t, err, runtime.ErrContainsIllegal)
}

func TestValidateMD_2(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		"key": []string{"value\twith non printables"},
	}

	err := runtime.ValidateMD(md)
	require.ErrorIs(t, err, runtime.ErrContainsNonPrintables)
}

func TestValidateMD_3(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		":": []string{"value\twith non printables"},
	}

	err := runtime.ValidateMD(md)
	require.NoError(t, err)
}

func TestValidateMD_4(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		"": []string{"value"},
	}

	err := runtime.ValidateMD(md)
	require.ErrorIs(t, err, runtime.ErrEmptyHeaderKey)
}

func TestValidateMD_5(t *testing.T) {
	t.Parallel()

	md := metadata.MD{
		"key-bin": []string{"some\tbin\tdata"},
	}

	err := runtime.ValidateMD(md)
	require.NoError(t, err)
}
