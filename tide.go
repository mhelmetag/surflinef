package surflinef

// Tide is a collection of tide predictions.
type Tide struct {
	DataPoints []DataPoint `json:"dataPoints"`
}

// DataPoint is a single tide prediction.
type DataPoint struct {
	Localtime string  `json:"Localtime"`
	Unixtime  int     `json:"time"`
	Type      string  `json:"type"`
	Height    float32 `json:"height"`
}
