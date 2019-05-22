package render

import (
    "github.com/go-gl/gl/v4.1-core/gl"
    mgl "github.com/go-gl/mathgl/mgl32"
    "github.com/harrisonthorne/chicky-chicky-go/maths"
)


// Camera is a camera
type Camera struct {
	position    maths.Vec3
	fov         float32
	perspective mgl.Mat4 // perspective matrix
	orientation mgl.Mat4 // stores position and rotation of the camera
	aspectRatio float32
}

// NewCamera constructs and returns a new Camera object
func NewCamera(position maths.Vec3, fov float32, aspectRatio float32) *Camera {
	c := &Camera{position: position, fov: fov, aspectRatio: aspectRatio}
	c.UpdatePerspectiveMatrix()
    c.orientation = mgl.LookAt(position.X, position.Y, position.Z, 0, 0, 0, 0, 1, 0)
	return c
}

// SetAspectRatio does what it's named for
func (c *Camera) SetAspectRatio(ratio float32) {
    c.aspectRatio = ratio
}

// SetPosition sets the position of Camera c
func (c *Camera) SetPosition(pos maths.Vec3) {
	c.position = pos
	c.UpdateOrientationMatrix()
}

// Position returns the position of Camera c
func (c *Camera) Position() maths.Vec3 {
	return c.position
}

// UpdateAllMatrices updates the camera's perspective and
// camera matrices
func (c *Camera) UpdateAllMatrices() {
	c.UpdatePerspectiveMatrix()
	c.UpdateOrientationMatrix()
}

// UpdatePerspectiveMatrix updates the perspective matrix
// of Camera c
func (c *Camera) UpdatePerspectiveMatrix() {
	c.perspective = mgl.Perspective(mgl.DegToRad(c.fov), c.aspectRatio, 0.1, 100)
}

// UpdateOrientationMatrix updates the orientation (position,
// rotation) matrix of Camera c
func (c *Camera) UpdateOrientationMatrix() mgl.Mat4 {
	return mgl.Translate3D(c.position.X, c.position.Y, c.position.Z);
}

// SetProgramAttributes sets the appropriate attributes in
// the current OpenGL program in use. The parameters passed
// in are the names of the attributes in the program.
func (c *Camera) SetProgramAttributes(p Program) {
    gl.UseProgram(p.id)
	gl.UniformMatrix4fv(p.Locations.PerspectiveMatrixLocation(), 1, false, &c.perspective[0])
	gl.UniformMatrix4fv(p.Locations.CameraMatrixLocation(), 1, false, &c.orientation[0])
}

// Matrices returns both the perspective and orientation
// matrices of Camera c.
func (c *Camera) Matrices() (perspective mgl.Mat4, orientation mgl.Mat4) {
	perspective = c.perspective
	orientation = c.orientation
	return
}
