package blocks

import (
	"github.com/harrisonthorne/chicky-chicky-go/types"
	"github.com/harrisonthorne/chicky-chicky-go/textures"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Block block block block block block block block
type Block interface {
	types.Killable
	Render()
}

var cubeVAO, cubeVBO uint32

func init() {
	// create a vertex array object
	gl.GenVertexArrays(1, &cubeVAO)
	gl.BindVertexArray(cubeVAO)

	// create a vertex buffer object
	gl.GenBuffers(1, &cubeVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, cubeVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	
	// attribute pointers for the texture program
	program := textures.GetTextureProgram()

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
}

// TODO render a cube with the texture
func renderBlock(texture uint32) {
	gl.UseProgram(textures.GetTextureProgram())
	gl.BindVertexArray(cubeVAO)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
