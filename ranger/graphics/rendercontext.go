// Package graphics provides for view projection
package graphics

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

// RenderContext provides top level rendering, for example, FreeTypeFont
type RenderContext struct {
	clearColor Colors
}

// NewRenderContext construct a View
func NewRenderContext() *RenderContext {
	rc := new(RenderContext)
	return rc
}

// SetClearColor set the OpenGL background clear color
func (rc *RenderContext) SetClearColor(r, g, b, a float32) {
	rc.clearColor.Set(r, g, b, a)
	gl.ClearColor(rc.clearColor.R, rc.clearColor.G, rc.clearColor.B, rc.clearColor.A)
}

// SetClearColors set the OpenGL background clear color
func (rc *RenderContext) SetClearColors(cs *Colors) {
	rc.clearColor.SetFromColors(cs)
	gl.ClearColor(rc.clearColor.R, rc.clearColor.G, rc.clearColor.B, rc.clearColor.A)
}

// Clear clears color buffer
func (rc *RenderContext) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
