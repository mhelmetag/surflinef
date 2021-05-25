package surflinef

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func setupFixtureServer(fixtureFilename string) (*httptest.Server, error) {
	d, err := ioutil.ReadFile(fixtureFilename)
	if err != nil {
		return nil, err
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))

	return ts, nil
}
