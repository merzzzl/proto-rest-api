package runtime

import (
	"net/http"

	"google.golang.org/grpc/status"
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

func GetHTTPStatusFromError(err error) *HTTPError {
	if err == nil {
		return NewError(http.StatusOK, "")
	}

	if grpcStatus, ok := status.FromError(err); ok {
		if httpStatus, ok := grpcToHTTPStatus[int(grpcStatus.Code())]; ok {
			return NewError(httpStatus, grpcStatus.Message())
		}
	}

	if errstatus, ok := err.(*HTTPError); ok {
		return errstatus
	}

	return NewError(http.StatusInternalServerError, err.Error())
}
