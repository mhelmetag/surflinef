package surflinef

import (
	"fmt"
	"net/http"
	"net/url"
)

// ConditionsBaseURL is the base URL for the regional conditions service
const ConditionsBaseURL = "https://services.surfline.com/kbyg/regions/forecasts/conditions"

// TidesBaseURL is the base URL for the spot tides service
const TidesBaseURL = "https://services.surfline.com/kbyg/spots/forecasts/tides"

// TaxonomyBaseURL is the base URL for the taxonomy service
const TaxonomyBaseURL = "https://services.surfline.com/taxonomy"

// WaveBaseURL is the base URL for the wave/swell service
const WaveBaseURL = "https://services.surfline.com/kbyg/spots/forecasts/wave"

// Client is the SurflineF HTTP Client.
type Client struct {
	BaseURL *url.URL
}

// FullURL formats the query string and Client BaseUrl
func (c *Client) FullURL(qs string) string {
	return fmt.Sprintf("%s?%s", c.BaseURL, qs)
}

// Get is just a wrapper for http.Get
func (c *Client) Get(u *url.URL) (*http.Response, error) {
	r, err := http.Get(u.String())

	return r, err
}
