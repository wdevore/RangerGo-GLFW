package components

import "github.com/wdevore/ranger/rmath"

// Group is a Node that collects other Nodes.
type Group struct {
	Node

	children []*Node

	highestZOrder      int
	hasNegativeZOrders bool
}

// NewGroup creates a new Group and initializes its Base Node
func (g *Group) NewGroup() *Group {
	gp := new(Group)
	gp.initialize()
	return gp
}

func (g *Group) initialize() {
	g.Node.initialize() // super
	g.highestZOrder = 10000000
	g.hasNegativeZOrders = false
}

// ---------------------------------------------------------------
// Node overrides
// ---------------------------------------------------------------

// Visit iterates through each node. Override.
func (g *Group) Visit(dt float32, modelT *rmath.Matrix4) bool {
	return true
}

// Render draws the node. Default is nothing
func (g *Group) Render() {
	// Typically an aabbox surrounding the group
}
