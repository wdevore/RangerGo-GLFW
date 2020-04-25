package main

import "github.com/wdevore/ranger/components"

// SplashScene shows the game logo amoung other things.
type SplashScene struct {
	// Splash is-a scene
	components.SceneBase
}

// NewSplashScene creates a new scene
func NewSplashScene() components.Scene {
	s := new(SplashScene)
	return s
}

// ------------------------------------------------------------------
// Scene interface
// ------------------------------------------------------------------

// Step takes a time step
func (ss *SplashScene) Step(dt float32) {

}

// GetInTransition generates a ...
func (ss *SplashScene) GetInTransition() components.Transition {
	return nil
}

// GetOutTransition generates a ...
func (ss *SplashScene) GetOutTransition() components.Transition {

	return nil
}

// IsAlive indicates that the Scene is moving onto the stage or is on the stage.
// func (ss *SplashScene) IsAlive() bool {
// 	return ss.alive
// }

// func (ss *SplashScene) SetAlive(live bool) {

// }
