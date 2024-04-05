package runtime_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestContextWithFieldMask_0(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	ifm := runtime.FieldMask{
		[]string{"foo", "bar"},
	}

	ctx = runtime.ContextWithFieldMask(ctx, ifm)
	ofm := runtime.FieldMaskFromContext(ctx)

	require.Equal(t, ifm, ofm)
}

func TestFieldMaskFromContext_0(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	ofm := runtime.FieldMaskFromContext(ctx)

	require.Nil(t, ofm)
}

func TestContextWithHeaders_0(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	ih := http.Header{
		"foo": {"bar"},
	}

	ctx = runtime.ContextWithHeaders(ctx, ih)
	oh := runtime.HeadersFromContext(ctx)

	require.Equal(t, ih, oh)
}

func TestHeadersFromContext_0(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	oh := runtime.HeadersFromContext(ctx)

	require.Nil(t, oh)
}
