// Package graphics provides for view projection
package graphics

import (
	"github.com/wdevore/ranger/rmath"
)

// View provides an orthographic project
type View struct {
	xOffset, yOffset, zOffset float32

	// Projection matrix (orthographic)
	Matrix rmath.Matrix4
}

// NewView construct a View
func NewView() *View {
	v := new(View)
	return v
}

// SetProjection sets orthographic projection used by shaders
func (v *View) SetProjection(xOffset, yOffset, zOffset float32) {

	v.xOffset = xOffset
	v.yOffset = yOffset
	v.zOffset = zOffset

	// For this engine x and y are always zero and z is a small negative number like -1.0
	v.Matrix.SetTranslate3Comp(v.xOffset, v.yOffset, v.zOffset)
}
