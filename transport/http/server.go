package http

import (
	"github.com/arcus/pkg/status"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewErrorHandler(mux *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		// Already an HTTP error, handling as normal.
		if _, ok := err.(*echo.HTTPError); ok {
			mux.DefaultHTTPErrorHandler(err, c)
			return
		}

		// Otherwise convert it into a service status error and
		s := status.Convert(err)
		c.JSON(s.HTTPCode(), map[string]string{
			"code":    s.Code().String(),
			"message": s.Message(),
		})
	}
}

func ServiceContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := ContextFromRequest(c.Request())
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// NewServer returns a new Echo server that can bind protobuf messages, converts
// errors to `arcus.pkg.status.Status`s, logs requests as JSON, and converts
// request contexts to the custom `arcus.pkg.service.Context`.
func NewServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Binder = &ProtoBinder{}
	e.HTTPErrorHandler = NewErrorHandler(e)
	e.Use(
		middleware.Logger(),
		ServiceContextMiddleware,
	)
	return e
}
