package status

import (
	"fmt"

	"google.golang.org/grpc/status"
)

// Status represents an RPC status code, message, and details.  It is immutable
// and should be created with New, Newf, or FromProto.
type Status struct {
	code    Code
	message string
}

// Code returns the status code contained in s.
func (s *Status) Code() Code {
	if s == nil {
		return OK
	}
	return s.code
}

// HTTPCode returns the mapped HTTP status code.
func (s *Status) HTTPCode() int {
	return HTTPStatusFromCode(s.Code())
}

// Message returns the message contained in s.
func (s *Status) Message() string {
	if s == nil {
		return ""
	}
	return s.message
}

func (s *Status) Error() string {
	return fmt.Sprintf("%s: %s", s.code, s.message)
}

// New returns a Status representing c and msg.
func New(c Code, msg string) *Status {
	return &Status{code: c, message: msg}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c Code, format string, a ...interface{}) *Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c Code, msg string) error {
	if c == OK {
		return nil
	}
	return New(c, msg)
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c Code, format string, a ...interface{}) error {
	return Error(c, fmt.Sprintf(format, a...))
}

// FromError returns a Status representing err if it was produced from this
// package. Otherwise, ok is false and a Status is returned with Unknown
// and the original error message.
func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return &Status{code: OK}, true
	}

	if s, ok := err.(*Status); ok {
		return s, true
	}

	if s, ok := status.FromError(err); ok {
		return New(Code(s.Code()), s.Message()), true
	}

	return New(Unknown, err.Error()), false
}

// Convert is a convenience function which removes the need to handle the
// boolean return value from FromError.
func Convert(err error) *Status {
	s, _ := FromError(err)
	return s
}

// Wrap returns nil if there is no error or returns a status.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return Convert(err)
}
