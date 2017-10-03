package dynapi

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server holds the routes available for extending the
// API surface.
type Server struct {
	router         *echo.Echo
	certsDir       string
	port           int
	routes         RouteConfigs
	buildVersion   string
	buildTimestamp string
}

// NewServer returns a pointer to a new instance of Server.
func NewServer(options ...Option) (s *Server) {
	s = &Server{}

	router := echo.New()
	router.Use(middleware.Recover())

	router.GET("/version", s.GetVersion)
	router.OPTIONS("/config", s.GetConfig)
	router.POST("/config", s.AddRoute)
	s.router = router

	for _, option := range options {
		if err := option(s); err != nil {
			log.Fatal(err)
		}
	}

	return
}

// ServeHTTP serves from the server's router.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Start start the server's router.
func (s *Server) Start() (err error) {
	return s.router.Start(fmt.Sprintf(":%d", s.port))
}

// Stop stops the server's router.
func (s *Server) Stop() (err error) {
	return s.router.Close()
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

// GetVersion displays the build version and timestamp
// of the server.
func (s *Server) GetVersion(c echo.Context) (err error) {
	return c.String(http.StatusOK, fmt.Sprintf("%s\n%s", s.buildVersion, s.buildTimestamp))
}

// GetConfig displays the available routes.
func (s *Server) GetConfig(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, s.routes)
}

func (s *Server) add(route RouteConfig) {
	// if a body template has been provided, parse it up-front
	// to make calls to the dynamic endpoint quicker.
	if route.Body != "" {
		route.BodyTemplate = template.Must(template.New(route.Body).Parse(route.Body))
	}

	handler := s.routeHandler(route)
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

func (s *Server) routeHandler(r RouteConfig) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		body := ParseArgs(c)

		if err = s.sleep(body, r, c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		// if a body hasn't be configured, don't bother continuing
		if r.Body == "" {
			return c.String(r.StatusCode, "")
		}

		if r.BodyContentType != "" {
			c.Response().Header().Set("Content-Type", r.BodyContentType)
		}

		c.Response().WriteHeader(r.StatusCode)
		template, err := r.BodyTemplate.Parse(r.Body)
		if err = template.Execute(c.Response().Writer, body); err != nil {
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

func (s *Server) sleep(args map[string]interface{}, r RouteConfig, c echo.Context) (err error) {
	if r.DurationArg == "" {
		return
	}

	if rawDuration := args[r.DurationArg]; rawDuration != "" {
		var duration time.Duration
		if duration, err = time.ParseDuration(rawDuration.(string)); err != nil {
			return
		}

		time.Sleep(duration)
	}

	return
}
