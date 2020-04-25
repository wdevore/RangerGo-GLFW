// Tests matrix4 functions
package rmath

import (
	"testing"
)

func Test_Identity(t *testing.T) {
	m := NewMatrix4()
	if m.e[M00] != 1.0 {
		t.Error("Expected m00 = 1.0")
	}
	if m.e[M11] != 1.0 {
		t.Error("Expected m11 = 1.0")
	}
	if m.e[M22] != 1.0 {
		t.Error("Expected m22 = 1.0")
	}
	if m.e[M33] != 1.0 {
		t.Error("Expected m33 = 1.0")
	}
	// TODO complete check for zeroes
}

func Test_Translate(t *testing.T) {
	m := NewMatrix4()
	m.SetTranslate(NewVector3With2Components(5.0, 6.0))

	// Note: in order for this log to show in the Output window you need to
	// add "go.testFlags": ["-v"] to your User-Settings json file in Visual Studio Code.
	// t.Log("\n" + m.String())

	if m.C(M03) != 5.0 {
		t.Error("Expected m03 = 5.0")
	}
	if m.C(M13) != 6.0 {
		t.Error("Expected m13 = 6.0")
	}
}
