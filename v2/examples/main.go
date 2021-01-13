package main

import (
	"fmt"
	"net/url"

	"github.com/mhelmetag/surflinef/v2"
)

func main() {
	getConditions()
	getTides()
	getTaxonomy()
	getWave()
}

func getConditions() {
	bu, err := url.Parse(surflinef.ConditionsBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.Query{
		Days:        3,
		SubregionID: "58581a836630e24c44878fd4", // Santa Barbara, CA
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Printf("Error building Query string: %v\n", err)
		return
	}

	cs, err := c.GetConditions(qs)
	if err != nil {
		fmt.Printf("Error fetching Conditions: %v\n", err)
		return
	}

	fmt.Printf("Conditions: %v\n", cs)
}

func getTides() {
	bu, err := url.Parse(surflinef.TidesBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.Query{
		Days:   3,
		SpotID: "5842041f4e65fad6a7708814", // Rincon, CA
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Printf("Error building Query string: %v\n", err)
		return
	}

	ts, err := c.GetTides(qs)
	if err != nil {
		fmt.Printf("Error fetching Tides: %v\n", err)
		return
	}

	fmt.Printf("Tides: %v\n", ts)
}

func getTaxonomy() {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	tq := surflinef.TaxonomyQuery{
		ID:       "58f7ed58dadb30820bb38f8b", // Ventura County, CA
		MaxDepth: 1,
		Type:     "taxonomy",
	}

	tqs, err := tq.TaxonomyQueryString()
	if err != nil {
		fmt.Printf("Error building Query string: %v\n", err)
		return
	}

	t, err := c.GetTaxonomy(tqs)
	if err != nil {
		fmt.Printf("Error fetching Taxonomy: %v\n", err)
		return
	}

	fmt.Printf("Taxonomy: %v\n", t)
}

func getWave() {
	bu, err := url.Parse(surflinef.WaveBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.Query{
		SpotID: "5842041f4e65fad6a7708814",
		Days:   1,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Printf("Error building Query string: %v\n", err)
		return
	}

	t, err := c.GetWave(qs)
	if err != nil {
		fmt.Printf("Error fetching Wave: %v\n", err)
		return
	}

	fmt.Printf("Wave: %v\n", t)
}
