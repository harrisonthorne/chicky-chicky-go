package render

var plainShader *Program

func initPlainShader() {

}

var plainShaderNames = ProgramAttrNames{
    PerspectiveMatrix: "perspective",
    CameraMatrix: "camera",
    ModelMatrix: "model",
    InColor: "color",
    OutColor: "outputColor",
}

// VertexPlainShaderSource is the source for the vertex shader of
// plain-color 3D programs
// {{{
var vertexPlainShaderSource = `
#version 330

uniform mat4 ` + plainShaderNames.PerspectiveMatrix + `;
uniform mat4 ` + plainShaderNames.CameraMatrix + `;
uniform mat4 ` + plainShaderNames.ModelMatrix + `;

in vec3 ` + plainShaderNames.InVertex + `;
in vec4 ` + plainShaderNames.InColor + `;
out vec4 fragColor;

void main() {
    ` + textureShaderNames.FragTexCoord + ` = ` + textureShaderNames.VertTexCoord + `;
    fragColor = ` + plainShaderNames.InColor + `
    gl_Position = ` + textureShaderNames.PerspectiveMatrix + ` * ` + textureShaderNames.CameraMatrix + ` * ` + textureShaderNames.ModelMatrix + ` * vec4(` + textureShaderNames.InVertex + `, 1);
}
` + "\x00" // any string being passed to OpenGL needs to terminate with the null character
// }}}

// FragmentPlainShaderSource is the source for the texture
// shader program
// {{{
var fragmentPlainShaderSource = `
#version 330

in vec4 fragColor;
out vec4 ` + textureShaderNames.OutColor + `;

void main() {
    ` + textureShaderNames.OutColor + ` = texture(` + textureShaderNames.TexSampler + `, ` + textureShaderNames.FragTexCoord + `);
}
` + "\x00"

// }}}

// vim: foldmethod=marker
