package dynapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/codingconcepts/dynapi/test"
	"github.com/labstack/echo"
)

func TestParseArgs(t *testing.T) {
	testCases := []struct {
		query       url.Values
		paramNames  []string
		paramValues []string
		expected    map[string]interface{}
	}{
		{
			query:       url.Values{"b": []string{"3"}, "d": []string{"4", "5"}},
			paramNames:  []string{"a", "b"},
			paramValues: []string{"1", "2"},
			expected: map[string]interface{}{
				"a": "1",
				"b": []string{"3"},
				"d": []string{"4", "5"},
			},
		},
		{
			query:       url.Values{"c": []string{"3"}, "d": []string{"4", "5"}},
			paramNames:  []string{"a", "b"},
			paramValues: []string{"1", "2"},
			expected: map[string]interface{}{
				"a": "1",
				"b": "2",
				"c": []string{"3"},
				"d": []string{"4", "5"},
			},
		},
		{
			query: url.Values{"a": []string{"1"}},
			expected: map[string]interface{}{
				"a": []string{"1"},
			},
		},
		{
			paramNames:  []string{"a", "b"},
			paramValues: []string{"1", "2"},
			expected: map[string]interface{}{
				"a": "1",
				"b": "2",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.expected), func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/anything?"+testCase.query.Encode(), nil)
			resp := httptest.NewRecorder()

			router := echo.New()
			context := router.NewContext(req, resp)

			context.SetParamNames(testCase.paramNames...)
			context.SetParamValues(testCase.paramValues...)

			args := ParseArgs(context)

			test.Equals(t, len(testCase.expected), len(args))
			for key, value := range testCase.expected {
				test.Equals(t, value, args[key])
			}
		})
	}
}

// func TestParseArgs(t *testing.T) {
// 	query := url.Values{
// 		"b": []string{"3"},
// 		"d": []string{"4", "5"},
// 	}

// 	req := httptest.NewRequest(http.MethodGet, "/anything?"+query.Encode(), nil)
// 	resp := httptest.NewRecorder()

// 	router := echo.New()
// 	context := router.NewContext(req, resp)
// 	context.SetParamNames("a", "b")
// 	context.SetParamValues("1", "2")

// 	args := ParseArgs(context)
// 	test.Equals(t, "1", args["a"])
// 	test.Equals(t, []string{"3"}, args["b"])
// 	test.Equals(t, []string{"4", "5"}, args["d"])
// }
