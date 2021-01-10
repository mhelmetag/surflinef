package surflinef

import (
	"encoding/json"
	"net/url"
)

// ConditionsResponse is the root JSON struct for condition data
type ConditionsResponse struct {
	Associated Associated    `json:"associated"`
	Data       conditonsData `json:"data"`
}

type conditonsData struct {
	Conditions []Condition `json:"conditions"`
}

// Condition is the JSON struct for a daily condition
type Condition struct {
	Timestamp   int        `json:"timestamp"`
	Forecaster  Forecaster `json:"forecaster"`
	Human       bool       `json:"human"`
	Observation string     `json:"observation"`
	AM          Report     `json:"am"`
	PM          Report     `json:"pm"`
}

// Forecaster is the JSON struct for the forecaster
type Forecaster struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// Report is the JSON struct of a report (for the AM or PM)
type Report struct {
	MaxHeight     float64 `json:"maxHeight"`
	MinHeight     float64 `json:"minHeight"`
	HumanRelation string  `json:"humanRelation"`
	Rating        string  `json:"rating"`
}

// GetConditions fetches a ConditionsResponse from the API
func (c *Client) GetConditions(qs string) (ConditionsResponse, error) {
	s := c.FullURL(qs)
	u, err := url.Parse(s)
	if err != nil {
		return ConditionsResponse{}, err
	}

	r, err := c.Get(u)
	if err != nil {
		return ConditionsResponse{}, err
	}

	defer r.Body.Close()

	var cr ConditionsResponse
	err = json.NewDecoder(r.Body).Decode(&cr)
	if err != nil {
		return ConditionsResponse{}, err
	}

	return cr, nil
}
