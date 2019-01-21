package world

import "chicky-chicky-go/game/space"

// Gravity is the gravity acceleration constant
const Gravity = -9.81

// PhysicalObject is an object with physics, position,
// velocity, and mass
type PhysicalObject struct {
	position space.Point
	velocity space.Velocity
    mass float32
	active   bool // Do we want physics to apply to this object?
}

// Physics calculates physics on the PhysicalObject p. 
func (p *PhysicalObject) Physics(float32 delta) {
    p.velocity.Dy += Gravity

    p.position.Add(p.velocity)
}
