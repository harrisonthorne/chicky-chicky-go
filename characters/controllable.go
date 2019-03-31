package characters

import (
	"github.com/harrisonthorne/chicky-chicky-go/input"
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"github.com/harrisonthorne/chicky-chicky-go/types"

	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Controllable is a type that can be controlled using the
// mouse or keyboard. It implements both KeyboardListener
// and MouseListener.
type Controllable struct {
	iControllable
	inControl bool
	action    CharacterAction // What is the character doing?
    direction Direction
}

// TakeControl lets the user take control of this
// Controllable.
func (c *Controllable) TakeControl() {
    c.inControl = true
}

type iControllable interface {
	input.KeyboardListener // embed this interface so we don't have to explicitly define KeyRepeat
    Move(direction Direction, super bool)
	Down(super bool)       // Do something downward
	Jump(super bool)       // Do something when the space bar is pressed
	Reset(super bool)      // Nothing is happening anymore

    // Initiates an attack by the Controllable. Returns
    // whatever the Controllable might be holding, the
    // attack power, and where the Controllable was aiming
	Attack() (with *types.Item, power float32, at maths.Vec2)
}

// KeyDown handles a ControllableCharacter action when a key
// is pressed. Implemented from KeyboardListener.
func (c *Controllable) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("Key down: %v\n", key)
    if !c.inControl { return }

	switch key {
	case glfw.KeyA:
		c.Move(DirectionLeft, false)
	case glfw.KeyS:
		c.Down(false)
	case glfw.KeyD:
		c.Move(DirectionRight, false)
	case glfw.KeySpace:
        c.Jump(false)
	}
}

// KeyUp handles a ControllableCharacter action when a key
// is released. Implemented from KeyboardListener.
func (c *Controllable) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("Key up: %v\n", key)
    if !c.inControl { return }

}
