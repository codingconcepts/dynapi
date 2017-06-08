package dynapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/codingconcepts/dynapi/test"
	"github.com/labstack/echo"
)

func TestParseArgs(t *testing.T) {
	query := url.Values{
		"b": []string{"3"},
		"d": []string{"4", "5"},
	}

	req := httptest.NewRequest(http.MethodGet, "/anything?"+query.Encode(), nil)
	resp := httptest.NewRecorder()

	router := echo.New()
	context := router.NewContext(req, resp)
	context.SetParamNames("a", "b")
	context.SetParamValues("1", "2")

	args := ParseArgs(context)
	test.Equals(t, "1", args["a"])
	test.Equals(t, []string{"3"}, args["b"])
	test.Equals(t, []string{"4", "5"}, args["d"])
}
