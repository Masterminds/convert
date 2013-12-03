// Package color provides conversion functions for color units.
package color

import (
	"errors"
	"strings"
	"strconv"
)

// HexToRGB converts a hexadecimal color value to RGB. The hex value can
// optionally begin with a # which is common in web environments. The hex value
// can be either 3 or 6 characters long. For example, both #FFF and #FFFFFF are
// acceptable. #FFF will be expanded to #FFFFFF.
func HexToRGB(hex string) (r, g, b int64, err error) {
	// In many representations (e.g., web) there is a # prefixing the value.
	hex = strings.TrimPrefix(hex, "#")

	length := len(hex)
	if length == 6 {
		tmpString := hex[0:2]
		r, err = strconv.ParseInt(tmpString, 16, 0)

		tmpString = hex[2:4]
		g, err = strconv.ParseInt(tmpString, 16, 0)

		tmpString = hex[4:6]
		b, err = strconv.ParseInt(tmpString, 16, 0)
	} else if length == 3 {
		tmpString := hex[0:1]
		r, err = strconv.ParseInt(tmpString + tmpString, 16, 0)

		tmpString = hex[1:2]
		g, err = strconv.ParseInt(tmpString + tmpString, 16, 0)

		tmpString = hex[2:3]
		b, err = strconv.ParseInt(tmpString + tmpString, 16, 0)
	} else {
		err = errors.New("convert.color: The length of the hexadecimal string passed to HexToRGB is invalid.")
	}

	return
}
