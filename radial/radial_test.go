package radial

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

func TestDeg2rad(t *testing.T) {
	if diff(DegToRad(1), 0.0174533) {
		t.Error("! DegToRad is not accurately converting between degrees and radians.")
	}

	if diff(DegToRad(5), 0.0872665) {
		t.Error("! DegToRad is not accurately converting between degrees and radians.")
	}
}

func TestRad2deg(t *testing.T) {
	if diff(RadToDeg(1), 57.2957795) {
		t.Error("! RadToDeg is not accurately converting between degrees and radians.")
	}

	if diff(RadToDeg(5), 286.4788975) {
		t.Error("! RadToDeg is not accurately converting between degrees and radians.")
	}
}
