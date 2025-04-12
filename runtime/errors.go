package runtime

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var grpcToHTTPStatus = map[int]int{
	0:  200, // OK
	1:  499, // CANCELLED
	2:  500, // UNKNOWN
	3:  400, // INVALID_ARGUMENT
	4:  504, // DEADLINE_EXCEEDED
	5:  404, // NOT_FOUND
	6:  409, // ALREADY_EXISTS
	7:  403, // PERMISSION_DENIED
	8:  429, // RESOURCE_EXHAUSTED
	9:  400, // FAILED_PRECONDITION
	10: 409, // ABORTED
	11: 400, // OUT_OF_RANGE
	12: 501, // UNIMPLEMENTED
	13: 500, // INTERNAL
	14: 503, // UNAVAILABLE
	15: 500, // DATA_LOSS
	16: 401, // UNAUTHENTICATED
}

type ErrResponse struct {
	Message string          `json:"message"`
	Details []proto.Message `json:"details"`
}

func (e ErrResponse) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Message string            `json:"message"`
		Details []json.RawMessage `json:"details,omitempty"`
	}

	var details []json.RawMessage

	for _, detail := range e.Details {
		if detail != nil {
			jsonBytes, err := ProtoMarshal(detail)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal proto.Message to JSON: %w", err)
			}

			details = append(details, json.RawMessage(jsonBytes))
		}
	}

	alias := Alias{
		Message: e.Message,
		Details: details,
	}

	return json.Marshal(alias)
}

func GetHTTPStatusFromError(err error) *HTTPError {
	if err == nil {
		return NewError(http.StatusOK, "")
	}

	if errstatus, ok := err.(*HTTPError); ok {
		return errstatus
	}

	errCode := http.StatusInternalServerError
	errResp := ErrResponse{}

	if grpcStatus, ok := status.FromError(err); ok {
		errResp.Message = grpcStatus.Message()

		if httpStatus, ok := grpcToHTTPStatus[int(grpcStatus.Code())]; ok {
			errCode = httpStatus

			if len(grpcStatus.Details()) > 0 {
				details := make([]proto.Message, 0, len(grpcStatus.Details()))

				for _, detail := range grpcStatus.Details() {
					if anyMsg, ok := detail.(proto.Message); ok {
						details = append(details, anyMsg)
					}
				}

				errResp.Details = details
			}
		}
	}

	message, err := json.Marshal(&errResp)
	if err == nil {
		return NewError(errCode, string(message))
	}

	return NewError(errCode, string(message))
}
