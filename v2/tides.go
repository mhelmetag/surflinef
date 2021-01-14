package surflinef

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TideQuery is used to build Tide query params
type TideQuery struct {
	SpotID string `url:"spotId"`
	Days   int    `url:"days"`
}

// TidesResponse is the root JSON struct for tide data
type TidesResponse struct {
	Associated Associated `json:"associated"`
	Data       tidesData  `json:"data"`
}

type tidesData struct {
	Tides []Tide `json:"tides"`
}

// Tide is the JSON struct for a tide reading
type Tide struct {
	Timestamp int     `json:"timestamp"`
	Type      string  `json:"type"`
	Height    float64 `json:"height"`
}

// GetTides fetches a TidesResponse from the API
func (c *Client) GetTides(tq TideQuery) (TidesResponse, error) {
	vs, err := query.Values(tq)

	if err != nil {
		return TidesResponse{}, err
	}

	qs := vs.Encode()

	s := c.FullURL(qs)
	u, err := url.Parse(s)
	if err != nil {
		return TidesResponse{}, err
	}

	r, err := c.Get(u)
	if err != nil {
		return TidesResponse{}, err
	}

	defer r.Body.Close()

	var tr TidesResponse
	err = json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		return TidesResponse{}, err
	}

	return tr, nil
}
