package surflinef

// Associated is associated info to go along with the API response
// It usually includes units of measurement, utc offset for timezones, etc
type Associated struct {
	Units            Units    `json:"units"`
	UTCOffset        int32    `json:"utcOffset"`
	Location         Location `json:"location"`
	ForecastLocation Location `json:"forecastLocation"`
	OffshoreLocation Location `json:"offshoreLocation"`
	TideLocation     Location `json:"tideLocation"`
}

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
