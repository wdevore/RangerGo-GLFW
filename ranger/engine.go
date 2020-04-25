// Package ranger : This implements the main Engine.
package ranger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/wdevore/ranger/config"
	"github.com/wdevore/ranger/graphics"
	"github.com/wdevore/ranger/window"
)

// Engine is the core component for launching and running the game.
type Engine struct {
	// The game is the client of the Engine.
	game GameShell

	fullScreen bool

	config config.Settings

	engineError error

	// ---------------------------------------------------------------------
	// Timing
	// ---------------------------------------------------------------------
	currentUpdateTime float64
	deltaUpdateTime   float64
	deltaTime         float64
	deltaRenderTime   float64
	currentRenderTime float64
	currentSwapTime   float64
	deltaSwapTime     float64

	// ---------------------------------------------------------------------
	// Window
	// ---------------------------------------------------------------------
	rWindow *window.RWindow

	// ---------------------------------------------------------------------
	// OpenGL
	// ---------------------------------------------------------------------
	Viewport graphics.Viewport
	Camera   graphics.Camera
	View     graphics.View

	renderContext graphics.RenderContext

	// ---------------------------------------------------------------------
	// Stage and Scene
	// ---------------------------------------------------------------------
	stage *Stage
}

// NewEngine requires a GameShell for notifying the developer that
// it is valid to configure their game.
func NewEngine(gs GameShell) *Engine {
	e := new(Engine)
	e.game = gs

	return e
}

// Launch the game.
func (e *Engine) Launch() error {
	println("Engine configuring...")

	e.loadConfig()

	// Now notify developer that they can configure their game.
	configured := e.game.Configure(e)

	// Finally start the engine.
	if configured {
		if !e.config.Engine.Enabled {
			return errors.New("engine is NOT enabled in config file")
		}

		e.engineError = e.start()

		if e.engineError != nil {
			return e.engineError
		}
	}

	return nil
}

func (e *Engine) loadConfig() {
	// Read configuration JSON file for pre-config settings.
	// ex := os.Getenv("GOPATH")
	// // if err != nil {
	// // 	panic("An error occurred trying to get directory of executable")
	// // }
	// fmt.Printf("executable working directory: %s\n", ex)

	// workingDirectory, errw := filepath.Abs(filepath.Dir(os.Args[0]))
	// if errw != nil {
	// 	panic("An error occurred trying to get working directory")
	// }

	// fmt.Printf("working directory: %s\n", workingDirectory)

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		panic("An error occurred trying to open ./config.json")
	}

	// The JSON must be preprocessed first to remove non-compliant comments.
	pattern := `(\/\*[\w .>=:\"\-\n\r\t]*\*\/|\/\/[ .>=:\"\w\-]*)`

	re, _ := regexp.Compile(pattern)

	cleanedJSON := re.ReplaceAllString(string(file), "")

	// fmt.Printf("%s\n", string(cleanedJSON))

	err = json.Unmarshal([]byte(cleanedJSON), &e.config)
	if err != nil {
		s := fmt.Sprintf("Failed to Unmarshal json object: %s\n", err.Error())
		panic(s)
	}
}

// ----------------------------------------------------------------------------
// Life cycles
// ----------------------------------------------------------------------------

// start will not exit until engine is told to exit.
func (e *Engine) start() error {
	println("Ranger Engine is starting...")

	e.rWindow = window.NewRWindow()

	err := e.rWindow.Construct(&e.config)

	if err != nil {
		return err
	}

	e.configureStage(&e.config)

	e.renderContext.SetClearColors(graphics.Orange)

	e.loop()

	return nil
}

// Stop performs any last minute resource cleanups
func (e *Engine) Stop() {
	println("Engine stopping...")
	println("Engine stopped")

	if e.engineError == nil {
		glfw.Terminate()
	}
}

func (e *Engine) loop() {
	for e.rWindow.IsRunning() {
		e.rWindow.Poll()

		// ---------------- Update BEGIN -----------------------------
		e.currentUpdateTime = glfw.GetTime()
		e.deltaUpdateTime = glfw.GetTime() - e.currentUpdateTime
		// ---------------- Update END -----------------------------

		// This clear sync locked with the vertical refresh. The clear itself
		// takes ~30 microseconds on a mid-range mobile nvidia GPU.
		e.renderContext.Clear()

		e.rWindow.Swap()
	}
}

func (e *Engine) configureStage(config *config.Settings) {
	e.Viewport.SetDimensions(0, 0, config.Window.DeviceRes.Width, config.Window.DeviceRes.Height)
	e.Viewport.Apply()

	// Calc the aspect ratio between the physical (aka device) dimensions and the
	// the virtual (aka user's design choice) dimensions.

	deviceRatio := float64(config.Window.DeviceRes.Width) / float64(config.Window.DeviceRes.Height)
	virtualRatio := float64(config.Window.VirtualRes.Width) / float64(config.Window.VirtualRes.Height)

	xRatioCorrection := float64(config.Window.DeviceRes.Width) / float64(config.Window.VirtualRes.Width)
	yRatioCorrection := float64(config.Window.DeviceRes.Height) / float64(config.Window.VirtualRes.Height)

	var ratioCorrection float64

	if virtualRatio < deviceRatio {
		ratioCorrection = yRatioCorrection
	} else {
		ratioCorrection = xRatioCorrection
	}

	e.Camera.SetProjection(
		float32(ratioCorrection),
		0.0, 0.0,
		float32(config.Window.DeviceRes.Height), float32(config.Window.DeviceRes.Width))

	if config.Camera.Centered {
		e.Camera.Centered()
	}

	e.View.SetProjection(config.Camera.View.X, config.Camera.View.Y, config.Camera.View.Z)

	// -----------------------------------------------------------------
	// Create stage
	// -----------------------------------------------------------------

}
