package runtime

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc/metadata"
)

type Conn interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (int, []byte, error)
	WriteControl(messageType int, data []byte, deadline time.Time) error
	SetCloseHandler(h func(code int, text string) error)
	Close() error
}

type Stream struct {
	context context.Context
	conn    Conn
	header  metadata.MD
	trailer metadata.MD
	closed  atomic.Bool
}

var (
	ErrPingTimeout       = errors.New("ping timeout")
	ErrConnAlreadyClosed = errors.New("connection already closed")
)

func NewWebSocketStream(w http.ResponseWriter, r *http.Request) (*Stream, error) {
	ctx, cancel := context.WithCancel(r.Context())
	r = r.WithContext(ctx)

	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		cancel()

		return nil, fmt.Errorf("upgrader conn: %w", err)
	}

	s := &Stream{
		context: ctx,
		conn:    conn,
		header:  metadata.MD{},
		trailer: metadata.MD{},
	}

	s.conn.SetCloseHandler(func(int, string) error {
		cancel()

		return nil
	})

	return s, nil
}

func (s *Stream) Context() context.Context {
	return s.context
}

func (s *Stream) WriteError(e error) error {
	if err := s.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, e.Error()), time.Now().Add(time.Second*5)); err != nil {
		return fmt.Errorf("close conn: %w", err)
	}

	return nil
}

func (s *Stream) Close() error {
	if !s.closed.CompareAndSwap(false, true) {
		return ErrConnAlreadyClosed
	}

	b, err := json.Marshal(s.trailer)
	if err != nil {
		return fmt.Errorf("json marshal trailer: %w", err)
	}

	b = append([]byte(`{"trailer":`), b...)
	b = append(b, []byte(`}`)...)

	if err := s.conn.WriteMessage(websocket.TextMessage, b); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	if err := s.conn.Close(); err != nil {
		return fmt.Errorf("close conn: %w", err)
	}

	return nil
}

func (s *Stream) SendMsg(m any) error {
	b, err := ProtoMarshal(m)
	if err != nil {
		return fmt.Errorf("json marshal message: %w", err)
	}

	if err := s.SendHeader(metadata.MD{}); err != nil {
		return fmt.Errorf("send header: %w", err)
	}

	if err := s.conn.WriteMessage(websocket.TextMessage, b); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	return nil
}

func (s *Stream) RecvMsg(m any) error {
	_, b, err := s.conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read message: %w", err)
	}

	if err := ProtoUnmarshal(b, m); err != nil {
		return fmt.Errorf("json unmarshal message: %w", err)
	}

	return nil
}

func (s *Stream) SetHeader(md metadata.MD) error {
	if md.Len() == 0 {
		return nil
	}

	if err := ValidateMD(md); err != nil {
		return fmt.Errorf("validate metadate: %w", err)
	}

	s.header = metadata.Join(s.header, md)

	return nil
}

func (s *Stream) SendHeader(md metadata.MD) error {
	if err := s.SetHeader(md); err != nil {
		return err
	}

	if len(s.header) == 0 {
		return nil
	}

	b, err := json.Marshal(s.header)
	if err != nil {
		return fmt.Errorf("json marshal header: %w", err)
	}

	b = append([]byte(`{"header":`), b...)
	b = append(b, []byte(`}`)...)

	if err := s.conn.WriteMessage(websocket.TextMessage, b); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	s.header = metadata.MD{}

	return nil
}

func (s *Stream) SetTrailer(md metadata.MD) {
	if md.Len() == 0 {
		return
	}

	if err := ValidateMD(md); err != nil {
		return
	}

	s.trailer = metadata.Join(s.trailer, md)
}
