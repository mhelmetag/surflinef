# SurflineF

[![Build Status](https://travis-ci.org/mhelmetag/surflinef.svg?branch=master)](https://travis-ci.org/mhelmetag/surflinef)

An API client for fetching Forecasts from the Surfline API.

## Installation

Simply run `go get github.com/mhelmetag/surflinef` and start using it in your own apps!

## Usage

The full example for fetching a Forecast (and other info) can be found in `examples/main.go` and can be run with `go run examples/main.go`.

## Surfline API URL

```
http://api.surfline.com/v1/forecasts
```

## Known Query Params

- **resources (string)** - possible values: \[surf, analysis, wind, weather, tide, sort\] (optional resources for forecast)
- **days (integer)** - greater than 1 (unsure of upper limit; confirmed to go above 10 but data usually becomes unknown/blank after that)
- **getAllSpots (boolean)** - _true_ will get all spot forecasts in the subregion (meaning an array of forecasts)
- **aggregate (boolean)** - _true_ enables aggregate fields for the Surf resource
- **units (string)** - possible values: \[e, m\] (e is feet and m is meters)
- **fullAnalysis (boolean)** - _true_ adds fields like `brief_outlook`, `best_bet`, `extended_outlook` and others to the larger Analysis JSON object
- **showOptimal (boolean)** - not sure what this does yet
- **interpolate (boolean)** - not sure what this does yet

## Example API Calls:

Found in the network tab when browsing around the Surfline site:

```
http://api.surfline.com/v1/forecasts/2141?&callback=jQuery18001517130109550182_1474259459851&resources=resources%3Dwind%2Csurf%2Canalysis%2Cweather%2Ctide%2Csort&days=17&aggregate=true&units=e&_=1474259492858

http://api.surfline.com/v1/forecasts/2141?resources=analysis&units=e&days=1&fullAnalysis=true

http://api.surfline.com/v1/forecasts/2141?resources=surf,analysis&days=1&getAllSpots=true&units=e&interpolate=false&showOptimal=false

http://api.surfline.com/v1/forecasts/4991?resources=surf&days=1&getAllSpots=false&units=e&interpolate=true&showOptimal=false
```
