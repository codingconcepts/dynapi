package main

import (
	"net/http"

	"github.com/codingconcepts/dynapi"
)

func main() {
	server := dynapi.NewServer(configuration...)
	server.Start(":1234")
}

var configuration = dynapi.RouteConfigs{
	dynapi.RouteConfig{
		Method:     http.MethodGet,
		URI:        "/person/:name/:age",
		Example:    "/person/Rob/30",
		StatusCode: http.StatusOK,
		Body:       "Name: {{.name}} Age: {{.age}}",
	},
	dynapi.RouteConfig{
		Method:      http.MethodGet,
		URI:         "/timeout/:duration",
		Example:     "/timeout/1s",
		StatusCode:  http.StatusTeapot,
		DurationArg: "duration",
	},
}
