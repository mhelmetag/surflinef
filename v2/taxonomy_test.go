package surflinef

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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

func TestTaxonomy(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/taxonomy.json")
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, httpClient: http.DefaultClient}

	tq := TaxonomyQuery{
		ID:       "58581a836630e24c44878fd4",
		MaxDepth: 0,
		Type:     "taxonomy",
	}

	tqs, err := tq.TaxonomyQueryString()
	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.GetTaxonomy(tqs)
	if err != nil {
		t.Fatal(err)
	}

	en := "North America"
	an := tx.Name
	if an != en {
		t.Errorf("Got '%s', expected '%s'", an, en)
	}

	cn := "Nicaragua"
	acn := tx.Contains[0].Name
	if acn != cn {
		t.Errorf("Got '%s', expected '%s'", acn, cn)
	}
}

func TestTaxonomyWithSpots(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/taxonomy-with-spots.json")
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, httpClient: http.DefaultClient}

	tq := TaxonomyQuery{
		ID:       "58f7ed58dadb30820bb38f8b",
		MaxDepth: 1,
		Type:     "taxonomy",
	}

	tqs, err := tq.TaxonomyQueryString()
	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.GetTaxonomy(tqs)
	if err != nil {
		t.Fatal(err)
	}

	en := "Ventura County"
	an := tx.Name
	if an != en {
		t.Errorf("Got '%s', expected '%s'", an, en)
	}

	var cs []Taxonomy
	for i := range tx.Contains {
		if tx.Contains[i].Type == "spot" {
			cs = append(cs, tx.Contains[i])
		}
	}

	csn := "Summer Beach"
	acsn := cs[0].Name
	if acsn != csn {
		t.Errorf("Got '%s', expected '%s'", acsn, csn)
	}

	csi := "5842041f4e65fad6a770895e"
	acsi := cs[0].Spot
	if acsn != csn {
		t.Errorf("Got '%s', expected '%s'", acsi, csi)
	}
}

func TestTaxonomyWithSubregions(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/taxonomy-with-subregions.json")
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, httpClient: http.DefaultClient}

	tq := TaxonomyQuery{
		ID:       "58f7ed51dadb30820bb387a6",
		MaxDepth: 0,
		Type:     "taxonomy",
	}

	tqs, err := tq.TaxonomyQueryString()
	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.GetTaxonomy(tqs)
	if err != nil {
		t.Fatal(err)
	}

	en := "California"
	an := tx.Name
	if an != en {
		t.Errorf("Got '%s', expected '%s'", an, en)
	}

	var cs []Taxonomy
	for i := range tx.Contains {
		if tx.Contains[i].Type == "subregion" {
			cs = append(cs, tx.Contains[i])
		}
	}

	csn := "Marin County"
	acsn := cs[0].Name
	if acsn != csn {
		t.Errorf("Got '%s', expected '%s'", acsn, csn)
	}

	csi := "58581a836630e24c44879009"
	acsi := cs[0].Subregion
	if acsn != csn {
		t.Errorf("Got '%s', expected '%s'", acsi, csi)
	}
}
