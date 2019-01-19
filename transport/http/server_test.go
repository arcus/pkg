package http

import (
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
	msgJSON   = `{"message":"ok"}`
	msgParams = "?message=ok"
)

func EchoHandler(ctx *service.Context, m proto.Message) (proto.Message, error) {
	return m, nil
}

func TestDefaultAdaptor_GET(t *testing.T) {
	s := NewServer()
	s.GET(
		"/",
		DefaultAdaptor(EchoHandler, &testpb.Message{}),
	)

	c, b := request(s, http.MethodGet, "/"+msgParams, "", nil)
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func TestDefaultAdaptor_POST(t *testing.T) {
	s := NewServer()
	s.POST(
		"/",
		DefaultAdaptor(EchoHandler, &testpb.Message{}),
	)

	c, b := request(s, http.MethodPost, "/", echo.MIMEApplicationJSON, strings.NewReader(msgJSON))
	if c != http.StatusOK {
		t.Errorf("unexpected status: %d\n", c)
	}
	if b != msgJSON {
		t.Errorf("unexpected response: %s\n", b)
	}
}

func request(s http.Handler, method, path, contentType string, body io.Reader) (int, string) {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}
