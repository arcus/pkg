package http

import (
	"github.com/arcus/pkg/service"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

// Route is an endpoint that routes traffic to a handler.
type Route struct {
	// adaptor is the function used to convert the service.Handler
	// into an echo.HandlerFunc.
	adaptor Adaptor

	method  string
	path    string
	handler service.Handler
	reqType proto.Message

	httpMiddleware []echo.MiddlewareFunc
	middleware     []service.Middleware

	echo *echo.Echo
}

// Use wraps the handler in middleware, which will be applied last first.
func (r *Route) Use(m ...service.Middleware) {
	r.middleware = append(r.middleware, m...)
	r.build()
}

// UseForHTTP uses middleware at the transport layer, applied last first.
func (r *Route) UseForHTTP(m ...echo.MiddlewareFunc) {
	r.httpMiddleware = append(r.httpMiddleware, m...)
	r.build()
}

// build creates an echo.Route from the registered middleware and
// registers it on the underlying echo.Echo instance at the configured
// and path, using the configured Adaptor.
func (r *Route) build() {
	h := r.handler
	for i := len(r.middleware) - 1; i >= 0; i-- {
		h = r.middleware[i](h)
	}
	r.echo.Add(r.method, r.path, r.adaptor(h), r.httpMiddleware...)
}
