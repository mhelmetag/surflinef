package surflinef

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// WaveQuery is used to build Wave query params
type WaveQuery struct {
	SpotID        string `url:"spotId"`
	Days          int    `url:"days"`
	IntervalHours int    `url:"intervalHours"`
	MaxHeights    bool   `url:"maxHeights"`
}

// WaveResponse is the root JSON struct for wave data
type WaveResponse struct {
	Associated Associated `json:"associated"`
	Data       waveData   `json:"data"`
}

type waveData struct {
	Wave []Wave `json:"wave"`
}

// Wave is the JSON struct for an interval's swells
type Wave struct {
	Timestamp int     `json:"timestamp"`
	Surf      Surf    `json:"surf"`
	Swells    []Swell `json:"swells"`
}

// Surf is the JSON struct for an interval's nearshore height
// At least... I think it is nearshore height
type Surf struct {
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	OptimalScore int     `json:"optimalScore"`
}

// Swell is the JSON struct for an interval's individual swell
type Swell struct {
	Height       float64 `json:"height"`
	Period       int     `json:"period"`
	Direction    float64 `json:"direction"`
	DirectionMin float64 `json:"directionMin"`
	OptimalScore int     `json:"optimalScore"`
}

// GetWave fetches a WaveResponse from the API
func (c *Client) GetWave(wq WaveQuery) (WaveResponse, error) {
	vs, err := query.Values(wq)

	if err != nil {
		return WaveResponse{}, err
	}

	qs := vs.Encode()

	s := c.FullURL(qs)
	u, err := url.Parse(s)
	if err != nil {
		return WaveResponse{}, err
	}

	r, err := c.Get(u)
	if err != nil {
		return WaveResponse{}, err
	}

	defer r.Body.Close()

	var wr WaveResponse
	err = json.NewDecoder(r.Body).Decode(&wr)
	if err != nil {
		return WaveResponse{}, err
	}

	return wr, nil
}
