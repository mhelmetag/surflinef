package main

import (
	"fmt"

	"github.com/mhelmetag/surflinef"
)

func main() {
	c, err := surflinef.DefaultClient()
	if err != nil {
		fmt.Printf("Error building Client: %v\n", err)
		return
	}

	q := surflinef.Query{
		Resources:    []string{"analysis"},
		Days:         3,
		Units:        "e",
		FullAnalysis: true,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Printf("Error building Query string: %v\n", err)
		return
	}

	f, err := c.GetForecast("2141", qs)
	if err != nil {
		fmt.Printf("Error fetching Forecast: %v\n", err)
		return
	}

	fmt.Printf("Forecast fetched!\n%#v\n", f)
}
