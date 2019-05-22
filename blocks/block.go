package blocks

import (
	"github.com/harrisonthorne/chicky-chicky-go/render"
	"github.com/harrisonthorne/chicky-chicky-go/types"
	"github.com/harrisonthorne/chicky-chicky-go/utils"
	"github.com/go-gl/gl/v4.1-core/gl"

	mgl "github.com/go-gl/mathgl/mgl32"
)

// Block block block block block block block block
type Block interface {
	types.Killable
	types.Renderable
	SetGridPos(row, col int)
	GridPos() (row, col int)
}

var cubeVAO, cubeVBO uint32

// InitGL initializes OpenGL-specific functionality for the
// blocks package.
func InitGL() {
	cubeVAO, cubeVBO = utils.NewTextureVAO(render.TextureProgram(), &cubeVertices)
}

var rotation mgl.Vec3

func renderBlock(c *render.Camera, texture uint32) {
	rotation[0] += 0.01
	rotation[1] += 0.02
	rotation[2] += 0.03
	mat := mgl.HomogRotate3DY(rotation[2])
	// mat = mat.Mul4(mgl.HomogRotate3DY(rotation[1]))
	// mat = mat.Mul4(mgl.HomogRotate3DZ(rotation[2]))

	println(texture)
	gl.UseProgram(render.TextureProgram().ID())
	gl.BindVertexArray(cubeVAO)

	c.SetProgramAttributes(render.TextureProgram())

	modelAttrLocation := render.TextureProgram().Locations.ModelMatrixLocation()
	gl.UniformMatrix4fv(modelAttrLocation, 1, false, &mat[0])

	textureUniform := render.TextureProgram().Locations.TextureLocation()
	gl.Uniform1i(textureUniform, 0) // number bound here must match the active texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// six faces times two triangles per face times three
	// vertices per triangle
	gl.BindVertexArray(cubeVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
