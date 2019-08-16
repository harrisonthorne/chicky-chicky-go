package blocks

const halfBW = BlockWidth / 2

// TODO fix uv coordinates
var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-halfBW, -halfBW, -halfBW, 0.5, 0.5,
	halfBW, -halfBW, -halfBW, 1.0, 0.5,
	-halfBW, -halfBW, halfBW, 0.5, 0.0,

	halfBW, -halfBW, -halfBW, 1.0, 0.5,
	halfBW, -halfBW, halfBW, 1.0, 0.0,
	-halfBW, -halfBW, halfBW, 0.5, 0.0,

	// Top
	-halfBW, halfBW, -halfBW, 0.0, 0.5,
	-halfBW, halfBW, halfBW, 0.0, 0.0,
	halfBW, halfBW, -halfBW, 0.5, 0.5,

	halfBW, halfBW, -halfBW, 0.5, 0.5,
	-halfBW, halfBW, halfBW, 0.0, 0.0,
	halfBW, halfBW, halfBW, 0.5, 0.0,

	// Front
	-halfBW, -halfBW, halfBW, 0.0, 1.0,
	halfBW, -halfBW, halfBW, 0.5, 1.0,
	-halfBW, halfBW, halfBW, 0.0, 0.5,

	halfBW, -halfBW, halfBW, 0.5, 1.0,
	halfBW, halfBW, halfBW, 0.5, 0.5,
	-halfBW, halfBW, halfBW, 0.0, 0.5,

	// Back
	-halfBW, -halfBW, -halfBW, 0.5, 1.0,
	-halfBW, halfBW, -halfBW, 0.5, 0.5,
	halfBW, -halfBW, -halfBW, 0.0, 1.0,

	halfBW, -halfBW, -halfBW, 0.0, 1.0,
	-halfBW, halfBW, -halfBW, 0.5, 0.5,
	halfBW, halfBW, -halfBW, 0.0, 0.5,

	// Left
	-halfBW, -halfBW, halfBW, 1.0, 1.0,
	-halfBW, halfBW, -halfBW, 0.5, 0.5,
	-halfBW, -halfBW, -halfBW, 0.5, 1.0,

	-halfBW, -halfBW, halfBW, 1.0, 1.0,
	-halfBW, halfBW, halfBW, 1.0, 0.5,
	-halfBW, halfBW, -halfBW, 0.5, 0.5,

	// Right
	halfBW, -halfBW, halfBW, 0.5, 1.0,
	halfBW, -halfBW, -halfBW, 1.0, 1.0,
	halfBW, halfBW, -halfBW, 1.0, 0.5,

	halfBW, -halfBW, halfBW, 0.5, 1.0,
	halfBW, halfBW, -halfBW, 1.0, 0.5,
	halfBW, halfBW, halfBW, 0.5, 0.5,
}
