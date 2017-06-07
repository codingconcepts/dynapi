package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var (
	router *echo.Echo

	configuration = routeConfigs{
		routeConfig{http.MethodGet, "/cat", "a cat", http.StatusOK, time.Second},
		routeConfig{http.MethodPost, "/dog", "a dog", http.StatusCreated, time.Second * 2},
	}
)

func main() {
	router = echo.New()
	router.POST("/config", postConfig)
	router.GET("/config", getConfig)

	wireupRoutes(configuration...)

	router.Start(":1234")
}

type routeConfig struct {
	Method     string        `json:"method"`
	URI        string        `json:"uri"`
	Body       string        `json:"body"`
	StatusCode int           `json:"statusCode"`
	Duration   time.Duration `json:"duration"`
}

type routeConfigs []routeConfig

func (r routeConfigs) mergeRoutes(other routeConfigs) {
	for _, newR := range other {
		r.mergeRoute(newR)
	}
}

func (r routeConfigs) mergeRoute(other routeConfig) {
	for _, oldR := range r {
		if other == oldR {
			return
		}
	}
}

func wireupRoutes(routes ...routeConfig) {
	for _, r := range routes {
		copyR := r
		wireupRoute(copyR)
	}
}

func wireupRoute(route routeConfig) {
	handler := func(c echo.Context) (err error) {
		time.Sleep(route.Duration)
		return c.String(route.StatusCode, route.Body)
	}

	switch route.Method {
	case http.MethodGet:
		router.GET(route.URI, handler)
	case http.MethodPost:
		router.POST(route.URI, handler)
	}
}

func handle(r routeConfig) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		time.Sleep(r.Duration)
		return c.String(r.StatusCode, r.Body)
	}
}

func postConfig(c echo.Context) (err error) {
	var config routeConfig
	if err = c.Bind(&config); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	configuration = append(configuration, config)
	return
}

func getConfig(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, configuration)
}
