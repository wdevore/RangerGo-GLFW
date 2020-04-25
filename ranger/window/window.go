// Package window manages GLFW window functionality
package window

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/wdevore/ranger/config"
)

func init() {
	// GLFW is only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// RWindow manages basic window construction and functionality
type RWindow struct {
	window *glfw.Window

	quitTriggered bool
}

// NewRWindow creates a new Window
func NewRWindow() *RWindow {
	w := new(RWindow)
	return w
}

// IsRunning checks if the window has closed
func (w *RWindow) IsRunning() bool {
	return !w.window.ShouldClose()
}

// Poll checks for quit or polls events
func (w *RWindow) Poll() {
	if w.quitTriggered {
		w.window.SetShouldClose(true)
	} else {
		glfw.PollEvents()
	}
}

// Swap swaps buffers
// SwapBuffers is synced to the vertical which means it is waits based on the monitor refresh rate.
// The Clear is also locked to the sync, so if we don't swap the display just waits/locks thus the
// engine appears frozen.
func (w *RWindow) Swap() {
	w.window.SwapBuffers()
}

// Construct initializes
func (w *RWindow) Construct(config *config.Settings) error {

	err := w.initGLFW(config)

	if err != nil {
		return err
	}

	err = w.initGL(config)

	if err != nil {
		return err
	}

	return nil
}

func (w *RWindow) initGLFW(config *config.Settings) error {
	// Init will call glfw.Terminate if it fails.
	println("Initializing GLFW...")
	err := glfw.Init()
	if err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// NOTE: Required by OSX! Otherwise the app crashes.
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	// Create a GLFWwindow object that we can use for GLFW's functions
	w.window, err = glfw.CreateWindow(
		config.Window.DeviceRes.Width,
		config.Window.DeviceRes.Height,
		config.Window.Title,
		nil, nil)

	if err != nil {
		glfw.Terminate()
		return errors.New("Failed to create GLFW window")
	}

	w.window.SetSizeCallback(w.framebufferSizeCallback)
	// w.window.SetUserPointer()

	w.window.SetPos(config.Window.Position.X, config.Window.Position.Y)

	w.window.MakeContextCurrent()

	if config.Engine.ShowMonitorInfo {
		println("---------------------------- Monitor Info ---------------------------------------")
		monitor := glfw.GetPrimaryMonitor()
		mode := monitor.GetVideoMode()
		fmt.Printf("Monitor refresh rate: %d Hz\n", mode.RefreshRate)
		fmt.Printf("Monitor colors: RGB(%d, %d, %d)\n", mode.RedBits, mode.GreenBits, mode.BlueBits)
		fmt.Printf("Monitor size: %d x %d\n", mode.Width, mode.Height)
		pWidth, pHeight := monitor.GetPhysicalSize()

		fmt.Printf("Physical size: %d x %d\n", pWidth, pHeight)

		fbWidth, fbHeight := w.window.GetFramebufferSize()
		fmt.Printf("Framebuffer size: %d x %d\n", fbWidth, fbHeight)
		println("-------------------------------------------------------------------")
	}

	w.window.SetKeyCallback(w.keyCallback)

	if config.Window.LockToVSync {
		println("Locking to VSync")
		glfw.SwapInterval(1)
	}

	return nil
}

func (w *RWindow) initGL(config *config.Settings) error {
	println("Initializing OpenGL...")

	err := gl.Init()

	if err != nil {
		return nil
	}

	if config.Engine.ShowGLInfo {
		println("---------------------------- GL Info ---------------------------------------")
		fmt.Printf("Requesting OpenGL minimum of: %d.%d\n", config.Engine.GLMajorVersion, config.Engine.GLMinorVersion)

		version := gl.GoStr(gl.GetString(gl.VERSION))
		fmt.Printf("GL Version obtained: %s\n", version)

		vender := gl.GoStr(gl.GetString(gl.VENDOR))
		fmt.Printf("GL vender: %s\n", vender)

		renderer := gl.GoStr(gl.GetString(gl.RENDERER))
		fmt.Printf("GL renderer: %s\n", renderer)

		glslVersion := gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
		fmt.Printf("GLSL version: %s\n", glslVersion)

		var nrAttributes int32
		gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &nrAttributes)
		fmt.Printf("Max # of vertex attributes supported: %d\n", nrAttributes)
		println("-------------------------------------------------------------------")
	}

	return nil
}

func (w *RWindow) keyCallback(glfwW *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// println("key pressed")
	if key == glfw.KeyQ && action == glfw.Press {
		w.quitTriggered = true
	}
}

func (w *RWindow) framebufferSizeCallback(glfwW *glfw.Window, width int, height int) {
	fmt.Printf("Framebuffer re-size: %d x %d", width, height)
}
