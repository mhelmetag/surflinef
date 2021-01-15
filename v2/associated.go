package surflinef

// Units is a collection of units for measurements (heights, speeds, etc)
type Units struct {
	Temperature string `json:"temperature"`
	TideHeight  string `json:"tideHeight"`
	SwellHeight string `json:"swellHeight"`
	WaveHeight  string `json:"waveHeight"`
	WindSpeed   string `json:"windSpeed"`
	Model       string `json:"model"`
}

// Location contains Lat and Lon(g) for related places
type Location struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
