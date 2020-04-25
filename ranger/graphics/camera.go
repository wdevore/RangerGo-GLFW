// Package graphics provides visual
package graphics

import (
	"github.com/wdevore/ranger/rmath"
)

// Camera provides an orthographic project
type Camera struct {
	near, far                float32
	left, right, bottom, top float32
	Width, Height            float32
	ratioCorrection          float32

	// Projection matrix (orthographic)
	Matrix rmath.Matrix4
}

// NewCamera construct a Camera
func NewCamera() *Camera {
	c := new(Camera)
	return c
}

// SetProjection sets orthographic frustum dimensions
func (c *Camera) SetProjection(ratioCorrection, bottom, left, top, right float32) {
	c.ratioCorrection = ratioCorrection

	c.bottom = bottom
	c.left = left
	c.top = top
	c.right = right
	c.Width = right - left
	c.Height = top - bottom

	c.Matrix.SetToOrtho(0.0, 0.0, c.Width, c.Height, 0.1, 100.0)
}

// Centered centers the projection and adjusted for aspect ratio
// If you choose to center the view then don't call Centered otherwise you get
// compounded centering.
func (c *Camera) Centered() {
	// Adjust for aspect ratio
	left := -c.Width / 2.0 / c.ratioCorrection
	right := c.Width / 2.0 / c.ratioCorrection
	bottom := -c.Height / 2.0 / c.ratioCorrection
	top := c.Height / 2.0 / c.ratioCorrection

	c.Matrix.SetToOrtho(left, right, bottom, top, 0.1, 100.0)
}
