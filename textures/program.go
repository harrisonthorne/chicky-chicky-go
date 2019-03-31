package textures

import (
	"github.com/harrisonthorne/chicky-chicky-go/shaders"
)

var textureProgram uint32

// GetTextureProgram returns the app's texture program,
// used for drawing images or sprites
func GetTextureProgram() uint32 {
	return textureProgram
}

func init() {
	var err error
	textureProgram, err = shaders.NewProgram(vertexTextureShader, fragmentTextureShader)
	if err != nil {
		panic(err)
	}
}

// vertexTextureShader {{{
const vertexTextureShader = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00" // any string being passed to OpenGL needs to terminate with the null character
// }}}

// fragmentTextureShader {{{
const fragmentTextureShader = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

// }}}

// vim: foldmethod=marker
