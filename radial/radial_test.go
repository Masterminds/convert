package radial

import (
	"testing"
)

const MAX_DIFF float64 = 0.0000001

func diff(a, b float64) bool {
	if a > b {
		return (a - b) > MAX_DIFF
	}
	return (b - a) > MAX_DIFF
}

func TestDeg2rad(t *testing.T) {
	if diff(Deg2rad(1), 0.0174533) {
		t.Error("! Deg2rad is not accurately converting between degrees and radians.")
	}

	if diff(Deg2rad(5), 0.0872665) {
		t.Error("! Deg2rad is not accurately converting between degrees and radians.")
	}
}

func TestRad2deg(t *testing.T) {
	if diff(Rad2deg(1), 57.2957795) {
		t.Error("! Rad2deg is not accurately converting between degrees and radians.")
	}

	if diff(Rad2deg(5), 286.4788975) {
		t.Error("! Rad2deg is not accurately converting between degrees and radians.")
	}
}
