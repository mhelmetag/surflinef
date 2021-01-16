package surflinef

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// ConditionsQuery is used to build Conditions query params
type ConditionsQuery struct {
	SubregionID string `url:"subregionId"`
	Days        int    `url:"days"`
}

// ConditionsResponse is the root JSON struct for condition data
type ConditionsResponse struct {
	Associated ConditionsAssociated `json:"associated"`
	Data       conditonsData        `json:"data"`
}

// ConditionsAssociated is associated info to go along with the Conditions API response
// It includes units of measurement and utc offset for timezones
type ConditionsAssociated struct {
	Units     Units `json:"units"`
	UTCOffset int32 `json:"utcOffset"`
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
func (c *Client) GetConditions(cq ConditionsQuery) (ConditionsResponse, error) {
	vs, err := query.Values(cq)

	if err != nil {
		return ConditionsResponse{}, err
	}

	qs := vs.Encode()

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
