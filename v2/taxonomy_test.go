package surflinef

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestTaxonomyQueryString(t *testing.T) {
	tq := TaxonomyQuery{
		Type:     "taxonomy",
		ID:       "58581a836630e24c44878fd4",
		MaxDepth: 0,
	}

	vs, err := query.Values(tq)

	if err != nil {
		t.Fatal(err)
	}

	qs := vs.Encode()

	e := "id=58581a836630e24c44878fd4&maxDepth=0&type=taxonomy"
	if qs != e {
		t.Errorf("Got '%s', expected '%s'", qs, e)
	}
}

func TestTaxonomy(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/taxonomy.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	tq := TaxonomyQuery{
		ID:       "58581a836630e24c44878fd4",
		MaxDepth: 0,
		Type:     "taxonomy",
	}

	tx, err := c.GetTaxonomy(tq)
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
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	tq := TaxonomyQuery{
		ID:       "58f7ed58dadb30820bb38f8b",
		MaxDepth: 1,
		Type:     "taxonomy",
	}

	tx, err := c.GetTaxonomy(tq)
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
	if acsi != csi {
		t.Errorf("Got '%s', expected '%s'", acsi, csi)
	}
}

func TestTaxonomyWithSubregions(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/taxonomy-with-subregions.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	tq := TaxonomyQuery{
		ID:       "58f7ed51dadb30820bb387a6",
		MaxDepth: 0,
		Type:     "taxonomy",
	}

	tx, err := c.GetTaxonomy(tq)
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

	csri := "58581a836630e24c44879009"
	acsri := cs[0].Subregion
	if acsri != csri {
		t.Errorf("Got '%s', expected '%s'", acsri, csri)
	}
}
