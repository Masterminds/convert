package color

import (
	"image/color"
	"testing"
)

func TestHexToRGB(t *testing.T) {

	// Test the error conditions of unusable hex values.
	_, _, _, err := HexToRGB("#1234")
	if err == nil {
		t.Error("! HexToRGB failed to error with an unusable hex value.")
	}

	_, _, _, err = HexToRGB("#GGG")
	if err == nil {
		t.Error("! HexToRGB failed to error with an unusable hex value.")
	}

	r, g, b, _ := HexToRGB("#FFF")
	if r != 255 || g != 255 || b != 255 {
		t.Error("! HexToRGB failed to convert a 3 digit hex value to RGB.")
	}

	r, g, b, _ = HexToRGB("#3C4C5C")
	if r != 60 || g != 76 || b != 92 {
		t.Error("! HexToRGB failed to convert a 6 digit hex value to RGB.")
	}
}

func TestHexToRGBColor(t *testing.T) {
	r, g, b, _ := HexToRGB("#FFF")

	foo := color.RGBA{r, g, b, 255}

	if foo.R != r {
		t.Error("! HexToRGB incompatible with color.RGBA.")
	}
}
