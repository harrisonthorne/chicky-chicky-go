package maths

// Vec2 is a vector with two components. This is intended
// for use with velocity, position, and acceleration.
type Vec2 struct {
    X float32
    Y float32
}

// Add adds the components of v2 to the matching
// components of Vec2 v1
func (v *Vec2) Add(v2 Vec2) {
    v.X += v2.X
    v.Y += v2.Y
}

// Scale scales the vector
func (v *Vec2) Scale(scalar float32) {
    v.X *= scalar
    v.Y *= scalar
}

// Vec3 is a vector with three components. This is intended
// for use with three-dimensional velocity, position, and
// acceleration.
type Vec3 struct {
    X float32
    Y float32
    Z float32
}

// Add adds the components of v2 to the matching
// components of Vec2 v1
func (v *Vec3) Add(v2 Vec3) {
    v.X += v2.X
    v.Y += v2.Y
    v.Z += v2.Z
}
