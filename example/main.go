package main

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/merzzzl/proto-rest-api/example/api"
	pb "github.com/merzzzl/proto-rest-api/example/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/merzzzl/proto-rest-api/swagger"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	router := runtime.NewRouter()

	pb.RegisterExampleServiceHandler(router, NewExampleService())
	pb.RegisterEchoServiceHandler(router, NewEchoService())

	mux := router.Mux()

	mux.Handle("/swagger/", swagger.Handler(api.GetExampleSwagger()))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	_, _ = fmt.Println("Server is running on port 8080")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type EchoService struct {
	pb.UnimplementedEchoServiceWebServer
}

func NewEchoService() *EchoService {
	return &EchoService{}
}

func (EchoService) Ticker(m *pb.TickerRequest, s pb.EchoServiceTickerWebSocket) error {
	sendMessageCount := 0

	for range m.GetCount() {
		err := s.Send(&pb.TickerResponse{
			Timestamp: timestamppb.Now(),
		})
		if err != nil {
			return fmt.Errorf("send: %w", err)
		}

		sendMessageCount++

		time.Sleep(time.Second)
	}

	s.SetTrailer(metadata.MD{
		"ticker-count":   []string{strconv.Itoa(int(m.GetCount()))},
		"message-sended": []string{strconv.Itoa(sendMessageCount)},
	})

	return nil
}

func (EchoService) Echo(s pb.EchoServiceEchoWebSocket) error {
	recvMessageCount := 0
	sendMessageCount := 0

	for {
		m, err := s.Recv()
		if err != nil {
			return fmt.Errorf("recv: %w", err)
		}

		recvMessageCount++

		if err = s.SetHeader(metadata.MD{
			"message-received": []string{strconv.Itoa(recvMessageCount)},
		}); err != nil {
			return fmt.Errorf("set header: %w", err)
		}

		if m.GetMessage() == "" {
			break
		}

		err = s.Send(&pb.EchoResponse{
			Message: fmt.Sprintf("%s: %s", m.GetChannel(), m.GetMessage()),
		})
		if err != nil {
			return fmt.Errorf("send: %w", err)
		}

		sendMessageCount++
	}

	s.SetTrailer(metadata.MD{
		"message-received": []string{strconv.Itoa(recvMessageCount)},
		"message-sended":   []string{strconv.Itoa(sendMessageCount)},
	})

	return nil
}

type ExampleService struct {
	pb.UnimplementedExampleServiceWebServer

	storage map[int32]*pb.Message
	mutex   sync.RWMutex
}

func NewExampleService() *ExampleService {
	return &ExampleService{
		storage: make(map[int32]*pb.Message),
	}
}

func (s *ExampleService) PostMessage(_ context.Context, req *pb.PostMessageRequest) (*pb.PostMessageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := int32(len(s.storage) + 1)
	s.storage[id] = &pb.Message{
		Id:      id,
		Message: req.GetMessage(),
		Author:  req.GetAuthor(),
		Status:  req.GetStatus(),
	}

	return &pb.PostMessageResponse{Message: s.storage[id]}, nil
}

func (s *ExampleService) GetMessage(_ context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	message, ok := s.storage[req.GetId()]
	if !ok {
		st := status.New(codes.NotFound, "message not found")

		stDet, err := st.WithDetails(&errdetails.ErrorInfo{
			Reason:   "MESSAGE_NOT_FOUND",
			Domain:   "example.com",
			Metadata: map[string]string{"id": strconv.Itoa(int(req.GetId()))},
		})
		if err != nil {
			return nil, st.Err()
		}

		return nil, stDet.Err()
	}

	return &pb.GetMessageResponse{
		Message: message,
	}, nil
}

func (s *ExampleService) DeleteMessage(_ context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.storage[int32(req.GetId())]
	if !ok {
		return nil, status.Errorf(http.StatusNotFound, "message not found")
	}

	delete(s.storage, int32(req.GetId()))

	return &pb.DeleteMessageResponse{}, nil
}

func (s *ExampleService) ListMessages(_ context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	messages := make([]*pb.Message, 0, len(s.storage))

	for id, message := range s.storage {
		if len(req.GetIds()) > 0 {
			if !slices.Contains(req.GetIds(), id) {
				continue
			}
		}

		messages = append(messages, message)
	}

	from := (req.GetPage() - 1) * req.GetPerPage()
	to := (req.GetPage()-1)*req.GetPerPage() + req.GetPerPage()

	if to > int32(len(messages)) {
		to = int32(len(messages))
	}

	if from > to {
		from = to
	}

	messages = messages[from:to]

	return &pb.ListMessagesResponse{
		Messages: messages,
	}, nil
}

func (s *ExampleService) PutMessage(_ context.Context, req *pb.PutMessageRequest) (*pb.PutMessageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.storage[req.GetMessage().GetId()]
	if !ok {
		return nil, status.Errorf(http.StatusNotFound, "message not found")
	}

	s.storage[req.GetMessage().GetId()] = req.GetMessage()

	return &pb.PutMessageResponse{}, nil
}

func (s *ExampleService) PatchMessage(ctx context.Context, req *pb.PatchMessageRequest) (*pb.PatchMessageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	message, ok := s.storage[req.GetMessage().GetId()]
	if !ok {
		return nil, status.Errorf(http.StatusNotFound, "message not found")
	}

	fm := runtime.FieldMaskFromContext(ctx)
	if fm != nil {
		runtime.MergeByMask(req.GetMessage(), message, fm)
	}

	s.storage[req.GetMessage().GetId()] = message

	return &pb.PatchMessageResponse{
		Message: s.storage[req.GetMessage().GetId()],
	}, nil
}
