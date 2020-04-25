package rmath

import (
	"testing"
)

func Test_Construction(t *testing.T) {
	v := NewVector3()
	if v.X != 0.0 {
		t.Error("Expected Vector3's x component to = 0.0")
	}

	if v.Y != 0.0 {
		t.Error("Expected Vector3's y component to = 0.0")
	}

	if v.Z != 0.0 {
		t.Error("Expected Vector3's z component to = 0.0")
	}
}

func Test_3Component_Construction(t *testing.T) {
	v := NewVector3With3Components(1.0, 2.0, 3.0)
	if v.X != 1.0 {
		t.Error("Expected Vector3's x component to = 1.0")
	}

	if v.Y != 2.0 {
		t.Error("Expected Vector3's y component to = 2.0")
	}

	if v.Z != 3.0 {
		t.Error("Expected Vector3's z component to = 3.0")
	}
}

func Test_2Component_Construction(t *testing.T) {
	v := NewVector3With2Components(1.0, 2.0)
	if v.X != 1.0 {
		t.Error("Expected Vector3's x component to = 1.0")
	}

	if v.Y != 2.0 {
		t.Error("Expected Vector3's y component to = 2.0")
	}

	if v.Z != 0.0 {
		t.Error("Expected Vector3's z component to = 0.0")
	}
}

func Test_Clone(t *testing.T) {
	v := NewVector3With2Components(1.0, 2.0)
	c := v.Clone()
	if c.X != 1.0 {
		t.Error("Expected copy's x component to = 1.0")
	}

	if c.Y != 2.0 {
		t.Error("Expected copy's y component to = 2.0")
	}

	if c.Z != 0.0 {
		t.Error("Expected copy's z component to = 0.0")
	}
}

func Test_Set3Components(t *testing.T) {
	v := NewVector3()
	v.Set3Components(4.0, 5.0, 6.0)
	if v.X != 4.0 {
		t.Error("Expected Vector3's x component to = 4.0")
	}

	if v.Y != 5.0 {
		t.Error("Expected Vector3's y component to = 5.0")
	}

	if v.Z != 6.0 {
		t.Error("Expected Vector3's z component to = 6.0")
	}
}

func Test_Set2Components(t *testing.T) {
	v := NewVector3()
	v.Set2Components(44.0, 55.0)
	if v.X != 44.0 {
		t.Error("Expected Vector3's x component to = 44.0")
	}

	if v.Y != 55.0 {
		t.Error("Expected Vector3's y component to = 55.0")
	}

	if v.Z != 0.0 {
		t.Error("Expected Vector3's z component to = 0.0")
	}
}

func Test_SetComponents(t *testing.T) {
	v := NewVector3()
	s := NewVector3With3Components(44.0, 55.0, 66.0)
	v.Set(s)

	if v.X != 44.0 {
		t.Error("Expected Vector3's x component to = 44.0")
	}

	if v.Y != 55.0 {
		t.Error("Expected Vector3's y component to = 55.0")
	}

	if v.Z != 66.0 {
		t.Error("Expected Vector3's z component to = 66.0")
	}
}

func Test_Vector3_RotateBy(t *testing.T) {
	m := NewMatrix4()

	m.SetRotation(ToRadians(45.0))

	v := NewVector3With2Components(1.0, 0.0)

	// Rotate v using rotation matrix.
	// Y axis <0.0, 1.0>
	// ^
	// |     ^
	// |    / v vector rotated 45 degrees to <0.707, 0.707>
	// |  /
	// |/
	// .---------> X axis  <1.0, 0>
	//
	// A positive rotation yields X vector pointing upwards.
	v.Mul(m)

	// Note: in order for this log to show in the Output window you need to
	// add "go.testFlags": ["-v"] to your User-Settings json file in Visual Studio Code.
	// t.Log("\n" + m.String())

	if !IsEqual(v.X, 0.707107) {
		t.Error("Expected v.X = 0.707107")
	}
	if !IsEqual(v.Y, 0.707107) {
		t.Error("Expected v.Y = 0.707107")
	}
	if !IsEqual(v.Z, 0.0) {
		t.Error("Expected v.Z = 0.0")
	}

}
