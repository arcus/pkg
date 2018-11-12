package service

import (
	"github.com/arcus/pkg/status"
	"github.com/golang/protobuf/proto"
)

var (
	ErrMismatchedRequestHandler = status.New(status.Internal, "mismatched request type for handler")
	ErrHandlerNotFound          = status.New(status.NotFound, "handler not found")
	ErrHandlerNotImplemented    = status.New(status.Unimplemented, "handler not implemented")
)

// Handler is a function that receives a command or query and returns a response and an error.
type Handler func(ctx *Context, msg proto.Message) (proto.Message, error)

// Middleware is a function that wraps a Handler with some additional processing.
type Middleware func(next Handler) Handler

// NotFoundHandler returns ErrHandlerNotFound.
func NotFoundHandler(ctx *Context, req proto.Message) (proto.Message, error) {
	return nil, ErrHandlerNotFound
}

// NotImplementedHandler returns ErrHandlerNotImplemented.
func NotImplementedHandler(ctx *Context, req proto.Message) (proto.Message, error) {
	return nil, ErrHandlerNotImplemented
}
