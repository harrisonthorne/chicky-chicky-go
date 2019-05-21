package blocks

import (
	"github.com/harrisonthorne/chicky-chicky-go/render"
	"github.com/harrisonthorne/chicky-chicky-go/types"
	"github.com/harrisonthorne/chicky-chicky-go/utils"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Block block block block block block block block
type Block interface {
	types.Killable
	types.Renderable
}

var cubeVAO, cubeVBO uint32

// InitGL initializes OpenGL-specific functionality for the
// blocks package.
func InitGL() {
	cubeVAO, cubeVBO = utils.NewTextureVAO(render.TextureProgram(), &cubeVertices)
}

func renderBlock(texture uint32) {
	gl.UseProgram(render.TextureProgram().ID())
	gl.BindVertexArray(cubeVAO)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// six faces times two triangles per face times three
	// vertices per triangle
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
