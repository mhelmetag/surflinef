# SurflineF

[![Build Status](https://travis-ci.org/mhelmetag/surflinef.svg?branch=master)](https://travis-ci.org/mhelmetag/surflinef)

An API client for fetching data from the Surfline API v2.

## Installation

Make sure to use versions 2.x.x or greater in your `go.mod` file.

## Usage

The full example for fetching a Forecast (and other info) can be found in `examples/main.go` and can be run with `go run examples/main.go`.

## Conditions

### Conditions - Base URL

`https://services.surfline.com/kbyg/regions/forecasts/conditions`

### Conditions - Known Query Params

- **subregionId (string)** - can get this from the taxonomy API
- **days (integer)** - greater than 1 and less than 6 (unless logged in)

### Conditions - Example URL

`https://services.surfline.com/kbyg/regions/forecasts/conditions?subregionId=58581a836630e24c44878fd4&days=6`

### Conditions - Data Structure

```json
{
  "timestamp": 1581148800,
  "forecaster": {
    "name": "Schaler Perry",
    "avatar": "https://www.gravatar.com/avatar/ea1e9a0c570c61d61dec3cf6ea26a85e?d=mm"
  },
  "human": true,
  "observation": "NW swell continues. Deep mid-morning high tide.",
  "am": {
    "maxHeight": 3,
    "minHeight": 2,
    "plus": false,
    "humanRelation": "2-3 ft – knee to waist high",
    "occasionalHeight": null,
    "rating": "FAIR"
  },
  "pm": {
    "maxHeight": 3,
    "minHeight": 2,
    "plus": false,
    "humanRelation": "2-3 ft – knee to waist high",
    "occasionalHeight": null,
    "rating": "POOR_TO_FAIR"
  }
}
```

## Tides

### Tides - Base URL

`https://services.surfline.com/kbyg/spots/forecasts/tides`

### Tides - Known Query Params

- **subregionId (string)** - can get this from the taxonomy API
- **days (integer)** - greater than 1 and less than 6 (unless logged in)

### Tides - Example URL

`https://services.surfline.com/kbyg/spots/forecasts/tides?subregionId=58581a836630e24c44878fd4&days=6`

### Tides - Data Structure

```json
{
  "timestamp": 1581062400,
  "type": "NORMAL",
  "height": 2.33
}
```
