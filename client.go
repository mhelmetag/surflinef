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
	url, err := url.Parse("http://api.surfline.com/v1/forecasts")
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	client := Client{BaseURL: url, httpClient: httpClient}

	return &client, err
}
