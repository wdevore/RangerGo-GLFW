package rmath

import (
	"testing"
)

func Test_UnCenteredRectangleConstruction(t *testing.T) {
	r := NewRectangleUnCentered(0.0, 0.0, 5.0, 10.0)
	t.Log("\n" + r.String())
	if r.Left != 0.0 {
		t.Error("Expected Rectangle's left component to = 0.0")
	}

	if r.Top != 0.0 {
		t.Error("Expected Rectangle's top component to = 0.0")
	}
	if r.Bottom != 10.0 {
		t.Error("Expected Rectangle's top component to = 10.0")
	}
	if r.Right != 5.0 {
		t.Error("Expected Rectangle's top component to = 5.0")
	}
}

func Test_CenteredRectangleConstruction(t *testing.T) {
	r := NewRectangleCentered(0.0, 0.0, 5.0, 10.0)
	t.Log("\n" + r.String())
	if r.Left != -2.5 {
		t.Errorf("Expected Rectangle's Left component to = -2.5, got: %f", r.Left)
	}

	if r.Top != -5.0 {
		t.Errorf("Expected Rectangle's Top component to = -5.0, got: %f", r.Top)
	}
	if r.Bottom != 5.0 {
		t.Errorf("Expected Rectangle's Bottom component to = 5.0, got: %f", r.Bottom)
	}
	if r.Right != 2.5 {
		t.Errorf("Expected Rectangle's Right component to = 2.5, got: %f", r.Right)
	}
}

func Test_Rectangle_SetSize(t *testing.T) {
	r := NewRectangleCentered(0.0, 0.0, 5.0, 10.0)
	t.Log("\n" + r.String())

	r.SetSize(10.0, 10.0)
	t.Log("\n" + r.String())

	if r.Left != -5.0 {
		t.Errorf("Expected Rectangle's Left component to = -5.0, got: %f", r.Left)
	}

	if r.Top != -5.0 {
		t.Errorf("Expected Rectangle's Top component to = -5.0, got: %f", r.Top)
	}
	if r.Bottom != 5.0 {
		t.Errorf("Expected Rectangle's Bottom component to = 5.0, got: %f", r.Bottom)
	}
	if r.Right != 5.0 {
		t.Errorf("Expected Rectangle's Right component to = 2.5, got: %f", r.Right)
	}
}

func Test_Rectangle_ContainsPoint(t *testing.T) {
	r := NewRectangleCentered(0.0, 0.0, 5.0, 10.0)
	t.Log("\n" + r.String())

	contains := r.ContainsPoint(2.0, 2.0)

	if !contains {
		t.Error("Expected <2.0, 2.0> to be in Rectangle")
	}

	contains = r.ContainsPoint(0.0, 2.0)

	if !contains {
		t.Error("Expected <0.0, 2.0> to be in Rectangle")
	}

	contains = r.ContainsPoint(-2.5, 2.0)

	if !contains {
		t.Error("Expected <-2.5, 2.0> to be in Rectangle")
	}

	contains = r.ContainsPoint(-2.6, 2.0)

	if contains {
		t.Error("Expected <-2.6, 2.0> NOT to be in Rectangle")
	}

	contains = r.ContainsPoint(-2.0, 9.0)

	if contains {
		t.Error("Expected <-2.0, 9.0> to be in Rectangle")
	}
}

func Test_Rectangle_ContainsRectangle(t *testing.T) {
	r1 := NewRectangleCentered(0.0, 0.0, 10.0, 10.0)
	// t.Log("\n" + r1.String())

	r2 := NewRectangleCentered(5.0, 5.0, 10.0, 10.0)
	// t.Log("\n" + r2.String())

	contains := r1.ContainsRectangle(r2)

	if contains {
		t.Error("Expected r1 to NOT be in Rectangle r2")
	}

	r3 := NewRectangleUnCentered(0.0, 0.0, 10.0, 10.0)
	// t.Log("\n" + r3.String())

	r4 := NewRectangleUnCentered(5.0, 5.0, 9.0, 9.0)
	// t.Log("\n" + r4.String())

	contains = r3.ContainsRectangle(r4)

	if !contains {
		t.Error("Expected r3 to be in Rectangle r4")
	}

	r4.Set(5.0, 5.0, 12.0, 12.0, false)
	// t.Log("\n" + r4.String())

	contains = r3.ContainsRectangle(r4)

	if contains {
		t.Error("Expected r3 to NOT be in Rectangle r4")
	}
}

func Test_Rectangle_Overlaps(t *testing.T) {
	r3 := NewRectangleUnCentered(0.0, 0.0, 10.0, 10.0)
	// t.Log("\n" + r3.String())

	r4 := NewRectangleUnCentered(5.0, 5.0, 9.0, 9.0)
	// t.Log("\n" + r4.String())

	overlaps := r3.Overlaps(r4)

	if !overlaps {
		t.Error("Expected r3 to overlap Rectangle r4")
	}

	r4.Set(5.0, 5.0, 12.0, 12.0, false)
	// t.Log("\n" + r4.String())

	overlaps = r3.Overlaps(r4)

	if !overlaps {
		t.Error("Expected r3 to overlap Rectangle r4 extended")
	}

	r4.Set(11.0, 11.0, 15.0, 15.0, false)
	// t.Log("\n" + r4.String())

	overlaps = r3.Overlaps(r4)

	if overlaps {
		t.Error("Expected r3 to NOT overlap Rectangle r4")
	}

}

// func Test_Rectangle_Encompass(t *testing.T) {
// 	r3 := NewRectangleByCorners(5.0, 20.0, 20.0, 10.0)
// 	t.Log("\n" + r3.String())

// 	r4 := NewRectangleByCorners(15.0, 15.0, 25.0, 5.0)
// 	t.Log("\n" + r4.String())

// 	ur := r3.Encompass(r4)
// 	t.Log("\n" + ur.String())

// 	// if !overlaps {
// 	// 	t.Error("Expected r3 to overlap Rectangle r4")
// 	// }

// }

func Test_Rectangle_Union(t *testing.T) {
	r3 := NewRectangleByCorners(5.0, 20.0, 20.0, 10.0)
	// t.Log("\n" + r3.String())

	r4 := NewRectangleByCorners(15.0, 15.0, 25.0, 5.0)
	// t.Log("\n" + r4.String())

	ur := r3.Union(r4)
	// t.Log("\n" + ur.String())

	if ur.Left != 5.0 {
		t.Errorf("Expected Rectangle's Left component to = 5.0, got: %f", ur.Left)
	}

	if ur.Top != 20.0 {
		t.Errorf("Expected Rectangle's Top component to = 20.0, got: %f", ur.Top)
	}
	if ur.Bottom != 5.0 {
		t.Errorf("Expected Rectangle's Bottom component to = 5.0, got: %f", ur.Bottom)
	}
	if ur.Right != 25.0 {
		t.Errorf("Expected Rectangle's Right component to = 25.0, got: %f", ur.Right)
	}
}
