package main

import (
	"net/http"

	"github.com/codingconcepts/dynapi"
)

var (
	buildVersion   string
	buildTimestamp string
)

func main() {
	server := dynapi.NewServer(buildVersion, buildTimestamp, configuration...)
	server.Start(":1234")
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
