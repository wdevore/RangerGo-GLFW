package components

import "github.com/wdevore/ranger"

// Scene represents nodes on stage.
type Scene interface {
	Step(dt float32)
	GetInTransition() Transition
	GetOutTransition() Transition

	IsAlive() bool
	SetAlive(live bool)
}

// SceneBase is a common base for typical scenes.
type SceneBase struct {
	// A scene is-a Node
	Node

	stage *ranger.Stage

	alive bool
}

// IsAlive indicates that the Scene is moving onto the stage or is on the stage.
func (sb *SceneBase) IsAlive() bool {
	return sb.alive
}

// SetAlive either marks the Scene as live = true or dead = false
func (sb *SceneBase) SetAlive(live bool) {
	sb.alive = live
}

// NewScene creates a Scene
// func NewScene(st *ranger.Stage) *Scene {
// 	s := new(Scene)
// 	s.stage = st
// 	return s
// }

// // GetInTransition returns a default "instant" transition.
// func (sc *Scene) GetInTransition() Transition {
// 	it := NewInstantTransition(sc)
// 	return it
// }
