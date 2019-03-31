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

	velocity     maths.Vec2
	acceleration maths.Vec2
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

	Hitbox *maths.AABB // Hitbox for collision calculation, but not kill calculation
}

// Position returns the PhysicalObject's position
func (p *PhysicalObject) Position() maths.Vec2 {
	return p.Hitbox.CenterPos
}

// SetPosition modifies the position of the PhysicalObject.
func (p *PhysicalObject) SetPosition(pos maths.Vec2) {
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
}

// ApplyForce applies a force, in Newtons, to the
// PhysicalObject. This is the only way to move a
// PhysicalObject in the game; velocity and acceleration are
// not publicly accessible.
func (p *PhysicalObject) ApplyForce(newtons maths.Vec2) {
	p.acceleration.X += newtons.X / p.mass
	p.acceleration.Y += newtons.Y / p.mass
}

// StopMotion immediately stops the motion of the
// PhysicalObject. Velocity and acceleration are set to
// zero.
func (p *PhysicalObject) StopMotion() {
	p.velocity.X = 0
	p.velocity.Y = 0
	p.acceleration.X = 0
	p.acceleration.Y = 0
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
		// TODO use momentum to calculate new velocities?
		firstBreach := calculateBreach(p, other)
		secondBreach := calculateBreach(other, p)

		firstBreach.Scale(0.5)
		secondBreach.Scale(0.5)

		fix(p, firstBreach)
		fix(other, secondBreach)
	case other.frozen:
		breach := calculateBreach(p, other)

		// fix p
		fix(p, breach)

		// now, we need to determine on which side of other
		// p is on. if it's on the top or bottom, velocity
		// stops on the y axis. if left or right, x axis
		if float64(p.Hitbox.HalfSize.X + other.Hitbox.HalfSize.X) < math.Abs(float64(p.Hitbox.CenterPos.X - other.Hitbox.CenterPos.X)) {
			// object is on top or bottom, so set y
			// velocity to zero
			p.velocity.Y = 0
		} else if float64(p.Hitbox.HalfSize.Y + other.Hitbox.HalfSize.Y) < math.Abs(float64(p.Hitbox.CenterPos.Y - other.Hitbox.CenterPos.Y)) {
			// object is on top or bottom, so set y
			// velocity to zero
			p.velocity.X = 0
		}
	}
}

// breach really depends on which direction the moving
// PhysicalObject is travelling
func calculateBreach(moving, static *PhysicalObject) (breach maths.Vec2) {
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
	return
}

func fix(p *PhysicalObject, breach maths.Vec2) {
	if p.velocity.X == 0 {
		p.Hitbox.CenterPos.Y += -breach.Y
		return
	}
	velSlope := p.velocity.Y / p.velocity.X

	// try along x axis
	dxOnX := -breach.X
	dyOnX := -breach.X * velSlope
	dOnX := math.Sqrt(math.Pow(float64(dxOnX), 2) + math.Pow(float64(dyOnX), 2))

	// try along x axis
	dxOnY := -breach.Y / velSlope
	dyOnY := -breach.Y
	dOnY := math.Sqrt(math.Pow(float64(dxOnY), 2) + math.Pow(float64(dyOnY), 2))

	if dOnX < dOnY {
		p.Hitbox.CenterPos.X += dxOnX
		p.Hitbox.CenterPos.Y += dyOnX
	} else {
		p.Hitbox.CenterPos.X += dxOnY
		p.Hitbox.CenterPos.Y += dyOnY
	}
}
