package surflinef

import (
	"net/url"
	"testing"
)

func TestTidesResponse(t *testing.T) {
	s, err := setupFixtureServer("fixtures/tides.json")
	if err != nil {
		t.Fatal(err)
	}

	defer s.Close()

	bu, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	tq := TidesQuery{
		SpotID: "5842041f4e65fad6a7708814",
		Days:   6,
	}

	tr, err := c.GetTides(tq)
	if err != nil {
		t.Fatal(err)
	}

	ts := tr.Data.Tides

	eT := "NORMAL"
	aT := ts[0].Type
	if aT != eT {
		t.Errorf("Got '%s', expected '%s'", aT, eT)
	}

	eH := 2.33
	aH := ts[0].Height
	if aH != eH {
		t.Errorf("Got '%f', expected '%f'", aH, eH)
	}
}
