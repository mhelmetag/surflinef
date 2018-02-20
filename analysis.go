package surflinef

import (
	"encoding/json"
	"strings"
)

// Analysis is a collection of high-level surf info.
type Analysis struct {
	GeneralCondition []string
	GeneralText      []string
	SurfMax          []string
	SurfMin          []string
	SurfRange        []string
	SurfText         []string
	Units            string
}

type iAnalysis struct {
	GeneralCondition []string      `json:"generalCondition"`
	GeneralText      []string      `json:"generalText"`
	SurfMax          []json.Number `json:"surfMax"`
	SurfMin          []json.Number `json:"surfMin"`
	SurfRange        []string      `json:"surfRange"`
	SurfText         []string      `json:"surfText"`
	Units            string        `json:"units"`
}

func (iA iAnalysis) analysis() Analysis {
	var fSurfMax []string
	for i := range iA.SurfMax {
		sm := iA.SurfMax[i].String()
		nsm := normalize(sm)
		fSurfMax = append(fSurfMax, nsm)
	}

	var fSurfMin []string
	for i := range iA.SurfMin {
		sm := iA.SurfMin[i].String()
		nsm := normalize(sm)
		fSurfMin = append(fSurfMin, nsm)
	}

	var fGeneralText []string
	for i := range iA.GeneralText {
		gt := iA.GeneralText[i]
		sgt := strings.TrimSpace(gt)
		fGeneralText = append(fGeneralText, sgt)
	}

	return Analysis{
		iA.GeneralCondition,
		fGeneralText,
		fSurfMax,
		fSurfMin,
		iA.SurfRange,
		iA.SurfText,
		iA.Units,
	}
}

func normalize(sm string) string {
	if sm == "" {
		return "0"
	}

	return sm
}
