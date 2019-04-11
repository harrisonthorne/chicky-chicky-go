package game

import (
	// "github.com/go-gl/gl/v4.1-core/gl"
	// "github.com/go-gl/glfw/v3.2/gl"
	// "github.com/go-gl/mathgl/mgl32"
	"github.com/harrisonthorne/chicky-chicky-go/characters"
	"github.com/harrisonthorne/chicky-chicky-go/input"
)

// InputManager is the game's main input manager for
// keyboard and mouse handling
var InputManager input.Manager

var characterInControl characters.Controllable

func init() {
	println("adding keyboard listener")
	InputManager.AddKeyboardListener(&TheChicken)
}

// Logic performs logic for the game. This includes movement, physics,
// clocks, animation, etc
func Logic(delta float32) {
	TheChicken.Animate(delta)

}

// Render renders the game.
func Render() {
	
}

type keyListener input.KeyboardListener

func (k *keyListener) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {

}

func (k *keyListener) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {

}

func (k *keyListener) KeyRepeat(key glfw.Key, scancode int, mods glfw.ModifierKey) {

}

