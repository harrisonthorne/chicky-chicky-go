package maths

import "math"

// AABB is an Axis-Aligned Bounding Box. it is used to check
// for collisions in collision detection.
type AABB struct {
	CenterPos Vec2
	HalfSize  Vec2
}

// CollidesWith returns true if the AABB is touching the
// other AABB
func (a *AABB) CollidesWith(other *AABB) bool {
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

