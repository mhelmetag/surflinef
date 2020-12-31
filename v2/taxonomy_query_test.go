package surflinef

import "testing"

func TestTaxonomyQueryString(t *testing.T) {
	q := TaxonomyQuery{
		Type:     "taxonomy",
		ID:       "58581a836630e24c44878fd4",
		MaxDepth: 0,
	}

	qs, err := q.TaxonomyQueryString()
	if err != nil {
		t.Fatal(err)
	}

	e := "id=58581a836630e24c44878fd4&maxDepth=0&type=taxonomy"
	if qs != e {
		t.Errorf("Got '%s', expected '%s'", qs, e)
	}
}
