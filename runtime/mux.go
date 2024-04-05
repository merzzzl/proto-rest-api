package runtime

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Pattern []string

type ServeMuxer interface {
	Handle(pattern string, handler http.Handler)
}

type Router struct {
	Router *httprouter.Router
}

type HandlerFunc func(http.ResponseWriter, *http.Request, Params)

type Params interface {
	ByName(name string) string
}

func NewRouter() *Router {
	return &Router{
		Router: httprouter.New(),
	}
}

func (router *Router) Handle(method, path string, handler HandlerFunc) {
	router.Router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handler(w, r, p)
	})
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.Router.ServeHTTP(w, r)
}

type Interceptor func(ctx context.Context, req *http.Request) (context.Context, error)

func ApplyInterceptors(ctx context.Context, req *http.Request, interceptors ...Interceptor) (context.Context, error) {
	for _, interceptor := range interceptors {
		var err error

		ictx, err := interceptor(ctx, req)
		if ictx != nil {
			ctx = ictx
		}

		if err != nil {
			return ctx, err
		}
	}

	return ctx, nil
}
