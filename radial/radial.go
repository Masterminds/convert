// Radial math conversions.
package radial

import (
	"math"
)

// Convert degrees to radians.
func DegToRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}

// Convert radians to degrees.
func RadToDeg(rad float64) float64 {
	return (rad * 180) / math.Pi
}
