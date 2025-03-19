// Code generated by protoc-gen-go-rest. DO NOT EDIT.
// versions:
// - protoc-gen-go-rest v0.0.0-alpha.0
// - protoc             v3.21.12
// source: example/proto/example.proto

package api

import (
	context "context"
	json "encoding/json"
	runtime "github.com/merzzzl/proto-rest-api/runtime"
	grpc "google.golang.org/grpc"
	io "io"
	http "net/http"
)

// EchoServiceWebServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceWebServer for forward compatibility.
type EchoServiceWebServer interface {
	Echo(EchoServiceEchoWebSocket) error
	Ticker(*TickerRequest, EchoServiceTickerWebSocket) error
	mustEmbedUnimplementedEchoServiceWebServer()
}

// UnimplementedEchoServiceWebServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceWebServer struct{}

func (UnimplementedEchoServiceWebServer) Echo(EchoServiceEchoWebSocket) error {
	return runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedEchoServiceWebServer) Ticker(*TickerRequest, EchoServiceTickerWebSocket) error {
	return runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedEchoServiceWebServer) mustEmbedUnimplementedEchoServiceWebServer() {}

type EchoServiceEchoWebSocket interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type x_EchoServiceEchoWebSocket struct {
	Channel string
	grpc.ServerStream
}

func (x *x_EchoServiceEchoWebSocket) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *x_EchoServiceEchoWebSocket) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)

	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}

	m.Channel = x.Channel

	return m, nil
}

type EchoServiceTickerWebSocket interface {
	Send(*TickerResponse) error
	grpc.ServerStream
}

type x_EchoServiceTickerWebSocket struct {
	Count int32
	grpc.ServerStream
}

func (x *x_EchoServiceTickerWebSocket) Send(m *TickerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *x_EchoServiceTickerWebSocket) Recv() (*TickerRequest, error) {
	m := new(TickerRequest)

	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}

	m.Count = x.Count

	return m, nil
}

// ExampleServiceWebServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceWebServer for forward compatibility.
type ExampleServiceWebServer interface {
	// POST new message to the server.
	PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error)
	// GET message from the server.
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	// DELETE message from the server.
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	// LIST messages from the server.
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	// PUT message to the server.
	PutMessage(context.Context, *PutMessageRequest) (*PutMessageResponse, error)
	// PATCH message to the server.
	PatchMessage(context.Context, *PatchMessageRequest) (*PatchMessageResponse, error)
	mustEmbedUnimplementedExampleServiceWebServer()
}

// UnimplementedExampleServiceWebServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceWebServer struct{}

func (UnimplementedExampleServiceWebServer) PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) PutMessage(context.Context, *PutMessageRequest) (*PutMessageResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) PatchMessage(context.Context, *PatchMessageRequest) (*PatchMessageResponse, error) {
	return nil, runtime.Errorf(http.StatusNotImplemented, "method not implemented")
}

func (UnimplementedExampleServiceWebServer) mustEmbedUnimplementedExampleServiceWebServer() {}

// RegisterEchoServiceHandler registers the http handlers for service EchoService to "mux".
func RegisterEchoServiceHandler(mux runtime.ServeMuxer, server EchoServiceWebServer, interceptors ...runtime.Interceptor) {
	router := runtime.NewRouter()

	router.Handle("GET", "/api/v1/echo/:channel", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerEchoServiceWebServerEcho(server, w, r, p, interceptors)
	})

	router.Handle("GET", "/api/v1/ticker/:count", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerEchoServiceWebServerTicker(server, w, r, p, interceptors)
	})

	mux.Handle("/api/v1/", router)
}

// RegisterExampleServiceHandler registers the http handlers for service ExampleService to "mux".
func RegisterExampleServiceHandler(mux runtime.ServeMuxer, server ExampleServiceWebServer, interceptors ...runtime.Interceptor) {
	router := runtime.NewRouter()

	router.Handle("POST", "/api/v1/example/messages", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerPostMessage(server, w, r, p, interceptors)
	})

	router.Handle("GET", "/api/v1/example/messages/:id", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerGetMessage(server, w, r, p, interceptors)
	})

	router.Handle("DELETE", "/api/v1/example/messages/:id", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerDeleteMessage(server, w, r, p, interceptors)
	})

	router.Handle("GET", "/api/v1/example/messages", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerListMessages(server, w, r, p, interceptors)
	})

	router.Handle("PUT", "/api/v1/example/messages/:message.id", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerPutMessage(server, w, r, p, interceptors)
	})

	router.Handle("PATCH", "/api/v1/example/messages/:message.id", func(w http.ResponseWriter, r *http.Request, p runtime.Params) {
		handlerExampleServiceWebServerPatchMessage(server, w, r, p, interceptors)
	})

	mux.Handle("/api/v1/example/", router)
}

func handlerEchoServiceWebServerEcho(server EchoServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	stream, err := runtime.NewWebSocketStream(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	defer stream.Close()

	streamReq := x_EchoServiceEchoWebSocket{
		ServerStream: stream,
	}

	streamReq.Channel = p.ByName("channel")

	if err := server.Echo(&streamReq); err != nil {
		stream.WriteError(err)

		return
	}

}

func handlerEchoServiceWebServerTicker(server EchoServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	stream, err := runtime.NewWebSocketStream(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	defer stream.Close()

	streamReq := x_EchoServiceTickerWebSocket{
		ServerStream: stream,
	}

	if v, err := runtime.ParseInt32(p.ByName("count")); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		streamReq.Count = v
	}

	protoReq, err := streamReq.Recv()
	if err != nil {
		stream.WriteError(err)

		return
	}

	if err := server.Ticker(protoReq, &streamReq); err != nil {
		stream.WriteError(err)

		return
	}

}

func handlerExampleServiceWebServerPostMessage(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, _ runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq PostMessageRequest

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if len(data) != 0 {
		fm, err := runtime.GetFieldMaskJS(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		ctx = runtime.ContextWithFieldMask(ctx, fm)
		err = runtime.ProtoUnmarshal(data, &protoReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}
	}

	msg, err := server.PostMessage(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	raw, err := runtime.ProtoMarshal(msg.GetMessage())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(raw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}

func handlerExampleServiceWebServerGetMessage(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq GetMessageRequest

	if v, err := runtime.ParseInt32(p.ByName("id")); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		protoReq.Id = v
	}

	msg, err := server.GetMessage(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	raw, err := runtime.ProtoMarshal(msg.GetMessage())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(raw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}

func handlerExampleServiceWebServerDeleteMessage(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq DeleteMessageRequest

	if v, err := runtime.ParseInt32(p.ByName("id")); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		protoReq.Id = v
	}

	_, err = server.DeleteMessage(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}

func handlerExampleServiceWebServerListMessages(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, _ runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq ListMessagesRequest

	if l, ok := r.URL.Query()["page"]; ok {
		for _, s := range l {
			v, err := runtime.ParseInt32(s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				if _, err := w.Write([]byte(err.Error())); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}

				return
			}

			protoReq.Page = v

			continue
		}
	}

	if l, ok := r.URL.Query()["per_page"]; ok {
		for _, s := range l {
			v, err := runtime.ParseInt32(s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				if _, err := w.Write([]byte(err.Error())); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}

				return
			}

			protoReq.PerPage = v

			continue
		}
	}

	if l, ok := r.URL.Query()["ids"]; ok {
		for _, s := range l {
			v, err := runtime.ParseInt32(s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				if _, err := w.Write([]byte(err.Error())); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}

				return
			}

			protoReq.Ids = append(protoReq.Ids, v)
		}
	}

	msg, err := server.ListMessages(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	raw, err := json.Marshal(msg.GetMessages())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(raw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}

func (x *Message) UnmarshalJSON(data []byte) error {
	return runtime.ProtoUnmarshal(data, x)
}

func (x *Message) MarshalJSON() ([]byte, error) {
	return runtime.ProtoMarshal(x)
}

func handlerExampleServiceWebServerPutMessage(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq PutMessageRequest

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if len(data) != 0 {
		fm, err := runtime.GetFieldMaskJS(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		ctx = runtime.ContextWithFieldMask(ctx, fm)

		var sub Message

		err = runtime.ProtoUnmarshal(data, &sub)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		protoReq.Message = &sub
	}
	if v, err := runtime.ParseInt32(p.ByName("message.id")); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		protoReq.Message.Id = v
	}

	_, err = server.PutMessage(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}

func handlerExampleServiceWebServerPatchMessage(server ExampleServiceWebServer, w http.ResponseWriter, r *http.Request, p runtime.Params, il []runtime.Interceptor) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx, err := runtime.ApplyInterceptors(ctx, r, il...)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	ctx = runtime.ContextWithHeaders(ctx, r.Header)

	var protoReq PatchMessageRequest

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if len(data) != 0 {
		fm, err := runtime.GetFieldMaskJS(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		ctx = runtime.ContextWithFieldMask(ctx, fm)

		var sub Message

		err = runtime.ProtoUnmarshal(data, &sub)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		protoReq.Message = &sub
	}
	if v, err := runtime.ParseInt32(p.ByName("message.id")); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		protoReq.Message.Id = v
	}

	_, err = server.PatchMessage(ctx, &protoReq)
	if err != nil {
		errstatus := runtime.GetHTTPStatusFromError(err)

		w.WriteHeader(errstatus.Code())
		if _, err := w.Write([]byte(errstatus.Message())); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}
