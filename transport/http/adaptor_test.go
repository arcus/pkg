package http

import (
	"net/http"
	"testing"

	"github.com/arcus/pkg/service"
	"github.com/arcus/pkg/status"
	testpb "github.com/arcus/pkg/transport/http/internal/test"
	"github.com/labstack/echo"
)

func customAdaptor(next service.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, ok := c.Request().Context().(*service.Context)
		if !ok {
			ctx = service.WithContext(c.Request().Context())
		}

		msg := &testpb.Message{}
		if err := c.Bind(msg); err != nil {
			return err
		}

		r, err := next(ctx, msg)
		if err != nil {
			return err
		}

		rep, ok := r.(*testpb.Message)
		if !ok {
			return status.New(status.Internal, "unknown return type")
		}

		return c.String(http.StatusOK, rep.Message)
	}
}

func TestCustomAdaptor(t *testing.T) {
	s := NewServer()
	s.Query("/", customAdaptor, EchoHandler)

	c, b := request(s, http.MethodGet, "/" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != "ok" {
		t.Errorf("unexpected response: %s\n", b)
	}
}
