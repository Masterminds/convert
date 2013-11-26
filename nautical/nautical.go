package nautical

// Convert meters into nautical miles.
func MetersToNauticalMiles(meters float64) float64 {
	return meters / 1852
}

// Convert nautical miles to meters.
func NauticalMilesToMeters(nm float64) float64 {
	return nm * 1852
}
