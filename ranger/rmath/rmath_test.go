// Tests basic math functions
package rmath

import (
	"testing"
)

func Test_ToRadians(t *testing.T) {
	results := ToRadians(45.0)

	if !IsEqual(results, 0.7853981634) {
		t.Error("Expected radians to be ~0.7853981634")
	}
}

func Test_ToDegrees(t *testing.T) {
	results := ToDegrees(0.7853981634)
	if !IsEqual(results, 45.0) {
		t.Error("Expected degrees to be ~45.0")
	}
}
