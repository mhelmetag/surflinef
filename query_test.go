package surflinef

import "testing"

func TestQueryString(t *testing.T) {
	eqs := "days=3&fullAnalysis=true&resources=analysis%2Cwind&units=e"
	q := Query{
		Resources:    []string{"analysis", "wind"},
		Days:         3,
		Units:        "e",
		FullAnalysis: true,
	}
	qs, err := q.QueryString()
	if err != nil {
		t.Fatal(err)
	}

	if qs != eqs {
		t.Errorf("Got '%s', expected '%s'", qs, eqs)
	}
}
