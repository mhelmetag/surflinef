package surflinef

// Wind is a collection of wind info.
type Wind struct {
	DateStamp      [][]string  `json:"dateStamp"`
	PeriodSchedule [][]string  `json:"periodSchedule"`
	WindDirection  [][]float32 `json:"wind_direction"`
	WindSpeed      [][]float32 `json:"wind_speed"`
	Units          string      `json:"units"`
}
