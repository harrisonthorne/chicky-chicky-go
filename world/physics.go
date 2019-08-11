package world

import (
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"math"
)

type axis int

const (
	xAxis axis = iota
	yAxis
	zAxis
)

// fixes a collision given a breach value
func fix(p *PhysicalObject, breach maths.Vec3) {
	smallestNonZeroBreachAxis := xAxis
	smallestNonZeroBreachValue := breach.X

	if (breach.Y < smallestNonZeroBreachValue && breach.Y != 0) || smallestNonZeroBreachValue == 0 {
		smallestNonZeroBreachAxis = yAxis
		smallestNonZeroBreachValue = breach.Y
	} 
	
	if (breach.Z < smallestNonZeroBreachValue && breach.Z != 0) || smallestNonZeroBreachValue == 0  {
		smallestNonZeroBreachAxis = zAxis
		smallestNonZeroBreachValue = breach.Z
	}

	if smallestNonZeroBreachValue == 0 {
		return
	}

	switch smallestNonZeroBreachAxis {
	case xAxis:
		yPerX := p.velocity.Y / p.velocity.X
		zPerX := p.velocity.Z / p.velocity.X

		p.AddPosition(maths.Vec3{X: -breach.X})
		p.AddPosition(maths.Vec3{Y: -breach.X * yPerX})
		p.AddPosition(maths.Vec3{Z: -breach.X * zPerX})
	case yAxis:
		xPerY := p.velocity.X / p.velocity.Y
		zPerY := p.velocity.Z / p.velocity.Y

		p.AddPosition(maths.Vec3{X: -breach.Y * xPerY})
		p.AddPosition(maths.Vec3{Y: -breach.Y})
		p.AddPosition(maths.Vec3{Z: -breach.Y * zPerY})
	case zAxis:
		xPerZ := p.velocity.X / p.velocity.Z
		yPerZ := p.velocity.Y / p.velocity.Z

		p.AddPosition(maths.Vec3{X: -breach.Z * xPerZ})
		p.AddPosition(maths.Vec3{Y: -breach.Z * yPerZ})
		p.AddPosition(maths.Vec3{Z: -breach.Z})
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
	pi1 := maths.Vec3{
		X: vi1.X * m1, 
		Y: vi1.Y * m1, 
		Z: vi1.Z * m1,
	}
	pi2 := maths.Vec3{
		X: vi2.X * m2, 
		Y: vi2.Y * m2, 
		Z: vi2.Z * m2,
	}

	// kinetic energy = momentum * vel / 2 (who comes up
	// with this crap?)
	ei1 := maths.Vec3{
		X: pi1.X * vi1.X / 2, 
		Y: pi1.Y * vi1.Y / 2, 
		Z: pi1.Z * vi1.Z / 2,
	}
	ei2 := maths.Vec3{
		X: pi2.X * vi2.X / 2, 
		Y: pi2.Y * vi2.Y / 2, 
		Z: pi2.Z * vi2.Z / 2,
	}

	// so...
	// sum of final momentums = sum of initial momentums
	// and
	// sum of final kinetic energies = sum of initial kinetic energies

	// sum of momentum
	sp := maths.Vec3{
		X: pi1.X + pi2.X, 
		Y: pi1.Y + pi2.Y, 
		Z: pi1.Z + pi2.Z,
	}

	// sum of kinetic energy
	sKE := maths.Vec3{
		X: ei1.X + ei2.X, 
		Y: ei1.Y + ei2.Y, 
		Z: ei1.Z + ei2.Z,
	}

	// final velocity calculation
	p1.velocity.X, p2.velocity.X = getFinalVelocities(m1, m2, sp.X, sKE.X)
	p1.velocity.Y, p2.velocity.Y = getFinalVelocities(m1, m2, sp.Y, sKE.Y)
	p1.velocity.Z, p2.velocity.Z = getFinalVelocities(m1, m2, sp.Z, sKE.Z)
}

// calculate final velocities along one axis
func getFinalVelocities(m1, m2, sp, sKE float32) (first, second float32) {
	sqrt := float32(math.Sqrt(float64(m2 * (sp*sp*(2*m2-m1) - 2*sKE*m1*(m1-m2)))))
	vf2 := (m2*sp + sqrt) / (m2 * (m1 + m2))
	vf1 := (sp - m2*vf2) / m1

	return vf1, vf2
}
