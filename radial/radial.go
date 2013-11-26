package radial

import (
	"math"
)

// Convert degrees to radians.
func Deg2rad(deg float64) float64 {
	return (deg * math.Pi) / 180
}

// Convert radians to degrees.
func Rad2deg(rad float64) float64 {
	return (rad * 180) / math.Pi
}
