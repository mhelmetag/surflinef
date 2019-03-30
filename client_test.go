package surflinef

import (
  "net/url"
  "testing"
)

func TestDefaultClient(t *testing.T) {
	bu, err := url.Parse("http://api.surfline.com/v1/forecasts")
	if err != nil {
		t.Fatal(err)
	}
	c, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	if c.BaseURL.String() != bu.String() {
		t.Errorf("Got '%s', expected '%s'", c.BaseURL.String(), bu.String())
	}
}
