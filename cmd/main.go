package main

import (
	"net/http"
	"time"

	"github.com/codingconcepts/dynapi"
)

func main() {
	server := dynapi.NewServer(configuration...)
	server.Start(":1234")
}

var configuration = dynapi.RouteConfigs{
	dynapi.RouteConfig{
		Method:       http.MethodGet,
		URI:          "/person/:name/:age",
		Example:      "/person/Rob/30",
		StatusCode:   http.StatusOK,
		BodyTemplate: "Name: {{.name}} Age: {{.age}}",
	},
	dynapi.RouteConfig{
		Method:     http.MethodGet,
		URI:        "/timeout",
		Example:    "/timeout",
		StatusCode: http.StatusTeapot,
		Duration:   time.Second * 10,
	},
}
