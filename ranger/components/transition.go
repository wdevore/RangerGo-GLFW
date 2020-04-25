package components

// Transition animates a Scene on-to and off-of the Stage
type Transition interface {
	Start()
	Stop()
	Step(dt float32) bool
}
