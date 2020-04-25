// Package components make up the basic core of the scene graph
package components

import (
	"github.com/wdevore/ranger/rmath"
)

// Node is an abstract core component of the Scene graph that
// should be embedded in components, i.e. is-a relationship.
type Node struct {
	Name           string
	Visible        bool
	parent         *Node
	cleanOnRelease bool

	/*
	 * [tag]s are used to identify [Node]s by Ids. They both can be
	 * handy for developement or for finding [Node]s.
	 *
	 * Ex:
	 * const TAG_PLAYER = 1;
	 * node1.tag = TAG_PLAYER;
	 * node2.tag = TAG_MONSTER;
	 * node3.tag = TAG_BOSS;
	 */

	Tag int

	//---------------------------------------------------------------------
	// Transforms
	//---------------------------------------------------------------------
	// Optimization flags to track transforms
	transformDirty bool
	inverseDirty   bool

	// Does this [Node] manage its own transform matrix, for example,
	// Zoom nodes manage their own.
	managedTransform bool

	transform    rmath.Matrix4
	invTransform rmath.Matrix4

	//---------------------------------------------------------------------
	// Discrete transform properties
	//---------------------------------------------------------------------
	position rmath.Vector3
	rotation float32
	scale    rmath.Vector3
}

// initialize should be called by a parent container.
func (n *Node) initialize() {
	// By default all nodes are dirty, this way their transforms are computed.
	n.SetDirty()
	n.Tag = -1
	n.scale.Z = 1.0
}

//---------------------------------------------------------------------
// Base properties
//---------------------------------------------------------------------

// SetDirty marks a node dirty
func (n *Node) SetDirty() {
	// Only mark Nodes that DON'T manage their own transforms.
	if n.managedTransform {
		return
	}

	n.transformDirty = true
	n.inverseDirty = true
}

//---------------------------------------------------------------------
// Rotation
//---------------------------------------------------------------------

// SetRotate sets rotation property given in radians
func (n *Node) SetRotate(angle float32) {
	n.rotation = angle
	n.transformDirty = true
}

// RotateByDegrees sets rotation property given in degrees
func (n *Node) RotateByDegrees(angle float32) {
	n.rotation = rmath.ToRadians(angle)
	n.transformDirty = true
}

// RotateBy increments rotation property by radians
func (n *Node) RotateBy(angle float32) {
	n.rotation += angle
	n.transformDirty = true
}

//---------------------------------------------------------------------
// Position
//---------------------------------------------------------------------

// SetPosition3Comp sets positional property
func (n *Node) SetPosition3Comp(x, y, z float32) {
	n.position.Set3Components(x, y, z)
	n.transformDirty = true
}

// SetPosition2Comp sets positional property
func (n *Node) SetPosition2Comp(x, y float32) {
	n.position.Set2Components(x, y)
	n.transformDirty = true
}

// SetPositionByVector sets positional property
func (n *Node) SetPositionByVector(v *rmath.Vector3) {
	n.position.Set3Components(v.X, v.Y, v.Z)
	n.transformDirty = true
}

// MoveBy increments positional property
func (n *Node) MoveBy(v *rmath.Vector3) {
	n.position.Add(v)
	n.transformDirty = true
}

// MoveBy2Comp increments positional property
func (n *Node) MoveBy2Comp(x, y float32) {
	n.position.Add2Components(x, y)
	n.transformDirty = true
}

//---------------------------------------------------------------------
// Scale
//---------------------------------------------------------------------

// SetScale set scale property
func (n *Node) SetScale(sv rmath.Vector3) {
	n.scale.Set3Components(sv.X, sv.Y, sv.Z)
	n.transformDirty = true
}

// ScaleBy increments scale property
func (n *Node) ScaleBy(s float32) {
	n.scale.ScaleBy(s)
	n.transformDirty = true
}

// ScaleBy2Components increments scale property
func (n *Node) ScaleBy2Components(s *rmath.Vector3) {
	n.position.ScaleBy2Components(s.X, s.Y)
	n.transformDirty = true
}

//---------------------------------------------------------------------
// Graph: Traversal
//---------------------------------------------------------------------

// Visit iterates through each node. Override.
func (n *Node) Visit(dt float32, modelT *rmath.Matrix4) bool {
	return true
}

// Render draws the node. Default is nothing
func (n *Node) Render() {

}

//---------------------------------------------------------------------
// Transforms
//---------------------------------------------------------------------

// CalcTransform computes this Node's transform from properties.
func (n *Node) CalcTransform() *rmath.Matrix4 {
	if n.transformDirty {
		n.transform.SetTranslateByVector(&n.position)

		// We Scale "before" Rotate because Matrix4 uses pre-multiply
		if n.scale.X != 1.0 || n.scale.Y != 1.0 {
			n.transform.ScaleBy(&n.scale)
		}

		if n.rotation != 0.0 {
			n.transform.RotateBy(n.rotation)
		}

		n.transformDirty = false
	}

	return &n.transform
}

// NodeToWorldTransform computes the "worldT" matrix to map from node-space to world-space.
func (n *Node) NodeToWorldTransform(worldT *rmath.Matrix4, relativeRootNode *Node) {
	// Start with this node's tranform.
	worldT.Set(n.CalcTransform())

	// Iterate "upwards" starting with the child towards the parents
	// starting with this child's parent.
	p := n.parent

	for p != nil {
		parentT := p.CalcTransform() // Non-pooled object

		// Because we are iterating upwards we need to post-multiply each
		// child. Ex: [child] x [parent]
		// ----------------------------------------------------------
		//           [worldT] x [parent]
		//                  |
		//                  v
		//                 [worldT] x [parent]
		//                        |
		//                        v
		//                      [worldT] x [parent]
		//                      ...
		//
		// In other words the child is mutiplied "into" the parent.
		worldT.PostMultiply(parentT)

		if p == relativeRootNode {
			break
		}

		// Move upwards towards next parent.
		p = p.parent
	}
}
