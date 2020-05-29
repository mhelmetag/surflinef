package surflinef

import (
	"net/url"
	"testing"
)

func TestClient(t *testing.T) {
	bu, err := url.Parse("https://services.surfline.com/kbyg/regions/forecasts/conditions")
	if err != nil {
		t.Fatal(err)
	}

	c := Client{BaseURL: bu}

	if c.BaseURL.String() != bu.String() {
		t.Errorf("Got '%s', expected '%s'", c.BaseURL.String(), bu.String())
	}
}
