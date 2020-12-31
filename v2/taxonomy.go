package surflinef

import (
	"encoding/json"
	"net/url"
)

// Taxonomy is the JSON struct for a daily condition
type Taxonomy struct {
	ID       string     `json:"_id"`
	Spot     string     `json:"spot,omitempty"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Contains []Taxonomy `json:"contains,omitempty"`
}

// GetTaxonomy fetches a Taxonomy from the API.
func (c *Client) GetTaxonomy(qs string) (Taxonomy, error) {
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
