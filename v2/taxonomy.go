package surflinef

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TaxonomyQuery is used to build Forecast query params
// False, nil or empty values are ignored
type TaxonomyQuery struct {
	ID       string `url:"id"`
	MaxDepth int    `url:"maxDepth"`
	Type     string `url:"type"`
}

// Taxonomy is the JSON struct for a daily condition
type Taxonomy struct {
	ID        string     `json:"_id"`
	Spot      string     `json:"spot"`
	Subregion string     `json:"subregion"`
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Contains  []Taxonomy `json:"contains"`
}

// GetTaxonomy fetches a Taxonomy from the API
func (c *Client) GetTaxonomy(tq TaxonomyQuery) (Taxonomy, error) {
	vs, err := query.Values(tq)

	if err != nil {
		return Taxonomy{}, err
	}

	qs := vs.Encode()

	s := c.FullURL(qs)
	u, err := url.Parse(s)
	if err != nil {
		return Taxonomy{}, err
	}

	r, err := c.Get(u)
	if err != nil {
		return Taxonomy{}, err
	}

	defer r.Body.Close()

	var t Taxonomy
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return Taxonomy{}, err
	}

	return t, nil
}
