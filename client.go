package surflinef

import (
	"net/http"
	"net/url"
)

// Client is the SurflineF HTTP Client.
type Client struct {
	BaseURL *url.URL

	httpClient *http.Client
}

// DefaultClient returns a default configured SurflineF Client.
func DefaultClient() (*Client, error) {
	u, err := url.Parse("http://api.surfline.com/v1/forecasts")
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	client := Client{BaseURL: u, httpClient: httpClient}

	return &client, err
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
