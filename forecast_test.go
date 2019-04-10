package surflinef

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAnalysis(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/forecast.json")
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
		Resources: []string{"surf,analysis,wind,weather,tide"},
		Days:      1,
		Units:     "e",
	}

	qs, err := q.QueryString()
	if err != nil {
		t.Fatal(err)
	}

	f, err := c.GetForecast("2141", qs)
	if err != nil {
		t.Fatal(err)
	}

	a := f.Analysis

	cond := "FAIR"
	if a.GeneralCondition[0] != cond {
		t.Errorf("Got '%s', expected '%s'", a.GeneralCondition[0], cond)
	}

	min := "0"
	if a.SurfMin[0] != min {
		t.Errorf("Got '%s', expected '%s'", a.SurfMin[0], min)
	}
}
