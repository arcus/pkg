package http

import (
	"net/http"
	"net/url"

	"github.com/arcus/pkg/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Group is a group of endpoints that use a common prefix and/or set of middleware.
type Group struct {
	prefix string

	httpMiddleware []echo.MiddlewareFunc
	middleware     []service.Middleware

	echo *echo.Echo
}

// Use wraps the handler in middleware, which will be applied last first.
func (g *Group) Use(m ...service.Middleware) {
	// TODO: echo implements this such that all subroutes of the Group
	// will pass through the middleware, even if there is no route defined.
	// See: https://github.com/labstack/echo/blob/master/group.go#L20
	g.middleware = append(g.middleware, m...)
}

// UseForHTTP uses middleware at the transport layer, applied last first.
func (g *Group) UseForHTTP(m ...echo.MiddlewareFunc) {
	g.httpMiddleware = append(g.httpMiddleware, m...)
}

// Query creates an endpoint for query operations at path that
// handles the request type with the handler wrapped in middleware.
func (g *Group) Query(path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	return g.add(http.MethodGet, path, a, h, m...)
}

// Command creates an endpoint for command operations at path that
// handles the request type with the handler wrapped in middleware.
func (g *Group) Command(path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	return g.add(http.MethodPost, path, a, h, m...)
}

// Group creates a group at path using middleware.
func (g *Group) Group(path string, m ...service.Middleware) *Group {
	nm := make([]service.Middleware, 0, len(g.middleware)+len(m))
	nm = append(nm, g.middleware...)
	nm = append(nm, m...)
	return &Group{
		prefix:         g.prefix + path,
		httpMiddleware: g.httpMiddleware,
		middleware:     nm,
		echo:           g.echo,
	}
}

// Proxy proxies all requests with path prefix to uris, removing
// path before relay.
func (g *Group) Proxy(path string, uris ...string) error {
	var ts []*middleware.ProxyTarget
	for _, uri := range uris {
		u, err := url.Parse(uri)
		if err != nil {
			return err
		}
		ts = append(ts, &middleware.ProxyTarget{URL: u})
	}

	m := append(g.httpMiddleware, middleware.ProxyWithConfig(
		middleware.ProxyConfig{
			Balancer: middleware.NewRoundRobinBalancer(ts),
			Rewrite:  map[string]string{g.prefix + path + "/*": "/$1"},
		},
	))

	g.echo.Group(g.prefix+path, m...)
	return nil
}

func (g *Group) add(method, path string, a Adaptor, h service.Handler, m ...service.Middleware) *Route {
	// Combine into a new slice to avoid accidentally passing the same slice for
	// multiple routes, which would lead to later add() calls overwriting the
	// middleware from earlier calls.
	nm := make([]service.Middleware, 0, len(g.middleware)+len(m))
	nm = append(nm, g.middleware...)
	nm = append(nm, m...)
	r := &Route{
		adaptor:        a,
		method:         method,
		path:           g.prefix + path,
		handler:        h,
		httpMiddleware: g.httpMiddleware,
		middleware:     nm,
		echo:           g.echo,
	}
	r.build()
	return r
}
