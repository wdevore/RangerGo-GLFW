// Package config : Defines JSON configuration
package config

// Settings top parent
type Settings struct {
	Engine EngineObj
	Window WindowObj
	Camera CameraObj
	Font   FontObj
}

// EngineObj settings
type EngineObj struct {
	Enabled          bool
	LoopFor          int
	ShowConfig       bool
	ShowGLInfo       bool
	ShowMonitorInfo  bool
	ShowTimingInfo   bool
	ShowJoystickInfo bool
	GLMajorVersion   int
	GLMinorVersion   int
	FPSRefreshRate   float32
}

// WindowObj settings
type WindowObj struct {
	BitsPerPixel int
	LockToVSync  bool
	ClearColor   ColorObj
	VirtualRes   DimesionsObj
	DeviceRes    DimesionsObj
	FullScreen   bool
	Orientation  string
	Position     CoordinateObj
	Title        string
}

// CameraObj camera settings
type CameraObj struct {
	Centered bool
	View     ViewObj
}

// FontObj settings
type FontObj struct {
	Path         string
	Name         string
	Size         int
	Scale        float32
	CharsFromSet int
}

// ColorObj color
type ColorObj struct {
	R float32
	G float32
	B float32
	A float32
}

// DimesionsObj dimensions
type DimesionsObj struct {
	Height int
	Width  int
}

// CoordinateObj point
type CoordinateObj struct {
	X int
	Y int
}

// ViewObj view orientation
type ViewObj struct {
	X float32
	Y float32
	Z float32
}
