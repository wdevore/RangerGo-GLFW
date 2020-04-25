package components

// InstantTransition animates a Scene on-to and off-of the Stage instantly
type InstantTransition struct {
	// The Scene this transition animates
	scene *Scene
}

// NewInstantTransition creates a instant Transition
func NewInstantTransition(sc *Scene) *InstantTransition {
	t := new(InstantTransition)
	t.scene = sc
	return t
}

// Step performs a single-step animation.
// Returns true immediately indicating transition is complete now.
func (it *InstantTransition) Step(dt float32) bool {
	// Translate the Scene out of view immediately.
	// This requires that we have access to the design dimensions.
	// it.scene.Node.SetPosition2Comp(0.0, 0.0)
	// vDimensions := it.scene.stage.GetSettings().Window.VirtualRes

	return true
}

// Start preps a transition for animation.
func (it *InstantTransition) Start() {

}

// Stop cleans up any resources
func (it *InstantTransition) Stop() {

}
