// Package ranger in the top most package providing the GameShell interface
package ranger

// GameShell is the interface by which the Engine coordinates with a user's Game
// during bootstrapping.
type GameShell interface {
	// Launch starts the engine
	// Launch()
	// Configure allows the user to perform port engine start configurations.
	Configure(*Engine) bool
}
