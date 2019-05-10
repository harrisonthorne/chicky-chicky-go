package render

// ProgramAttrLocations holds locations of attributes in a
// program
type ProgramAttrLocations struct {
    perspectiveMatrix int32
    cameraMatrix int32
    modelMatrix int32
    texture int32
}

// PerspectiveMatrixLocation returns the perspective matrix
// attribute location
func (p ProgramAttrLocations) PerspectiveMatrixLocation() int32 {
    return p.perspectiveMatrix
}

// CameraMatrixLocation returns the camera matrix attribute
// location
func (p ProgramAttrLocations) CameraMatrixLocation() int32 {
    return p.cameraMatrix
}

// ModelMatrixLocation returns the model matrix attribute
// location
func (p ProgramAttrLocations) ModelMatrixLocation() int32 {
    return p.modelMatrix
}

// TextureLocation returns the texture attribute location
func (p ProgramAttrLocations) TextureLocation() int32 {
    return p.texture
}
