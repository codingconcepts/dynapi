package dynapi

import (
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

// Server holds the routes available for extending the
// API surface.
type Server struct {
	router *echo.Echo
	routes RouteConfigs
}

// NewServer returns a pointer to a new instance of Server.
func NewServer(routes ...RouteConfig) (s *Server) {
	s = &Server{}

	router := echo.New()
	router.GET("/config", s.GetRoute)
	router.POST("/config", s.AddRoute)
	s.router = router

	for _, route := range routes {
		s.add(route)
	}

	return
}

// Start start the server's router.
func (s *Server) Start(addr string) (err error) {
	return s.router.Start(addr)
}

// AddRoute allows a user to add a new API route.
func (s *Server) AddRoute(c echo.Context) (err error) {
	var route RouteConfig
	if err = c.Bind(&route); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	s.add(route)
	return c.String(http.StatusOK, "successfully added route")
}

// GetRoute displays the available routes.
func (s *Server) GetRoute(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, s.routes)
}

func (s *Server) add(route RouteConfig) {
	handler := routeHandler(route)
	handlerOptions := routeHandlerOptions(route)

	switch route.Method {
	case http.MethodGet:
		s.router.GET(route.URI, handler)
	case http.MethodPost:
		s.router.POST(route.URI, handler)
	}

	s.router.OPTIONS(route.URI, handlerOptions)

	// keep track of the route within the server
	s.routes.Merge(route)
}

func routeHandler(r RouteConfig) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		time.Sleep(r.Duration)

		if r.BodyTemplate == "" {
			return c.String(r.StatusCode, "")
		}

		body := ParseArgs(c)

		t := template.Must(template.New(r.URI).Parse(r.BodyTemplate))
		if err = t.Execute(c.Response().Writer, body); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return
	}
}

func routeHandlerOptions(r RouteConfig) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, r)
	}
}
