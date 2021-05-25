package surflinef

import (
	"net/url"
	"testing"
)

func TestConditionsResponse(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/conditions.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	cq := ConditionsQuery{
		SubregionID: "58581a836630e24c44878fd4",
		Days:        6,
	}

	cr, err := c.GetConditions(cq)
	if err != nil {
		t.Fatal(err)
	}

	cs := cr.Data.Conditions
	cs0 := cs[0]

	eO := "Steep new NW swell peaks early. Steady NNE winds. "
	aO := cs0.Observation
	if aO != eO {
		t.Errorf("Got '%s', expected '%s'", aO, eO)
	}

	eM := 3.0
	aM := cs0.AM.MinHeight
	if aM != eM {
		t.Errorf("Got '%f', expected '%f'", aM, eM)
	}
}

func TestCondtionsResponseWithPlus(t *testing.T) {
	ts, err := setupFixtureServer("fixtures/conditions.json")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu}

	cq := ConditionsQuery{
		SubregionID: "58581a836630e24c44878fd4",
		Days:        6,
	}

	cr, err := c.GetConditions(cq)
	if err != nil {
		t.Fatal(err)
	}

	cs := cr.Data.Conditions
	cs0 := cs[1]

	eP := true
	aP := cs0.AM.Plus
	if aP != eP {
		t.Errorf("Got '%t', expected '%t'", aP, eP)
	}

	eOH := 4.0
	aOH := cs0.AM.OccasionalHeight
	if aP != eP {
		t.Errorf("Got '%f', expected '%f'", aOH, eOH)
	}
}
