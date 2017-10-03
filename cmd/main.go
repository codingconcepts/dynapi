package main

import (
	"log"
	"net/http"

	"github.com/codingconcepts/dynapi"
	"github.com/codingconcepts/env"
)

var (
	buildVersion   string
	buildTimestamp string
)

func main() {
	config := struct {
		Host     string `env:"HOST" required:"true"`
		Port     int    `env:"PORT" required:"true"`
		CertsDir string `env:"CERTS" required:"true" default:"certs"`
	}{}
	if err := env.Set(&config); err != nil {
		log.Fatal(err)
	}

	server := dynapi.NewServer(
		config.Host,
		config.Port,
		dynapi.CertsDir(config.CertsDir),
		dynapi.Routes(configuration...),
		dynapi.BuildInfo(buildVersion, buildTimestamp))

	server.Start()
}

var configuration = dynapi.RouteConfigs{
	dynapi.RouteConfig{
		Method:          http.MethodGet,
		URI:             "/xml",
		Example:         "/xml",
		StatusCode:      http.StatusOK,
		BodyContentType: "text/xml",
		Body: `<?xml version="1.0" encoding="UTF-8"?>
		<movies>
			<movie name="Shutter Island">
				<director firstName="Martin" lastName="Scorsese" />
			</movie>
			<movie name="Kill Bill II">
				<director firstName="Quentin" lastName="Tarantino" />
			</movie>
		</movies>`,
	},
	dynapi.RouteConfig{
		Method:          http.MethodGet,
		URI:             "/jsonNested",
		Example:         "/jsonNested",
		StatusCode:      http.StatusOK,
		BodyContentType: "application/json",
		Body: `{
				"movies": [
					{	"name": "Shutter Island",
						"director": {
							"firstName": "Martin",
							"lastName": "Scorsese"
						}
					},
					{	"name": "Kill Bill II",
						"director": {
							"firstName": "Quentin",
							"lastName": "Tarantino"
						}
					}
			]
		}`,
	},
	dynapi.RouteConfig{
		Method:          http.MethodGet,
		URI:             "/jsonArray",
		Example:         "/jsonArray",
		StatusCode:      http.StatusOK,
		BodyContentType: "application/json",
		Body: `[
			{	"name": "Shutter Island",
				"director": {
					"firstName": "Martin",
					"lastName": "Scorsese"
				}
			},
			{	"name": "Kill Bill II",
				"director": {
					"firstName": "Quentin",
					"lastName": "Tarantino"
				}
			}
		]`,
	},
	dynapi.RouteConfig{
		Method:     http.MethodGet,
		URI:        "/person/:name/:age",
		Example:    "/person/Rob/31",
		StatusCode: http.StatusOK,
		Body:       "Name: {{.name}} Age: {{.age}}",
	},
	dynapi.RouteConfig{
		Method:      http.MethodGet,
		URI:         "/timeout/:duration",
		Example:     "/timeout/1s",
		StatusCode:  http.StatusOK,
		DurationArg: "duration",
	},
}
