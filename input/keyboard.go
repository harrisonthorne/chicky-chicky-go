package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// KeyboardListener listens for keyboard input.
// Its functions are called from glfw's callback.
type KeyboardListener interface {
	KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey)
	KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey)
	KeyRepeat(key glfw.Key, scancode int, mods glfw.ModifierKey)
}

