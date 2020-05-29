package surflinef

import "testing"

func TestQueryString(t *testing.T) {
	q := Query{
		Days:        3,
		SubregionID: "58581a836630e24c44878fd4",
	}

	qs, err := q.QueryString()
	if err != nil {
		t.Fatal(err)
	}

	e := "days=3&subregionId=58581a836630e24c44878fd4"
	if qs != e {
		t.Errorf("Got '%s', expected '%s'", qs, e)
	}
}
