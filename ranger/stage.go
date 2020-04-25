package ranger

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/wdevore/ranger/config"
	"github.com/wdevore/ranger/rmath"
)

// Stage manages the view and projection.
type Stage struct {
	viewProjection rmath.Matrix4
	FillPolyMode   bool

	settings *config.Settings
}

// NewStage creates a stage
func (st *Stage) NewStage(se *config.Settings) *Stage {
	sa := new(Stage)
	sa.settings = se
	return sa
}

// Initialize configures a stage with a SceneManager
func (st *Stage) Initialize(e *Engine) {
	st.viewProjection.Set(&e.Camera.Matrix)
	st.viewProjection.PostMultiply(&e.View.Matrix)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
}

// Settings returns the engine's configuration settings.
func (st *Stage) Settings() *config.Settings {
	return st.settings
}

func (st *Stage) step(dt float32) bool {

	return true
}

func (st *Stage) exit() {

}
