package render

var textureShader *Program

// TextureProgram returns the texture program
func TextureProgram() Program {
    return *textureShader
}

func initTextureShader() {
    var err error
    textureShader, err = NewProgram(vertexTextureShaderSource, fragmentTextureShaderSource, textureShaderNames)
    if err != nil {
        panic(err)
    }
}

var textureShaderNames = ProgramAttrNames{
    PerspectiveMatrix: "perspective",
    CameraMatrix: "camera",
    ModelMatrix: "model",
    InVertex: "vert",
    VertTexCoord: "vertTexCoord",
    FragTexCoord: "fragTexCoord",
    TexSampler: "tex",
    OutColor: "outputColor",
}

// VertexTextureShaderSource is the source for the vertex shader of
// all 3D programs
// {{{
var vertexTextureShaderSource = `
#version 330

uniform mat4 ` + textureShaderNames.PerspectiveMatrix + `;
uniform mat4 ` + textureShaderNames.CameraMatrix + `;
uniform mat4 ` + textureShaderNames.ModelMatrix + `;

in vec3 ` + textureShaderNames.InVertex + `;
in vec2 ` + textureShaderNames.VertTexCoord + `;

out vec2 ` + textureShaderNames.FragTexCoord + `;

void main() {
    ` + textureShaderNames.FragTexCoord + ` = ` + textureShaderNames.VertTexCoord + `;
    gl_Position = ` + textureShaderNames.PerspectiveMatrix + ` * ` + textureShaderNames.CameraMatrix + ` * ` + textureShaderNames.ModelMatrix + ` * vec4(` + textureShaderNames.InVertex + `, 1);
}
` + "\x00" // any string being passed to OpenGL needs to terminate with the null character
// }}}

// FragmentTextureShaderSource is the source for the texture
// shader program
// {{{
var fragmentTextureShaderSource = `
#version 330

uniform sampler2D ` + textureShaderNames.TexSampler + `;

in vec2 ` + textureShaderNames.FragTexCoord + `;

out vec4 ` + textureShaderNames.OutColor + `;

void main() {
    vec4 texColor = texture(` + textureShaderNames.TexSampler + `, ` + textureShaderNames.FragTexCoord + `);
    if (texColor.a < 0.05)
        discard;

    ` + textureShaderNames.OutColor + ` = texColor;
}
` + "\x00"

// }}}

// vim: foldmethod=marker
