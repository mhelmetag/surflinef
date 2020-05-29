package surflinef

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestConditionsResponse(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/conditions.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, httpClient: http.DefaultClient}

	q := Query{
		SubregionID: "58581a836630e24c44878fd4",
		Days:        6,
	}

	qs, err := q.QueryString()
	if err != nil {
		t.Fatal(err)
	}

	cr, err := c.GetConditions(qs)
	if err != nil {
		t.Fatal(err)
	}

	cs := cr.Data.Conditions

	eO := "Steep new NW swell peaks early. Steady NNE winds. "
	aO := cs[0].Observation
	if aO != eO {
		t.Errorf("Got '%s', expected '%s'", aO, eO)
	}

	eM := 3.0
	aM := cs[0].AM.MinHeight
	if aM != eM {
		t.Errorf("Got '%f', expected '%f'", aM, eM)
	}
}
