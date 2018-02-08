package surflinef

import (
	"github.com/google/go-querystring/query"
)

// Query is used to build Forecast query params.
// False, nil or empty values are ignored.
type Query struct {
	Resources    []string `url:"resources,comma,omitempty"`
	Days         int      `url:"days,omitempty"`
	GetAllSpots  bool     `url:"getAllSpots,omitempty"`
	Aggregate    bool     `url:"aggregate,omitempty"`
	Units        string   `url:"units,omitempty"`
	FullAnalysis bool     `url:"fullAnalysis,omitempty"`
	ShowOptimal  bool     `url:"showOptimal,omitempty"`
	Interpolate  bool     `url:"interpolate,omitempty"`
}

// QueryString builds a query string from a Query.
func (q *Query) QueryString() (string, error) {
	values, err := query.Values(q)

	if err != nil {
		return "", err
	}

	return values.Encode(), nil
}
