package surflinef

import (
	"encoding/json"
	"net/url"
)

// WaveResponse is the root JSON struct for wave data
type WaveResponse struct {
	Data waveData `json:"data"`
}

type waveData struct {
	Wave []Wave `json:"wave"`
}

type Wave struct {
	Timestamp int     `json:"timestamp"`
	Surf      Surf    `json:"surf"`
	Swells    []Swell `json:"swells"`
}

type Surf struct {
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	OptimalScore int     `json:"optimalScore"`
}

type Swell struct {
	Height       float64 `json:"height"`
	Period       int     `json:"period"`
	Direction    float64 `json:"direction"`
	DirectionMin float64 `json:"directionMin"`
	OptimalScore int     `json:"optimalScore"`
}

// GetWave fetches a WaveResponse from the API
func (c *Client) GetWave(qs string) (WaveResponse, error) {
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
