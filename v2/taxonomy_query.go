package surflinef

import (
	"github.com/google/go-querystring/query"
)

// TaxonomyQuery is used to build Forecast query params
// False, nil or empty values are ignored
type TaxonomyQuery struct {
	ID       string `url:"id"`
	MaxDepth int    `url:"maxDepth"`
	Type     string `url:"type"`
}

// TaxonomyQueryString builds a query string from a Query
func (q *TaxonomyQuery) TaxonomyQueryString() (string, error) {
	vs, err := query.Values(q)

	if err != nil {
		return "", err
	}

	return vs.Encode(), nil
}
