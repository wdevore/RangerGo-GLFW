package components

import (
	"github.com/wdevore/ranger/rmath"
	"github.com/wdevore/utilities"
)

// SceneManager manages Scenes.
type SceneManager struct {
	scenes *utilities.Stack

	model rmath.Matrix4

	activeScene   Scene
	outgoingScene Scene

	// Each Transition controls a different Scene.
	// The Transition that animates a Scene on-to the Stage
	transitionIn Transition
	// The Transition that animates a Scene off-of the Stage.
	// If there is only one Scene on the stack then this transition is nil
	transitionOut Transition
}

// NewSceneManager creates a SceneManager with a default maximum of
// 10 Scenes allowed
func NewSceneManager(max int) *SceneManager {
	sm := new(SceneManager)
	sm.scenes = utilities.NewStack(10)
	return sm
}

// Push tells the SceneManager to start interacting with the new Scene.
func (sm *SceneManager) Push(newSc Scene) {
	// If there an active scene then get its outward transition and
	// start it.
	if sm.activeScene != nil {
		sm.transitionOut = newSc.GetOutTransition()
		sm.transitionOut.Start()
	}

	// The stack top becomes the outgoing scene
	var isScene bool

	sm.outgoingScene, isScene = sm.scenes.Pop().(Scene)

	if !isScene {
		panic("Top of scene stack contained something other than a Scene")
	}

	// Now push the new scene and make it active
	sm.Push(newSc)
	sm.activeScene = newSc

	// To animate the scene on-to the stage we ask the scene for a
	// transition
	sm.transitionIn = newSc.GetInTransition()
	sm.transitionIn.Start()

	newSc.SetAlive(true)
}

// Step steps any transitions or active scenes.
func (sm *SceneManager) Step(dt float32) bool {
	if sm.scenes.IsEmpty() {
		println("SceneManager.step: no more scenes to visit.")
		return false
	}

	if sm.transitionIn != nil {
		complete := sm.transitionIn.Step(dt)
		if complete {
			sm.transitionIn = nil
		}
	}

	if sm.transitionOut != nil {
		complete := sm.transitionOut.Step(dt)
		if complete {
			sm.transitionOut = nil
			// Pull scene from stack
		}
	}

	// If there is no active scene then attempt to pull one from
	// the stack.

	return true
}
