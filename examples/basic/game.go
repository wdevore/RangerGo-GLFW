package main

import (
	"os"
	"runtime"

	ranger "github.com/wdevore/ranger"
)

func main() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	// println("Launching game...")
	game := new(Game)

	engine := ranger.NewEngine(game)
	err := engine.Launch()

	if err != nil {
		os.Stderr.WriteString("### Engine failed to launch ###\n")
		os.Stderr.WriteString(err.Error())
	}

	defer engine.Stop()
}

// Game is for the developer to implement GameShell's Configure callback method.
type Game struct {
	engine *ranger.Engine

	// Various Scenes that make up the game.
}

// Configure is called by the Engine when it's ready for
// the developer to configure their game.
func (g *Game) Configure(e *ranger.Engine) bool {
	println("Configuring game...")

	g.engine = e

	// Create Scenes and Layers using the Engine API.
	splash

	return true
}
