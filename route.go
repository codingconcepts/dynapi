package dynapi

import (
	"text/template"
)

// RouteConfig holds the information about a dynamic API route.
type RouteConfig struct {
	Method       string             `json:"method"`
	URI          string             `json:"uri"`
	Example      string             `json:"example"`
	StatusCode   int                `json:"statusCode"`
	DurationArg  string             `json:"durationArg,omitempty"`
	Body         string             `json:"body,omitempty"`
	BodyTemplate *template.Template `json:"-"`
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
		if other == existing {
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
	if r.DurationArg != other.DurationArg {
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
