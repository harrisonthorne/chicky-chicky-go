package render

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

// Program is an OpenGL program
type Program struct {
	id                                 uint32
    Locations ProgramAttrLocations
}

// ProgramAttrNames holds names for program attributes
type ProgramAttrNames struct {
    PerspectiveMatrix string
    CameraMatrix string
    ModelMatrix string
    InVertex string
    OutVertex string
    InColor string
    OutColor string
    VertTexCoord string
    FragTexCoord string
    TexSampler string
	SpriteFrames string
	SpriteCurrentFrame string
}

// ID returns the program's OpenGL ID
func (p Program) ID() uint32 {
    return p.id
}

// NewProgram creates and returns a new OpenGL program.
func NewProgram(vertexShaderSource, fragmentShaderSource string, names ProgramAttrNames) (p *Program, err error) {
	p = new(Program)

	p.id, err = compileProgram(vertexShaderSource, fragmentShaderSource)
    if err != nil {
        return
    }

	p.Locations.perspectiveMatrix = gl.GetUniformLocation(p.id, gl.Str(names.PerspectiveMatrix+"\x00"))
	p.Locations.cameraMatrix = gl.GetUniformLocation(p.id, gl.Str(names.CameraMatrix+"\x00"))
	p.Locations.modelMatrix = gl.GetUniformLocation(p.id, gl.Str(names.ModelMatrix+"\x00"))
	p.Locations.spriteFrames = gl.GetUniformLocation(p.id, gl.Str(names.SpriteFrames+"\x00"))
	p.Locations.spriteCurrentFrame = gl.GetUniformLocation(p.id, gl.Str(names.SpriteCurrentFrame+"\x00"))

	gl.BindFragDataLocation(p.id, 0, gl.Str(names.OutColor+"\x00"))

    return
}

func compileProgram(vertexShaderStr, fragmentShaderStr string) (program uint32, err error) {
	vertexShader, err := compileShader(vertexShaderStr, gl.VERTEX_SHADER)
	if err != nil {
		return
	}

	fragmentShader, err := compileShader(fragmentShaderStr, gl.FRAGMENT_SHADER)
	if err != nil {
		return
	}

	program = gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		err = fmt.Errorf("failed to link program: %v", log)
		return
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return
}

func compileShader(source string, shaderType uint32) (shader uint32, err error) {
	shader = gl.CreateShader(shaderType)

	cString, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cString, nil)
	free() // MUST BE CALLED
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile shaderType %d: %v", shaderType, log)
	}

	return shader, nil
}
