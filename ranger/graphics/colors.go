// Package graphics provides for view projection
package graphics

import (
	"fmt"
	"regexp"
	"strconv"
)

// Detects 0xaabbcc(dd) or #aabbcc(dd)
const hexPattern = `^(#|0x)([\daAbBcCdDeEfF]{6,8})$`

var rexp, _ = regexp.Compile(hexPattern)

// ----------------------------------------------------------
// Predefined colors
// ----------------------------------------------------------
var (
	Transparent = NewColors().SetFromInts(0, 0, 0, 0)
	White       = NewColors().SetFromInts(255, 255, 255, 255)
	Black       = NewColors().SetFromInts(0, 0, 0, 255)
	Grey        = NewColors().SetFromInts(128, 128, 128, 255)
	Red         = NewColors().SetFromInts(255, 0, 0, 255)
	Green       = NewColors().SetFromInts(0, 255, 0, 255)
	Blue        = NewColors().SetFromInts(0, 0, 255, 255)
	Yellow      = NewColors().SetFromInts(255, 255, 0, 255)
	NavyBlue    = NewColors().SetFromInts(10, 20, 100, 255)
	LightBlue   = NewColors().SetFromInts(128, 128, 255, 255)
	GreyBlue    = NewColors().SetFromInts(72, 100, 180, 255)
	DarkBlue    = NewColors().SetFromInts(109, 157, 235, 255)
	Orange      = NewColors().SetFromInts(255, 127, 0, 255)
	GoldYellow  = NewColors().SetFromInts(255, 200, 0, 255)
	GreenYellow = NewColors().SetFromInts(173, 255, 47, 255)
	YellowGreen = NewColors().SetFromInts(154, 205, 50, 255)

	// Pantone colors
	// http://damonbauer.github.io/Pantone-Sass/
	PanSkin, _   = NewColors().SetColorFromHex("#fcc89bff")
	PanPurple, _ = NewColors().SetColorFromHex("#8031a7ff") // 527-C
)

// Colors provides conversions to and from different formats.
type Colors struct {
	R, G, B, A float32
}

// NewColors construct a Colors
func NewColors() *Colors {
	c := new(Colors)
	return c
}

// Set color
func (c *Colors) Set(r, g, b, a float32) *Colors {
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	return c
}

// SetFromInts color
func (c *Colors) SetFromInts(r, g, b, a int) *Colors {
	c.R = float32(r) / 255.0
	c.G = float32(g) / 255.0
	c.B = float32(b) / 255.0
	c.A = float32(a) / 255.0
	return c
}

// SetFromColors set color from Colors
func (c *Colors) SetFromColors(cs *Colors) *Colors {
	return c.Set(cs.R, cs.G, cs.B, cs.A)
}

// SetColorFromHex takes a hex string formatted as either "0xaabbcc(dd)" or "#aabbcc(dd)"
// where mixed case is allowed and "dd" is an optional Alpha value.
func (c *Colors) SetColorFromHex(hex string) (*Colors, error) {
	if rexp == nil {
		panic("Reqular expression not created")
	}

	subMatch := rexp.FindStringSubmatch(hex)

	if subMatch == nil {
		return nil, fmt.Errorf("Unable to find a match for hex value: %s", hex)
	}

	// The first element, 0, is the entire string match
	// We want the second group which is second element: subMatch[2]
	subHex := subMatch[2]

	// elements are specified as [low, high)
	// 012345
	// aabbcc
	//
	// "aa" starts at 0 and goes to 2 non-inclusive, thus [0,1] = [0,2)
	r := subHex[0:2]
	g := subHex[2:4]
	b := subHex[4:6]

	i, _ := strconv.ParseUint(r, 16, 32)
	c.R = float32(i) / 255.0

	i, _ = strconv.ParseUint(g, 16, 32)
	c.G = float32(i) / 255.0

	i, _ = strconv.ParseUint(b, 16, 32)
	c.B = float32(i) / 255.0

	if len(subHex) > 6 {
		a := subHex[6:8]
		i, _ = strconv.ParseUint(a, 16, 32)
		c.A = float32(i) / 255.0
	}

	// println(c.StringUnit())
	return c, nil
}

func (c Colors) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", c.R*255.0, c.G*255.0, c.B*255.0, c.A)
}

// StringUnit returns a unit representation
func (c Colors) StringUnit() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", c.R, c.G, c.B, c.A)
}
