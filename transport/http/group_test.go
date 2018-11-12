package http

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/arcus/pkg/service"
	testpb "github.com/arcus/pkg/transport/http/internal/test"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func TestGroupQuery(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodGet, "/group/msg" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func TestGroupCommand(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	g.Command("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodPost, "/group/msg", echo.MIMEApplicationJSON, strings.NewReader(msgJSON))
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func TestGroupUseForHTTP(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	buf := new(bytes.Buffer)
	g.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

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

func TestGroupMiddleware(t *testing.T) {
	s := NewServer()
	buf := new(bytes.Buffer)
	g := s.Group("/group", func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

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

func TestGroupUse(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	buf := new(bytes.Buffer)
	g.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

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

func TestGroupUseBoth(t *testing.T) {
	s := NewServer()
	g := s.Group("/group")
	buf := new(bytes.Buffer)
	g.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})
	g.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("2")
			return next(ctx, m)
		}
	})
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

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

func TestGroupWithServerMiddleware(t *testing.T) {
	s := NewServer()
	buf := new(bytes.Buffer)
	s.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})
	g := s.Group("/group")
	g.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("2")
			return next(ctx, m)
		}
	})
	g.Query("/msg", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

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
