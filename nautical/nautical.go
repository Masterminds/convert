// Package nautical provides conversion functions for nautical metrics.
package nautical

// MetersToNauticalMiles converts meters into nautical miles.
func MetersToNauticalMiles(meters float64) float64 {
	return meters / 1852
}

// NauticalMilesToMeters converts nautical miles to meters.
func NauticalMilesToMeters(nm float64) float64 {
	return nm * 1852
}
