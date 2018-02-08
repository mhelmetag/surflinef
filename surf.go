package surflinef

// Surf is a collection of detailed swell info.
// Agregate data is only available when the aggregate query param is true
// for the Forecast.
type Surf struct {
	SwellDirection1    [][]int     `json:"swell_direction1"`
	SwellDirection2    [][]int     `json:"swell_direction2"`
	SwellDirection3    [][]int     `json:"swell_direction3"`
	SwellPeriod1       [][]float32 `json:"swell_period1"`
	SwellPeriod2       [][]float32 `json:"swell_period2"`
	SwellPeriod3       [][]float32 `json:"swell_period3"`
	SwellHeight1       [][]float32 `json:"swell_height1"`
	SwellHeight2       [][]float32 `json:"swell_height2"`
	SwellHeight3       [][]float32 `json:"swell_height3"`
	SurfMin            [][]float32 `json:"surf_min"`
	SurfMax            [][]float32 `json:"surf_max"`
	PeriodSchedule     [][]string  `json:"periodSchedule"`
	DateStamp          [][]string  `json:"dateStamp"`
	Units              string      `json:"units"`
	AggregatePeriod    [][]float32 `json:"agg_period1"`
	AggregateHeight    [][]float32 `json:"agg_height1"`
	AggregateSpread    [][]int     `json:"agg_spread1"`
	AggregateDirection [][]int     `json:"agg_direction1"`
	AggregateLocation  [][]int     `json:"agg_location"`
	AggregateSurfMin   [][]float32 `json:"agg_surf_min"`
	AggregateSurfMax   [][]float32 `json:"agg_surf_max"`
}
