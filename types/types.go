package types

import (
	"chicky-chicky-go/maths"

    "math"
)

// AABB is an Axis-Aligned Bounding Box. it is used to check
// for collisions in collision detection.
type AABB struct {
	CenterPos maths.Vec2
	HalfSize  maths.Vec2
}

// CollidesWith returns true if the AABB is touching the
// other AABB
func (a *AABB) CollidesWith(other AABB) bool {
    // AABBs are in collision with each other if and only
    // if, on all axes, the distance between the center of
    // the AABBs is less than the sum of half of the size of
    // either AABB.

    // Distance between centers
    xCenterDelta := math.Abs(float64(other.CenterPos.X - a.CenterPos.X))
    yCenterDelta := math.Abs(float64(other.CenterPos.Y - a.CenterPos.Y))

    // Sum of half sizes
    xHalfSizeSum := other.HalfSize.X + a.HalfSize.X
    yHalfSizeSum := other.HalfSize.Y + a.HalfSize.Y

    // On which axes do the AABBs collide?
    xCollision := float32(xCenterDelta) < xHalfSizeSum
    yCollision := float32(yCenterDelta) < yHalfSizeSum
    
    return xCollision && yCollision
}

// Animatable is an interface that can be added to objects
// that animate. It calls its Animate(delta) method during
// every logical loop to compute whether or not the
// animation should advance to the next frame, which frame
// the animation should be on, or the animation
// itself.
type Animatable interface {
	Animate(delta float32)
}

// Ignitable can be added to objects that can be ignited.
// Such objects can burn any objects with the Burnable
// interface attached.
type Ignitable interface {
	Ignite()
}

// Burnable can be added to objects that can be burned.  Its
// Burn() method will be called when an offending object is
// placed next to it.
type Burnable interface {
	Burn()
}

// Killable is here to serve as an embed in blocks or
// characters. in other words, anything that can be killed.
type Killable struct {
	IKillable
	Lifespan int     // Max health the Killable can have
	Health   float32 // Health remaining on the Killable
}

// IKillable is implemented in aforementioned objects that
// can be killed
type IKillable interface {
	// Called when the Killable is hit. Returns any items that
	// the Killable might drop when hit.
	Hit(with Item, power int) []Item

	// Called when the Killable should be killed. Returns any
	// items that might be dropped with the Killable dies.
	Kill() []Item
}

// Renderable can be implemented by objects that are to be
// rendered in the game
type Renderable interface {
	Render()
}

// PhysicalObject is an object with physics, position, size,
// and velocity
type PhysicalObject struct {
	position maths.Vec2
	velocity maths.Vec2

	Acceleration   maths.Vec2
	OnGround       bool
	PushingWall    bool
	AtCeiling      bool
	WasOnGround    bool
	WasPushingWall bool
	WasAtCeiling   bool
}

// Physics performs physics calculations on the PhysicalObject
func (p *PhysicalObject) Physics(delta float32) {

}
