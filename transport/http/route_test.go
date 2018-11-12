package http

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/arcus/pkg/service"
	testpb "github.com/arcus/pkg/transport/http/internal/test"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func TestRouteUseForHTTP(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	r := g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)
	buf := new(bytes.Buffer)
	r.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})

	c, b := request(s, http.MethodGet, "/group/msg" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
	if buf.String() != "1" {
		t.Errorf("unexpected middleware output: %s\n", buf.String())
	}
}

func TestRouteMiddleware(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	buf := new(bytes.Buffer)
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler, func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})

	c, b := request(s, http.MethodGet, "/group/msg" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
	if buf.String() != "1" {
		t.Errorf("unexpected middleware output: %s\n", buf.String())
	}
}

func TestRouteUse(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	r := g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)
	buf := new(bytes.Buffer)
	r.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})

	c, b := request(s, http.MethodGet, "/group/msg" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
	if buf.String() != "1" {
		t.Errorf("unexpected middleware output: %s\n", buf.String())
	}
}

func TestRouteUseBoth(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	r := g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)
	buf := new(bytes.Buffer)
	r.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})
	r.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("2")
			return next(ctx, m)
		}
	})

	c, b := request(s, http.MethodGet, "/group/msg" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
	if buf.String() != "12" {
		t.Errorf("unexpected middleware output: %s\n", buf.String())
	}
}
