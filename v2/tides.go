package surflinef

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TidesQuery is used to build Tides query params
type TidesQuery struct {
	SpotID string `url:"spotId"`
	Days   int    `url:"days"`
}

// TidesResponse is the root JSON struct for tide data
type TidesResponse struct {
	Associated TidesAssociated `json:"associated"`
	Data       tidesData       `json:"data"`
}

// TidesAssociated is associated info to go along with the Tide API response
// It includes units of measurement, utc offset for timezones, related locations, etc
type TidesAssociated struct {
	Units        Units    `json:"units"`
	UTCOffset    int32    `json:"utcOffset"`
	TideLocation Location `json:"tideLocation"`
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
func (c *Client) GetTides(tq TidesQuery) (TidesResponse, error) {
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
