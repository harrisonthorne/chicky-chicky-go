package world

import (
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"math"
)

// Gravity is the gravity acceleration constant (m/s/s)
const Gravity = -9.81

// PhysicalObject is an object with physics, position,
// velocity, and mass
type PhysicalObject struct {
	frozen bool // if true, the PhysicalObject will not move

	velocity     maths.Vec3
	acceleration maths.Vec3
	mass         float32 // In kilograms.

	onGround    bool
	pushingWall bool
	atCeiling   bool

	wasOnGround    bool
	wasPushingWall bool
	wasAtCeiling   bool

	OnGroundHit  func()
	OnPush       func()
	OnCeilingHit func()

	Hitbox *maths.AABC // Hitbox for collision calculation, but not kill calculation
}

// Position returns the PhysicalObject's position
func (p *PhysicalObject) Position() maths.Vec3 {
	return p.Hitbox.CenterPos
}

func (p *PhysicalObject) AddPosition(v2 maths.Vec3) {
	p.Hitbox.CenterPos.Add(v2)
}

// SetPosition modifies the position of the PhysicalObject.
func (p *PhysicalObject) SetPosition(pos maths.Vec3) {
	p.Hitbox.CenterPos = pos
}

// Physics calculates physics on the PhysicalObject p.
func (p *PhysicalObject) Physics(delta float32) {
	// no physics if p is frozen
	if p.frozen {
		return
	}

	// gravity only applies if the PhysicalObject is not on
	// ground
	if !p.onGround {
		p.acceleration.Y += Gravity
	}

	p.acceleration.Scale(delta) // converts to velocity
	p.velocity.Add(p.acceleration)

	p.velocity.Scale(delta) // converts to displacement
	p.Hitbox.CenterPos.Add(p.velocity)

	// reset acceleration
	p.acceleration.X = 0
	p.acceleration.Y = 0
	p.acceleration.Z = 0
}

// ApplyForce applies a force, in Newtons, to the
// PhysicalObject. This is the only way to move a
// PhysicalObject in the game; velocity and acceleration are
// not publicly accessible.
func (p *PhysicalObject) ApplyForce(newtons maths.Vec3) {
	p.acceleration.X += newtons.X / p.mass
	p.acceleration.Y += newtons.Y / p.mass
	p.acceleration.Z += newtons.Z / p.mass
}

// StopMotion immediately stops the motion of the
// PhysicalObject. Velocity and acceleration are set to
// zero.
func (p *PhysicalObject) StopMotion() {
	p.velocity.X = 0
	p.velocity.Y = 0
	p.velocity.Z = 0
	p.acceleration.X = 0
	p.acceleration.Y = 0
	p.acceleration.Z = 0
}

// CollidesWith returns whether or not the Collider
// collides with another Collider
func (p *PhysicalObject) CollidesWith(other *PhysicalObject) bool {
	return p.Hitbox.CollidesWith(other.Hitbox)
}

// FixCollision fixes a collision between two
// PhysicalObjects. If both objects are actively subject to
// forces, momentum will take effect on both PhysicalObjects
// and force will be applied to both of  them
func (p *PhysicalObject) FixCollision(other *PhysicalObject) {
	if p.frozen || !p.CollidesWith(other) {
		return
	}

	// fix collisions
	switch {
	case !other.frozen:
		// fix both
		firstBreach := calculateBreach(p, other)
		secondBreach := calculateBreach(other, p)

		firstBreach.Scale(0.5)
		secondBreach.Scale(0.5)

		fix(p, firstBreach)
		fix(other, secondBreach)

		applyMomentum(p, other)
	case other.frozen:
		breach := calculateBreach(p, other)

		// fix p
		fix(p, breach)

		// now, we need to determine on which side of
		// `other` p is on. if it's on the top or bottom,
		// velocity stops on the y axis. if left or right, x
		// axis. first, re-calculate the breach now that
		// we've fixed the objects:
		breach = calculateBreach(p, other)

		// smallest breach determines which side the object is on
		minBreach := float32(math.Min(float64(breach.X), math.Min(float64(breach.Y), float64(breach.Z))))
		switch minBreach {
		case breach.X:
			p.velocity.X = 0
		case breach.Y:
			p.velocity.Y = 0
		case breach.Z:
			p.velocity.Z = 0
		}
	}
}

func calculateBreach(moving, static *PhysicalObject) (breach maths.Vec3) {
	// breach really depends on which direction the moving
	// PhysicalObject is travelling

	// calculate X
	switch {
	case moving.velocity.X > 0:
		breach.X = moving.Hitbox.CenterPos.X + moving.Hitbox.HalfSize.X - (static.Hitbox.CenterPos.X - static.Hitbox.HalfSize.X)
	case moving.velocity.X < 0:
		breach.X = moving.Hitbox.CenterPos.X - moving.Hitbox.HalfSize.X - (static.Hitbox.CenterPos.X + static.Hitbox.HalfSize.X)
	}

	// calculate Y
	switch {
	case moving.velocity.Y > 0:
		breach.Y = moving.Hitbox.CenterPos.Y + moving.Hitbox.HalfSize.Y - (static.Hitbox.CenterPos.Y - static.Hitbox.HalfSize.Y)
	case moving.velocity.Y < 0:
		breach.Y = moving.Hitbox.CenterPos.Y - moving.Hitbox.HalfSize.Y - (static.Hitbox.CenterPos.Y + static.Hitbox.HalfSize.Y)
	}

	// calculate Z
	switch {
	case moving.velocity.Z > 0:
		breach.Z = moving.Hitbox.CenterPos.Z + moving.Hitbox.HalfSize.Z - (static.Hitbox.CenterPos.Z - static.Hitbox.HalfSize.Z)
	case moving.velocity.Z < 0:
		breach.Z = moving.Hitbox.CenterPos.Z - moving.Hitbox.HalfSize.Z - (static.Hitbox.CenterPos.Z + static.Hitbox.HalfSize.Z)
	}
	return
}

// vim: foldmethod=syntax
