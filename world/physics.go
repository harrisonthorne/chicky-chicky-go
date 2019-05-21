package world

import (
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"math"
)

// fixes a collision given a breach value
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

func applyMomentum(p1, p2 *PhysicalObject) {
	// velocity
	vi1 := p1.velocity
	vi2 := p2.velocity

	// mass
	m1 := p1.mass
	m2 := p2.mass

	// momentum = velocity * mass
	pi1 := maths.Vec2{X: vi1.X * m1, Y: vi1.Y * m1}
	pi2 := maths.Vec2{X: vi2.X * m2, Y: vi2.Y * m2}

	// kinetic energy = momentum * vel / 2 (who comes up
	// with this crap?)
	ei1 := maths.Vec2{X: pi1.X * vi1.X / 2, Y: pi1.Y * vi1.Y / 2}
	ei2 := maths.Vec2{X: pi2.X * vi2.X / 2, Y: pi2.Y * vi2.Y / 2}

	// so...
	// sum of final momentums = sum of initial momentums
	// and
	// sum of final kinetic energies = sum of initial kinetic energies

	// sum of momentum
	sp := maths.Vec2{X: pi1.X + pi2.X, Y: pi1.Y + pi2.Y}

	// sum of kinetic energy
	sKE := maths.Vec2{X: ei1.X + ei2.X, Y: ei1.Y + ei2.Y}

	// final velocity calculation
	p1.velocity.X, p2.velocity.X = getFinalVelocities(m1, m2, sp.X, sKE.X)
	p1.velocity.Y, p2.velocity.X = getFinalVelocities(m1, m2, sp.X, sKE.X)
}

// calculate final velocities along one axis
func getFinalVelocities(m1, m2, sp, sKE float32) (first, second float32) {
	sqrt := float32(math.Sqrt(float64(m2 * (sp*sp*(2*m2-m1) - 2*sKE*m1*(m1-m2)))))
	vf2 := (m2*sp + sqrt) / (m2 * (m1 + m2))
	vf1 := (sp - m2*vf2) / m1

	return vf1, vf2
}
