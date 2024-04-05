package runtime_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestServeHTTP_0(t *testing.T) {
	t.Parallel()

	router := runtime.NewRouter()
	router.Handle("GET", "/", func(w http.ResponseWriter, _ *http.Request, _ runtime.Params) {
		w.WriteHeader(http.StatusOK)
	})

	mux := httptest.NewServer(router)
	defer mux.Close()

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, mux.URL+"/", http.NoBody)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestApplyInterceptors_0(t *testing.T) {
	t.Parallel()

	interceptors := []runtime.Interceptor{
		func(_ context.Context, _ *http.Request) (context.Context, error) {
			return context.Background(), nil
		},
	}

	ctx, err := runtime.ApplyInterceptors(context.Background(), nil, interceptors...)
	require.NoError(t, err)
	require.NotNil(t, ctx)
}

func TestApplyInterceptors_1(t *testing.T) {
	t.Parallel()

	interceptors := []runtime.Interceptor{
		func(_ context.Context, _ *http.Request) (context.Context, error) {
			return nil, runtime.ErrInvalidJSON
		},
	}

	ctx, err := runtime.ApplyInterceptors(context.Background(), nil, interceptors...)
	require.Equal(t, runtime.ErrInvalidJSON, err)
	require.NotNil(t, ctx)
}
