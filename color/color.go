// Package color provides conversion functions for color units.
package color

import (
	"errors"
	"strconv"
	"strings"
)

// HexToRGB converts a hexadecimal color value to RGB. The hex value can
// optionally begin with a # which is common in web environments. The hex value
// can be either 3 or 6 characters long. For example, both #FFF and #FFFFFF are
// acceptable. #FFF will be expanded to #FFFFFF.
// The returned values are designed to be used with the color package. For
// example, you can do foo := color.RGBA{r, g, b, 255}.
func HexToRGB(hex string) (r, g, b uint8, err error) {
	// In many representations (e.g., web) there is a # prefixing the value.
	hex = strings.TrimPrefix(hex, "#")

	length := len(hex)
	var tmpVal int64
	if length == 6 {
		tmpString := hex[0:2]
		tmpVal, err = strconv.ParseInt(tmpString, 16, 0)
		r = uint8(tmpVal)

		tmpString = hex[2:4]
		tmpVal, err = strconv.ParseInt(tmpString, 16, 0)
		g = uint8(tmpVal)

		tmpString = hex[4:6]
		tmpVal, err = strconv.ParseInt(tmpString, 16, 0)
		b = uint8(tmpVal)
	} else if length == 3 {
		tmpString := hex[0:1]
		tmpVal, err = strconv.ParseInt(tmpString+tmpString, 16, 0)
		r = uint8(tmpVal)

		tmpString = hex[1:2]
		tmpVal, err = strconv.ParseInt(tmpString+tmpString, 16, 0)
		g = uint8(tmpVal)

		tmpString = hex[2:3]
		tmpVal, err = strconv.ParseInt(tmpString+tmpString, 16, 0)
		b = uint8(tmpVal)
	} else {
		err = errors.New("convert.color: The length of the hexadecimal string passed to HexToRGB is invalid.")
	}

	return
}
