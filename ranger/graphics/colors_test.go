package graphics

import (
	"testing"

	"github.com/wdevore/ranger/rmath"
)

func Test_Colors_HexStringNonAlpha(t *testing.T) {
	c := NewColors()
	c.SetColorFromHex("#ffaa00")

	if !rmath.IsEqual(c.R, 1.0) {
		t.Error("Expected c.R = 1.0")
	}

	if !rmath.IsEqual(c.G, 0.666667) {
		t.Error("Expected c.G = 0.666667")
	}

	if !rmath.IsEqual(c.B, 0.0) {
		t.Error("Expected c.B = 0.0")
	}

	if !rmath.IsEqual(c.A, 0.0) {
		t.Error("Expected c.A = 0.0")
	}

	// println(c.String())
}

func Test_Colors_HexStringAlpha(t *testing.T) {
	c := NewColors()
	c.SetColorFromHex("#ffaa00a0")

	if !rmath.IsEqual(c.R, 1.0) {
		t.Error("Expected c.R = 1.0")
	}

	if !rmath.IsEqual(c.G, 0.666667) {
		t.Error("Expected c.G = 0.666667")
	}

	if !rmath.IsEqual(c.B, 0.0) {
		t.Error("Expected c.B = 0.0")
	}

	if !rmath.IsEqual(c.A, 0.627451) {
		t.Error("Expected c.A = 0.627451")
	}

	// println(c.String())
}
