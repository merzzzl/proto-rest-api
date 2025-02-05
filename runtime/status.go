package runtime

import "fmt"

type HTTPError struct {
	message string
	code    int
}

func NewError(code int, msg string) *HTTPError {
	return &HTTPError{code: code, message: msg}
}

func Error(code int, msg string) error {
	return NewError(code, msg)
}

func Errorf(code int, format string, a ...any) error {
	return Error(code, fmt.Sprintf(format, a...))
}

func (s *HTTPError) Code() int {
	return s.code
}

func (s *HTTPError) Message() string {
	return s.message
}

func (s *HTTPError) Error() string {
	return fmt.Sprintf("%s (%d)", s.message, s.code)
}
