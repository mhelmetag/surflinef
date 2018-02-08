package surflinef

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Forecast is the root JSON struct for forecast data.
type Forecast struct {
	Analysis `json:"Analysis"`
	Tide     `json:"Tide"`
	Surf     `json:"Surf"`
	Weather  `json:"Weather"`
	Wind     `json:"Wind"`
}

// GetForecast grabs a Forecast from surfline with the provided Sub Region ID
// and Query string.
func (c *Client) GetForecast(srID string, query string) (Forecast, error) {
	p := fmt.Sprintf("%s/%s?%s", c.BaseURL, srID, query)
	u, err := url.Parse(p)
	if err != nil {
		return Forecast{}, err
	}

	resp, err := c.get(u)
	if err != nil {
		return Forecast{}, err
	}
	defer resp.Body.Close()

	var forecast Forecast
	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		return Forecast{}, err
	}

	return forecast, nil
}

func (c *Client) get(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
