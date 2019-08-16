package game

import (
	"fmt"
	"time"

	"github.com/harrisonthorne/chicky-chicky-go/characters"
	"github.com/harrisonthorne/chicky-chicky-go/blocks"
	"github.com/harrisonthorne/chicky-chicky-go/input"
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"github.com/harrisonthorne/chicky-chicky-go/render"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var characterInControl characters.Character
var block = blocks.NewGrassBlock()

func init() {
	input.AddKeyboardListener(&keyListener{})
	chicken := characters.NewChicken()
	characterInControl = chicken
}

// Logic performs logic for the game. This includes movement, physics,
// clocks, animation, etc
func Logic(delta float32) {
	characterInControl.Logic(delta);
}

var cam = render.NewCamera(maths.Vec3{X:0, Y:0, Z:2}, 70, 800.0/600)
var plot = blocks.NewChunk(0)

var last = time.Now()

// Render renders the game.
func Render() {
	characterInControl.Render(cam)
	// plot.Render(cam)
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

