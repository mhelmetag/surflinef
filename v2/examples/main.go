package main

import (
	"fmt"
	"net/url"

	"github.com/mhelmetag/surflinef/v2"
)

func main() {
	bu, err := url.Parse("https://services.surfline.com/kbyg/regions/forecasts/conditions")
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.Query{
		Days:        3,
		SubregionID: "58581a836630e24c44878fd4",
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
