// Package radial provides conversion functions for angle units.
package radial

import (
	"math"
)

// DegToRad converts degrees to radians.
func DegToRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}

// RadToDeg converts radians to degrees.
func RadToDeg(rad float64) float64 {
	return (rad * 180) / math.Pi
}
