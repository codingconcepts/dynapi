package dynoapi

import (
	"text/template"
)

// RouteConfig holds the information about a dynamic API route.
type RouteConfig struct {
	Method          string             `json:"method" yaml:"method"`
	URI             string             `json:"uri" yaml:"uri"`
	Example         string             `json:"example" yaml:"example"`
	StatusCode      int                `json:"statusCode" yaml:"statusCode"`
	DurationParam   string             `json:"durationParam,omitempty" yaml:"durationParam"`
	Body            string             `json:"body,omitempty" yaml:"body,omitempty"`
	BodyTemplate    *template.Template `json:"-" yaml:"-"`
	BodyContentType string             `json:"contentType" yaml:"contentType,omitempty"`
}

// RouteConfigs is a slice of RouteConfig structs.
type RouteConfigs []RouteConfig

// Merge adds unique RouteConfig structs into the slice of
// RouteConfigs, ensuring there are no route collisions.
func (r *RouteConfigs) Merge(others ...RouteConfig) {
	for _, other := range others {
		r.MergeRoute(other)
	}
}

// MergeRoute adds single unique RouteConfig struct to the
// slice of RouteConfigs, ensuring there are no route collisions.
func (r *RouteConfigs) MergeRoute(other RouteConfig) {
	for _, existing := range *r {
		if other.Equals(existing) {
			return
		}
	}

	*r = append(*r, other)
}

// Equals performs a field-by-field comparison of two
// RouteConfig structs.  Required due to the template
// that's parsed and stored for the RouteConfig which
// is not comparable.
func (r RouteConfig) Equals(other RouteConfig) bool {
	if r.DurationParam != other.DurationParam {
		return false
	}
	if r.Example != other.Example {
		return false
	}
	if r.Method != other.Method {
		return false
	}
	if r.StatusCode != other.StatusCode {
		return false
	}
	if r.URI != other.URI {
		return false
	}
	if r.Body != other.Body {
		return false
	}

	return true
}
