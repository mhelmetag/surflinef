package surflinef

// Weather is a collection of weather info.
type Weather struct {
	DateStamp   []string  `json:"dateStamp"`
	TempMax     []float32 `json:"temp_max"`
	TempMin     []float32 `json:"temp_min"`
	WeatherType []string  `json:"weather_type"`
	Units       string    `json:"units"`
}
