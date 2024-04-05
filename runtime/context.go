package runtime

import (
	"context"
	"net/http"
)

type contextKeyFieldMask struct{}

var _contextKeyFieldMask = contextKeyFieldMask{}

func ContextWithFieldMask(ctx context.Context, fm FieldMask) context.Context {
	return context.WithValue(ctx, _contextKeyFieldMask, fm)
}

func FieldMaskFromContext(ctx context.Context) FieldMask {
	if fm, ok := ctx.Value(_contextKeyFieldMask).(FieldMask); ok {
		return fm
	}

	return nil
}

type contextKeyHeaders struct{}

var _contextKeyHeaders = contextKeyHeaders{}

func ContextWithHeaders(ctx context.Context, headers http.Header) context.Context {
	return context.WithValue(ctx, _contextKeyHeaders, headers)
}

func HeadersFromContext(ctx context.Context) http.Header {
	if headers, ok := ctx.Value(_contextKeyHeaders).(http.Header); ok {
		return headers
	}

	return nil
}
