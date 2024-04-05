package runtime_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestGetHTTPStatusFromError_0(t *testing.T) {
	t.Parallel()

	err := runtime.ErrInvalidJSON
	errstatus := runtime.GetHTTPStatusFromError(err)

	require.Equal(t, err.Error(), errstatus.Message())
	require.Equal(t, http.StatusInternalServerError, errstatus.Code())
}

func TestGetHTTPStatusFromError_1(t *testing.T) {
	t.Parallel()

	errstatus := runtime.GetHTTPStatusFromError(nil)

	require.Equal(t, "", errstatus.Message())
	require.Equal(t, http.StatusOK, errstatus.Code())
}

func TestGetHTTPStatusFromError_2(t *testing.T) {
	t.Parallel()

	err := runtime.Error(http.StatusBadRequest, "foo")
	errstatus := runtime.GetHTTPStatusFromError(err)

	require.EqualValues(t, err, errstatus)
}

func TestGetHTTPStatusFromError_3(t *testing.T) {
	t.Parallel()

	err := status.Error(codes.NotFound, "foo")
	errstatus := runtime.GetHTTPStatusFromError(err)

	require.Equal(t, "foo", errstatus.Message())
	require.Equal(t, http.StatusNotFound, errstatus.Code())
}
