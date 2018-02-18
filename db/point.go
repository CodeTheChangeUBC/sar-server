package db

import (
	"math"
	"regexp"
	"strconv"
)

var validLatLong = regexp.MustCompile(`^-?[0-9]+\.[0-9]+$`)

// A Point on the map.
type Point struct {
	Latitude  string
	Longitude string
}

// Validate the `Latitude` and `Longitude` of the given point.
func (p *Point) Validate() bool {
	lat, err := strconv.ParseFloat(p.Latitude, 64)
	if err != nil || math.Abs(lat) > 90.0 {
		return false
	}

	lon, err := strconv.ParseFloat(p.Longitude, 64)
	if err != nil || math.Abs(lon) > 90.0 {
		return false
	}

	return true
}
