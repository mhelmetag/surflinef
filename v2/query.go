package surflinef

import (
	"github.com/google/go-querystring/query"
)

// Query is used to build Forecast query params.
// False, nil or empty values are ignored.
type Query struct {
	SubregionID string `url:"subregionId,omitempty"`
	Days        int    `url:"days,omitempty"`
}

// QueryString builds a query string from a Query.
func (q *Query) QueryString() (string, error) {
	vs, err := query.Values(q)

	if err != nil {
		return "", err
	}

	return vs.Encode(), nil
}
