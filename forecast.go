package surflinef

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Forecast is the root JSON struct for forecast data.
type Forecast struct {
	Analysis
	Tide
	Surf
	Weather
	Wind
}

type iForecast struct {
	iAnalysis `json:"Analysis"`
	Tide      `json:"Tide,omitempty"`
	Surf      `json:"Surf,omitempty"`
	Weather   `json:"Weather,omitempty"`
	Wind      `json:"Wind,omitempty"`
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

// UnmarshalJSON is a custom unmarshaller for Forecast.
func (f *Forecast) UnmarshalJSON(data []byte) error {
	var iF iForecast

	err := json.Unmarshal(data, &iF)
	if err != nil {
		return err
	}

	*f = iF.forecast()

	return nil
}

func (iF iForecast) forecast() Forecast {
	iA := iF.iAnalysis
	a := iA.analysis()

	return Forecast{
		a,
		iF.Tide,
		iF.Surf,
		iF.Weather,
		iF.Wind,
	}
}
