package surflinef

import (
	"fmt"
	"net/http"
	"net/url"
)

// Client is the SurflineF HTTP Client.
type Client struct {
	BaseURL *url.URL

	httpClient *http.Client
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
