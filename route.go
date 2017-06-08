package dynapi

import "time"

// RouteConfig holds the information about a dynamic API route.
type RouteConfig struct {
	Method       string        `json:"method"`
	URI          string        `json:"uri"`
	Example      string        `json:"example"`
	BodyTemplate string        `json:"bodyTemplate"`
	StatusCode   int           `json:"statusCode"`
	Duration     time.Duration `json:"duration"`
}

// RouteConfigs is a slice of RouteConfig structs.
type RouteConfigs []RouteConfig

// Merge adds unique RouteConfig structs into the slice of
// RouteConfigs, ensuring there are no route collisions.
func (r RouteConfigs) Merge(others ...RouteConfig) {
	for _, other := range others {
		r.MergeRoute(other)
	}
}

// MergeRoute adds single unique RouteConfig struct to the
// slice of RouteConfigs, ensuring there are no route collisions.
func (r RouteConfigs) MergeRoute(other RouteConfig) {
	for _, existing := range r {
		if other == existing {
			return
		}
	}

	r = append(r, other)
}
