package game

import (
	"fmt"

	// "github.com/go-gl/gl/v4.1-core/gl"
	// "github.com/go-gl/mathgl/mgl32"
	"github.com/harrisonthorne/chicky-chicky-go/characters"
	"github.com/harrisonthorne/chicky-chicky-go/input"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var characterInControl characters.Character

func init() {
	println("adding keyboard listener")
	input.AddKeyboardListener(&keyListener{})
	chicken := characters.NewChicken()
	characterInControl = chicken
}

// Logic performs logic for the game. This includes movement, physics,
// clocks, animation, etc
func Logic(delta float32) {

}

// Render renders the game.
func Render() {
	characterInControl.Render()
}

type keyListener struct{}

func (k *keyListener) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key down:", int(key))
}

func (k *keyListener) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key up:", int(key))
}

func (k *keyListener) KeyRepeat(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key repeat:", int(key))
}

