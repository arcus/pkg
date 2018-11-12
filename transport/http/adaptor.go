package http

import (
	"bytes"
	"net/http"
	"reflect"

	"github.com/arcus/pkg/service"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo"
)

// Adaptor is a function that converts a service.Handler for a request type
// into an echo.HandlerFunc.
type Adaptor func(next service.Handler) echo.HandlerFunc

// MakeDefaultAdaptor returns an Adaptor that decodes the request into the
// proto.Message request type and passes it to the handler, returning any error
// that results, an HTTP OK with no body if the response is empty, or the JSON
// encoded response as the body of an HTTP OK response.
func MakeDefaultAdaptor(req proto.Message) Adaptor {
	return func(next service.Handler) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.Request().Context().(*service.Context)
			if !ok {
				ctx = service.WithContext(c.Request().Context())
			}

			// Make a copy.
			req = copyProto(req)

			// Decode the request body into the request type.
			// Note: this assumes that the custom protoBinder has been
			// registered with the echo instance.
			if err := c.Bind(req); err != nil {
				return err
			}

			rep, err := next(ctx, req)
			if err != nil {
				return err
			}

			// Write an empty response.
			if _, ok := rep.(*emptypb.Empty); ok || rep == nil {
				return c.NoContent(http.StatusOK)
			}

			// Write the response as a JSON-encoded value.
			// TODO: this is wasting buffers..
			var buf bytes.Buffer
			m := &jsonpb.Marshaler{
				OrigName: true,
			}
			if err := m.Marshal(&buf, rep); err != nil {
				return err
			}
			return c.JSONBlob(http.StatusOK, buf.Bytes())
		}
	}
}

func copyProto(src proto.Message) proto.Message {
	in := reflect.ValueOf(src)
	if in.IsNil() {
		return src
	}
	out := reflect.New(in.Type().Elem())
	dst := out.Interface().(proto.Message)
	return dst
}
