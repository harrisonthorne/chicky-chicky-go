package utils

import (
	"github.com/harrisonthorne/chicky-chicky-go/textures"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// NewTextureVAO creates and return a new vao and vbo built with
// the vertices passed in. Each vertex should consist of
// five float32 values. The first three values are XYZ
// coordinates. The last two are UV coordinates for 
func NewTextureVAO(vertices *[]float32) (vao, vbo uint32) {
	// create a vertex array object
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// create a vertex buffer object
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*4, gl.Ptr(*vertices), gl.STATIC_DRAW)

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
	return
}

func setAttribute(vao, program uint32, attribName string, value interface{}) {

}
