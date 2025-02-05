package runtime_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/stretchr/testify/require"
)

func TestContextWithFieldMask(t *testing.T) {
	t.Parallel()

	t.Run("set and get field mask", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		fieldMask := runtime.FieldMask{{"field1", "subfield1"}, {"field2"}}
		ctx = runtime.ContextWithFieldMask(ctx, fieldMask)

		result := runtime.FieldMaskFromContext(ctx)
		require.NotNil(t, result, "field mask should not be nil")
		require.Equal(t, fieldMask, result, "retrieved field mask should match set value")
	})

	t.Run("get field mask from empty context", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result := runtime.FieldMaskFromContext(ctx)
		require.Nil(t, result, "field mask should be nil for empty context")
	})
}

func TestContextWithHeaders(t *testing.T) {
	t.Parallel()

	t.Run("set and get headers", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		headers := http.Header{
			"Key1": []string{"Value1"},
			"Key2": []string{"Value2"},
		}
		ctx = runtime.ContextWithHeaders(ctx, headers)

		result := runtime.HeadersFromContext(ctx)
		require.NotNil(t, result, "headers should not be nil")
		require.Equal(t, headers, result, "retrieved headers should match set value")
	})

	t.Run("get headers from empty context", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result := runtime.HeadersFromContext(ctx)
		require.Nil(t, result, "headers should be nil for empty context")
	})
}
