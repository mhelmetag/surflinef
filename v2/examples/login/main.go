package main

import (
	"fmt"
	"net/url"

	"github.com/mhelmetag/surflinef/v2"
)

const USERNAME = ""
const PASSWORD = ""

func main() {
	accessToken := login()

	if accessToken == "" {
		fmt.Println("There was an unexpected error logging in")
	} else {
		fmt.Println("Logged in")
		getMoreConditions(accessToken)
	}
}

func login() string {
	bu, err := url.Parse(surflinef.LoginBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return ""
	}

	c := surflinef.Client{BaseURL: bu}

	lq := surflinef.LoginQuery{
		IsShortLived: false,
	}

	lp := surflinef.DefaultLoginPayload(USERNAME, PASSWORD)

	lr, err := c.PostLogin(lq, lp)
	if err != nil {
		fmt.Printf("Error logging in: %v\n", err)
		return ""
	}

	return lr.AccessToken
}

func getMoreConditions(accessToken string) {
	bu, err := url.Parse(surflinef.ConditionsBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	c := surflinef.Client{BaseURL: bu}

	cq := surflinef.ConditionsQuery{
		Days:        10,                         // This normally wouldn't work
		SubregionID: "58581a836630e24c44878fd4", // Santa Barbara, CA
		AccessToken: accessToken,
	}

	cs, err := c.GetConditions(cq)
	if err != nil {
		fmt.Printf("Error fetching Conditions: %v\n", err)
		return
	}

	fmt.Printf("Conditions: %v\n", cs)
}
