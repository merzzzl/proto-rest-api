package runtime

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockWebSocketConn struct {
	mock.Mock
}

func (m *MockWebSocketConn) WriteMessage(messageType int, data []byte) error {
	args := m.Called(messageType, data)

	return args.Error(0)
}

func (m *MockWebSocketConn) ReadMessage() (int, []byte, error) {
	args := m.Called()

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (m *MockWebSocketConn) WriteControl(messageType int, data []byte, deadline time.Time) error {
	args := m.Called(messageType, data, deadline)

	return args.Error(0)
}

func (m *MockWebSocketConn) SetCloseHandler(h func(code int, text string) error) {
	m.Called(h)
}

func (m *MockWebSocketConn) Close() error {
	args := m.Called()

	return args.Error(0)
}

func TestStreamWithMocks(t *testing.T) {
	t.Parallel()

	t.Run("NewWebSocketStream", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stream, err := NewWebSocketStream(w, r)
			require.NoError(t, err, "NewWebSocketStream should not return an error")
			require.NotNil(t, stream, "NewWebSocketStream should return a stream")
		}))
		defer server.Close()

		url := "ws" + server.URL[len("http"):] // Convert HTTP URL to WS URL
		_, _, err := websocket.DefaultDialer.Dial(url, nil)
		require.NoError(t, err, "Dial should not return an error")
	})

	t.Run("Context", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		stream := &Stream{
			context: context.Background(),
			conn:    mockConn,
		}

		ctx := stream.Context()
		require.NotNil(t, ctx, "Context should return a valid context")
	})

	t.Run("RecvMsg", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		msg := timestamppb.Now()
		bts, _ := ProtoMarshal(msg)

		mockConn.On("ReadMessage").Return(websocket.TextMessage, bts, nil)

		stream := &Stream{
			conn: mockConn,
		}

		var out timestamppb.Timestamp

		err := stream.RecvMsg(&out)
		require.NoError(t, err, "RecvMsg should not return an error")
		require.Equal(t, out.GetSeconds(), msg.GetSeconds(), "RecvMsg should correctly unmarshal the message")
	})

	t.Run("SetTrailer", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		stream := &Stream{
			conn:    mockConn,
			trailer: metadata.MD{},
		}

		trailer := metadata.Pairs("key", "value")
		stream.SetTrailer(trailer)
		require.Equal(t, trailer, stream.trailer, "SetTrailer should correctly set the trailer metadata")
	})

	t.Run("WriteError", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		mockConn.On("WriteControl", websocket.CloseMessage, mock.Anything, mock.Anything).
			Return(nil)

		stream := &Stream{
			conn: mockConn,
		}

		err := stream.WriteError(errors.New("test error"))
		require.NoError(t, err, "WriteError should not return an error")
		mockConn.AssertCalled(t, "WriteControl", websocket.CloseMessage, mock.Anything, mock.Anything)
	})

	t.Run("SendHeader", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		mockConn.On("WriteMessage", websocket.TextMessage, mock.Anything).
			Return(nil)

		stream := &Stream{
			conn:   mockConn,
			header: metadata.MD{},
		}

		headers := metadata.Pairs("key", "value")
		err := stream.SendHeader(headers)
		require.NoError(t, err, "SendHeader should not return an error")
		mockConn.AssertCalled(t, "WriteMessage", websocket.TextMessage, mock.Anything)
	})

	t.Run("Close", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		mockConn.On("WriteMessage", websocket.TextMessage, mock.Anything).
			Return(nil)
		mockConn.On("Close").
			Return(nil)

		stream := &Stream{
			conn:    mockConn,
			trailer: metadata.Pairs("trailer-key", "trailer-value"),
		}

		err := stream.Close()
		require.NoError(t, err, "Close should not return an error")
		mockConn.AssertCalled(t, "WriteMessage", websocket.TextMessage, mock.Anything)
		mockConn.AssertCalled(t, "Close")
	})

	t.Run("SendMsg", func(t *testing.T) {
		t.Parallel()

		mockConn := new(MockWebSocketConn)
		mockConn.On("WriteMessage", websocket.TextMessage, mock.Anything).
			Return(nil)

		stream := &Stream{
			conn: mockConn,
		}

		msg := &emptypb.Empty{}
		err := stream.SendMsg(msg)
		require.NoError(t, err, "SendMsg should not return an error")
		mockConn.AssertCalled(t, "WriteMessage", websocket.TextMessage, mock.Anything)
	})
}
