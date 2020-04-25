package rmath

// See math.go for the expectations of the coordinate system.

import (
	"fmt"
)

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewRectangle creates a Rectangle initialized as:
// X, Y, Width, Height = 0.0 and uncentered.
func NewRectangle() *Rectangle {
	r := new(Rectangle)
	r.Top = 0.0
	r.Left = 0.0
	r.Bottom = 0.0
	r.Right = 0.0
	r.Width = 0.0
	r.Height = 0.0
	r.Centered = false
	return r
}

// NewRectangleUnCentered creates a Rectangle where the position is
// at x,y with dimensions of width x height.
// x,y are specified in local coordinates and should be zeroes.
func NewRectangleUnCentered(x, y, width, height float32) *Rectangle {
	r := new(Rectangle)
	r.Set(x, y, width, height, false)
	return r
}

// NewRectangleCentered creates a Rectangle where the position is
// centered between the position and dimensions.
// x,y are specified in local coordinates
func NewRectangleCentered(x, y, width, height float32) *Rectangle {
	r := new(Rectangle)
	r.Set(x, y, width, height, true)
	return r
}

// NewRectangleByCorners creates a Rectangle where the position is
// uncentered between the positions.
// x,y are specified in local coordinates
func NewRectangleByCorners(left, top, right, bottom float32) *Rectangle {
	r := new(Rectangle)
	r.SetByComp(left, top, right, bottom)
	return r
}

// ---------------------------------------------------------------------
// Setter/Getters
// ---------------------------------------------------------------------

// Set sets this rectangle's position and dimensions.
func (r *Rectangle) Set(x, y, width, height float32, centered bool) {
	if centered {
		hw := width / 2.0
		hh := height / 2.0
		r.Top = y - hh
		r.Left = x - hw
		r.Bottom = hh + y
		r.Right = hw + x
	} else {
		r.Top = y
		r.Left = x
		r.Bottom = height
		r.Right = width
	}

	r.Width = width
	r.Height = height
	r.Centered = centered
}

// SetByComp sets this rectangle's corner positions.
func (r *Rectangle) SetByComp(left, top, right, bottom float32) {
	r.Top = top
	r.Left = left
	r.Bottom = bottom
	r.Right = right

	r.Width = right - left
	r.Height = top - bottom
	r.Centered = false
}

// SetWithRectangle sets this rectangle based on `src` while compensating for centering.
func (r *Rectangle) SetWithRectangle(src *Rectangle) {
	if src.Centered {
		// Remove centering such that Set() works against uncentered values.
		hw := src.Width / 2.0
		hh := src.Height / 2.0
		r.Set(src.Left+hw, src.Top-hh, src.Width, src.Height, src.Centered)
	} else {
		r.Set(src.Left, src.Top, src.Width, src.Height, src.Centered)
	}
}

// SetSize sets the dimensions while compensating for centering.
func (r *Rectangle) SetSize(width, height float32) {
	if r.Centered {
		r.Set(0.0, 0.0, width, height, r.Centered)
	} else {
		r.Width = width
		r.Height = height
	}
}

// SetUniformSize sets the dimensions while compensating for centering.
func (r *Rectangle) SetUniformSize(length float32) {
	if r.Centered {
		// Upper/Lower need adjusting, so we adjust the upper coordinate.
		hw := r.Width / 2.0
		hh := r.Height / 2.0
		r.Set(r.Left+hw, r.Bottom+hh, length, length, r.Centered)
	} else {
		r.Width = length
		r.Height = length
	}
}

// GetCenter calculates the center of the rectangle and returns the results
// in 'out' with the z-component = 0.0.
func (r *Rectangle) GetCenter(out *Vector3) {
	out.Set3Components(r.Left+(r.Width/2.0), r.Top-(r.Height/2.0), 0.0)
}

// ---------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------

// ContainsPoint checks if point (aka x,y) is inside "inclusively".
func (r *Rectangle) ContainsPoint(x, y float32) bool {
	return x >= r.Left && x <= r.Right && y >= r.Top && y <= r.Bottom
}

// ContainsVector checks if point specified as a vector (aka v.X, v.Y) is inside "inclusively".
func (r *Rectangle) ContainsVector(v *Vector3) bool {
	return v.X >= r.Left && v.X <= r.Right && v.Y >= r.Top && v.Y <= r.Bottom
}

// ContainsRectangle checks if 'o' is inside this rectangle "inclusively".
func (r *Rectangle) ContainsRectangle(o *Rectangle) bool {
	contains := r.ContainsPoint(o.Left, o.Top)
	contains = contains && r.ContainsPoint(o.Right, o.Bottom)
	return contains
}

// ContainsPointNonInclusive checks if point (aka x,y) is inside "non-inclusively".
func (r *Rectangle) ContainsPointNonInclusive(x, y float32) bool {
	return x > r.Left && x < r.Right && y > r.Top && y < r.Bottom
}

// Overlaps checks if either overlaps the other.
func (r *Rectangle) Overlaps(o *Rectangle) bool {
	// Is one rectangle is on the left side of the other.
	if o.Left > r.Right || r.Left > o.Right {
		return false
	}

	// Is one rectangle above the other.
	if o.Top > r.Bottom || r.Top > o.Bottom {
		return false
	}

	return true
}

// Union returns the inclosing rectangle between 'this' and 'o'.
// i.e. the bounding box that incloses the rectangles.
func (r *Rectangle) Union(o *Rectangle) *Rectangle {
	u := NewRectangle()

	u.SetByComp(
		Min32(r.Left, o.Left),
		Max32(r.Top, o.Top),
		Max32(r.Right, o.Right),
		Min32(r.Bottom, o.Bottom))

	return u
}

// ---------------------------------------------------------------------
// Misc
// ---------------------------------------------------------------------

func (r Rectangle) String() string {
	return fmt.Sprintf("<%f, %f>:<%f, %f> %f x %f : Centered (%t)", r.Left, r.Top, r.Right, r.Bottom, r.Width, r.Height, r.Centered)
}
