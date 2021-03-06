package surflinef

import (
	"net/url"
	"testing"
)

func TestWaveResponse(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/wave.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	wq := WaveQuery{
		SpotID: "5842041f4e65fad6a7708814",
		Days:   3,
	}

	cr, err := c.GetWave(wq)
	if err != nil {
		t.Fatal(err)
	}

	ws := cr.Data.Wave

	esmO := 1.8
	asmO := ws[0].Surf.Min
	if asmO != esmO {
		t.Errorf("Got '%f', expected '%f'", asmO, esmO)
	}
}
