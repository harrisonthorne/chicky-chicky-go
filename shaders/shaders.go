package shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// NewProgram creates and returns a new program
func NewProgram(vertexShaderStr, fragmentShaderStr string) (program uint32, err error) {
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
