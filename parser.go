package dynapi

import "github.com/labstack/echo"

// ParseArgs extracts arguments from the URL and the query string
// and ensures that any query string parameters provided overwrite
// the URL arguments.
func ParseArgs(c echo.Context) (args map[string]interface{}) {
	args = parseQueryParamArgs(c)
	queryArgs := parseQueryStringArgs(c)

	for key, value := range queryArgs {
		args[key] = value
	}

	return
}

// parseQueryParamArgs extracts parameter key value pairs from a
// given context and returns them in a map.
func parseQueryParamArgs(c echo.Context) (items map[string]interface{}) {
	items = make(map[string]interface{})

	if len(c.ParamNames()) == 0 {
		return
	}

	items = make(map[string]interface{}, len(c.ParamNames()))
	for _, name := range c.ParamNames() {
		items[name] = c.Param(name)
	}

	return
}

// parseQueryStringArgs extracts query string key value pairs from a
// given context and returns them in a map.
func parseQueryStringArgs(c echo.Context) (items map[string]interface{}) {
	items = make(map[string]interface{})

	if len(c.QueryParams()) == 0 {
		return
	}

	items = make(map[string]interface{}, len(c.QueryParams()))
	for name, value := range c.QueryParams() {
		items[name] = value
	}

	return
}
