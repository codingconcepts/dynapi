package dynapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/codingconcepts/dynapi/test"
)

const (
	testBuildVersion   = "1.2.3"
	testBuildTimestamp = "2017-06-09 07:02:31.187"
)

var (
	server *Server

	bodyGETRouteExample = RouteConfig{
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
		URI:        "/person/:name/:age",
		Example:    "/person/Rob/30",
		Body:       "Name: {{.name}}, Age: {{.age}}",
	}

	bodyPOSTRouteExample = RouteConfig{
		Method:     http.MethodPost,
		StatusCode: http.StatusCreated,
		URI:        "/stock",
		Example:    "/stock",
	}

	durationRouteExample = RouteConfig{
		Method:      http.MethodGet,
		StatusCode:  http.StatusTeapot,
		URI:         "/wait/:timeout",
		Example:     "/wait/1s",
		DurationArg: "timeout",
	}

	routeConfig = RouteConfigs{
		bodyGETRouteExample,
		bodyPOSTRouteExample,
		durationRouteExample,
	}
)

func TestMain(t *testing.M) {
	server = NewServer("host", 1234, BuildInfo(testBuildVersion, testBuildTimestamp), Routes(routeConfig...))

	os.Exit(t.Run())
}

func TestVersion(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/version", nil)
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	test.Equals(t, fmt.Sprintf("%s\n%s", testBuildVersion, testBuildTimestamp), resp.Body.String())
	test.Equals(t, http.StatusOK, resp.Result().StatusCode)
}

func TestGetConfig(t *testing.T) {
	req, _ := http.NewRequest(http.MethodOptions, "/config", nil)
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	test.Equals(t, http.StatusOK, resp.Result().StatusCode)

	var configs RouteConfigs
	err := json.NewDecoder(resp.Body).Decode(&configs)
	test.ErrorNil(t, err)

	test.Equals(t, len(routeConfig), len(configs))
}

func TestAddConfig(t *testing.T) {
	config := RouteConfig{
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
		URI:        "/things/:id",
		Example:    "/things/19a88b749f",
		Body:       "successfully deleted {{.id}}",
	}

	// add route to server
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(config)

	req, _ := http.NewRequest(http.MethodPost, "/config", buf)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	test.Equals(t, http.StatusOK, resp.Result().StatusCode)
	test.Equals(t, "successfully added route", resp.Body.String())

	// invoke new route
	req, _ = http.NewRequest(http.MethodGet, "/things/19a88b749f", nil)
	resp = httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	test.Equals(t, http.StatusOK, resp.Result().StatusCode)
	test.Equals(t, "successfully deleted 19a88b749f", resp.Body.String())
}

func TestBody(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/person/Rob/30", nil)
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	test.Equals(t, "Name: Rob, Age: 30", resp.Body.String())
	test.Equals(t, http.StatusOK, resp.Result().StatusCode)
}

func TestRouteHandlerOptions(t *testing.T) {
	testCases := []struct {
		config       RouteConfig
		expectedBody string
	}{
		{config: bodyGETRouteExample, expectedBody: "Name: Rob, Age: 30"},
		{config: bodyPOSTRouteExample, expectedBody: ""},
		{config: durationRouteExample, expectedBody: ""},
	}
	for _, testCase := range testCases {
		// TODO: DOESN'T CURRENTLY RUN DURATION TESTS DUE TO CLOCK ISSUE, UNCOMMENT WHEN WORKING
		if testCase.config.DurationArg != "" {
			continue
		}

		t.Run(testCase.config.Example, func(t *testing.T) {
			req, _ := http.NewRequest(testCase.config.Method, testCase.config.Example, nil)
			resp := httptest.NewRecorder()
			server.ServeHTTP(resp, req)

			test.Equals(t, testCase.config.StatusCode, resp.Result().StatusCode)
			test.Equals(t, testCase.expectedBody, resp.Body.String())
		})
	}
}

func TestDuration(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/wait/1ms", nil)
	resp := httptest.NewRecorder()

	start := time.Now()

	server.ServeHTTP(resp, req)

	stop := time.Now()
	duration := stop.Sub(start)

	test.Equals(t, http.StatusTeapot, resp.Result().StatusCode)
	test.Assert(t, duration >= time.Millisecond)
}
