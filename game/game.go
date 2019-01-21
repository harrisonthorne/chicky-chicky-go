package game

import (
	// "github.com/go-gl/gl/v4.1-core/gl"
	// "github.com/go-gl/glfw/v3.2/gl"
	// "github.com/go-gl/mathgl/mgl32"
	"chicky-chicky-go/characters"
	"chicky-chicky-go/input"
)

// InputManager is the game's main input manager for
// keyboard and mouse handling
var InputManager input.Manager

// TheChicken is THE CHICKEN. THE ONE AND ONLY. THE MASTER
// AND THE EPITOME OF THIS GREAT GAME
var TheChicken characters.Chicken

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
