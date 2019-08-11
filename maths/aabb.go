package maths

import "math"

// AABC is an Axis-Aligned Bounding Cube. it is used to check
// for collisions in collision detection.
type AABC struct {
	CenterPos Vec3
	HalfSize  Vec3
}

// CollidesWith returns true if the AABC is touching the
// other AABB
func (a *AABC) CollidesWith(other *AABC) bool {
    // AABBs are in collision with each other if and only
    // if, on all axes, the distance between the center of
    // the AABBs is less than the sum of half of the size of
    // either AABB.

    // Distance between centers
    xCenterDelta := math.Abs(float64(other.CenterPos.X - a.CenterPos.X))
    yCenterDelta := math.Abs(float64(other.CenterPos.Y - a.CenterPos.Y))
    zCenterDelta := math.Abs(float64(other.CenterPos.Z - a.CenterPos.Z))

    // Sum of half sizes
    xHalfSizeSum := other.HalfSize.X + a.HalfSize.X
    yHalfSizeSum := other.HalfSize.Y + a.HalfSize.Y
    zHalfSizeSum := other.HalfSize.Z + a.HalfSize.Z

    // On which axes do the AABBs collide?
    xCollision := float32(xCenterDelta) < xHalfSizeSum
    yCollision := float32(yCenterDelta) < yHalfSizeSum
    zCollision := float32(zCenterDelta) < zHalfSizeSum
    
    return xCollision && yCollision && zCollision
}

