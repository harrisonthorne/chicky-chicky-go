package main

import (
	"chicky-chicky-go/game"

	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	// I believe this ensures that our program always runs on the same process
    runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw: ", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Chicky Chicky Go", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(game.InputManager.KeyCallback)
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// vertex and fragment shaders
	// program, err := newProgram(vertexShader, fragmentShader)
	// if err != nil {
	// 	panic(err)
	// }

	// gl.UseProgram(program)
	gl.ClearColor(0, 0.5, 0.7, 1)

	lastUpdate := time.Now()

	for !window.ShouldClose() {
		deltaSeconds := float32(time.Now().Sub(lastUpdate) / time.Second)

		game.Logic(deltaSeconds)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		game.Render()

		window.SwapBuffers()
	}
}
