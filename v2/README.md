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

- **subregionId (string)** - Can get this from the taxonomy API
- **days (integer)** - Greater than 1 and less than 6 (unless logged in)

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

- **spotId (string)** - Can get this from the taxonomy API
- **days (integer)** - Greater than 1 and less than 6 (unless logged in)

### Tides - Example URL

`https://services.surfline.com/kbyg/spots/forecasts/tides?spotId=5842041f4e65fad6a7708814&days=6`

### Tides - Data Structure

```json
{
  "timestamp": 1581062400,
  "type": "NORMAL",
  "height": 2.33
}
```

## Taxonomy

### Taxonomy - Base URL

`https://services.surfline.com/taxonomy`

### Taxonomy - Known Query Params

- **type (string)** - Can be a few things... taxonomy is one but not sure about others
- **id (string)** - ID of the taxonomy record (continent, country, region or area)
- **maxDepth (integer)** - Depth of the taxonomy search. Use 0 for searches for continents, country, regions and areas. Use 1 for area to find contained spots.

### Taxonomy - Example URL

`https://services.surfline.com/taxonomy?type=taxonomy&id=58f7ed51dadb30820bb38791&maxDepth=0`

### Taxonomy - Data Structure

Taxonomy is a recursive data structure. Basically, a taxonomy record `contains` a list of other taxonomy records that lie within in (`liesIn`).

```json
{
  "_id": "58f7ed51dadb30820bb38791",
  "geonameId": 6255149,
  "type": "geoname",
  "liesIn": [],
  "geonames": {},
  "location": {},
  "enumeratedPath": ",Earth,North America",
  "name": "North America",
  "category": "geonames",
  "hasSpots": true,
  "associated": {},
  "in": [],
  "contains": []
}
```
