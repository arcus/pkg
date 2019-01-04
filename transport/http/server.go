package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/arcus/pkg/service"
	"github.com/arcus/pkg/status"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewServer returns a new server that can bind protobuf messages, converts
// errors to `arcus.pkg.status.Status`s, logs requests as JSON, and converts
// request contexts to the custom `arcus.pkg.arcus.Context`.
func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Binder = &protoBinder{}
	e.HTTPErrorHandler = newErrorHandler(e)
	e.Use(
		middleware.Logger(),
		serviceContextMiddleware,
	)
	return &Server{echo: e}
}

// Server serves traffic using the routes defined on it.
type Server struct {
	httpMiddleware []echo.MiddlewareFunc
	middleware     []service.Middleware

	echo *echo.Echo
}

// Use wraps the handler in middleware, which will be applied last first.
func (s *Server) Use(m ...service.Middleware) {
	s.middleware = append(s.middleware, m...)
}

// UseForHTTP uses middleware at the transport layer, applied last first.
func (s *Server) UseForHTTP(m ...echo.MiddlewareFunc) {
	s.httpMiddleware = append(s.httpMiddleware, m...)
}

// Query creates an endpoint for query operations at path that
// handles the request type with the handler wrapped in middleware.
func (s *Server) Query(path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	return s.add(http.MethodGet, path, a, h, m...)
}

// Command creates an endpoint for command operations at path that
// handles the request type with the handler wrapped in middleware.
func (s *Server) Command(path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	return s.add(http.MethodPost, path, a, h, m...)
}

// Group creates a group at path using middleware.
func (s *Server) Group(path string, m ...service.Middleware) *Group {
	nm := make([]service.Middleware, 0, len(s.middleware)+len(m))
	nm = append(nm, s.middleware...)
	nm = append(nm, m...)
	return &Group{
		prefix:         path,
		httpMiddleware: s.httpMiddleware,
		middleware:     nm,
		echo:           s.echo,
	}
}

// Proxy proxies all requests with path prefix to uris, removing
// path before relay.
func (s *Server) Proxy(path string, uris ...string) error {
	var ts []*middleware.ProxyTarget

	for _, uri := range uris {
		u, err := url.Parse(uri)
		if err != nil {
			return err
		}
		ts = append(ts, &middleware.ProxyTarget{URL: u})
	}

	m := append(s.httpMiddleware, middleware.ProxyWithConfig(
		middleware.ProxyConfig{
			Balancer: middleware.NewRoundRobinBalancer(ts),
			Rewrite:  map[string]string{path + "/*": "/$1"},
		},
	))

	s.echo.Group(path, m...)
	return nil
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.echo.ServeHTTP(w, r)
}

// Start starts the server, listening at addr.
func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

// Shutdowns shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) add(method, path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	// Combine into a new slice to avoid accidentally passing the same slice for
	// multiple routes, which would lead to later add() calls overwriting the
	// middleware from earlier calls.
	nm := make([]service.Middleware, 0, len(s.middleware)+len(m))
	nm = append(nm, s.middleware...)
	nm = append(nm, m...)
	r := &Route{
		adaptor:        a,
		method:         method,
		path:           path,
		handler:        h,
		httpMiddleware: s.httpMiddleware,
		middleware:     nm,
		echo:           s.echo,
	}
	r.build()
	return r
}

func newErrorHandler(mux *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		s := status.Convert(err)
		c.JSON(s.HTTPCode(), map[string]string{
			"code":    s.Code().String(),
			"message": s.Message(),
		})
	}
}

func serviceContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := ContextFromRequest(c.Request())
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
