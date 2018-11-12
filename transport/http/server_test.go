package http

import (
	"bytes"

	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arcus/pkg/service"
	testpb "github.com/arcus/pkg/transport/http/internal/test"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

const (
	msgJSON = `{"message":"ok"}`
	msgParams = "?message=ok"
)

func EchoHandler(ctx *service.Context, m proto.Message) (proto.Message, error) {
	return m, nil
}

func TestQuery(t *testing.T) {
	s := NewServer()
	s.Query("/", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodGet, "/" + msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func TestCommand(t *testing.T) {
	s := NewServer()
	s.Command("/", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodPost, "/", echo.MIMEApplicationJSON, strings.NewReader(msgJSON))
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func TestUseForHTTP(t *testing.T) {
	s := NewServer()
	buf := new(bytes.Buffer)
	s.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})
	s.Query("/", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodGet, "/" + msgParams, "", nil)
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

func TestUse(t *testing.T) {
	s := NewServer()
	buf := new(bytes.Buffer)
	s.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("1")
			return next(ctx, m)
		}
	})
	s.Query("/", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodGet, "/" + msgParams, "", nil)
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

func TestUseBoth(t *testing.T) {
	s := NewServer()
	buf := new(bytes.Buffer)
	s.UseForHTTP(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			buf.WriteString("1")
			return next(c)
		}
	})
	s.Use(func(next service.Handler) service.Handler {
		return func(ctx *service.Context, m proto.Message) (proto.Message, error) {
			buf.WriteString("2")
			return next(ctx, m)
		}
	})
	s.Query("/", MakeDefaultAdaptor(&testpb.Message{}), EchoHandler)

	c, b := request(s, http.MethodGet, "/" + msgParams, "", nil)
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

func request(s *Server, method, path, contentType string, body io.Reader) (int, string) {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}
