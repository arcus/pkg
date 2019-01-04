package service

import (
	"context"
	"time"

	"github.com/arcus/pkg/service/internal/pb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// Context is a custom implementation of the context.Context interface that
// adds application specific data.
type Context struct {
	ServiceName    string
	ServiceVersion string

	Command proto.Message

	AccessToken   string
	PrincipalId   string
	PrincipalRole string

	ctx context.Context
}

// WithContext wraps a context.Context with the custom Context.
func WithContext(ctx context.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}

func (c *Context) Deadline() (time.Time, bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

// Proto outputs a proto.Message that contains the custom application data from the Context.
func (c *Context) Proto() proto.Message {
	cmd, _ := ptypes.MarshalAny(c.Command)
	return &pb.Context{
		ServiceName: c.ServiceName,
		ServiceVersion: c.ServiceVersion,
		Command: cmd,
		AccessToken: c.AccessToken,
		PrincipalId: c.PrincipalId,
		PrincipalRole: c.PrincipalRole,
	}
}
