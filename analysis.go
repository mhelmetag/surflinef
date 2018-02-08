package surflinef

import "encoding/json"

// Analysis is a collection of high-level surf info.
type Analysis struct {
	GeneralCondition []string      `json:"generalCondition"`
	GeneralText      []string      `json:"generalText"`
	SurfMax          []json.Number `json:"surfMax"`
	SurfMin          []json.Number `json:"surfMin"`
	SurfRange        []string      `json:"surfRange"`
	SurfText         []string      `json:"surfText"`
	Units            string        `json:"units"`
}
