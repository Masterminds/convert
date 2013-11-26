package nautical

import (
	"testing"
)

const MaxDiff float64 = 0.0000001

func diff(a, b float64) bool {
	if a > b {
		return (a - b) > MaxDiff
	}
	return (b - a) > MaxDiff
}

func TestConvertNautical(t *testing.T) {

	if NauticalMilesToMeters(5) != 9260 {
		t.Error("! ConvertNauticalMilesToMeters incorrectly converting.")
	}

	if diff(MetersToNauticalMiles(1234), .6663067) {
		t.Error("! ConvertMetersToNauticalMiles incorrectly converting.")
	}
}
