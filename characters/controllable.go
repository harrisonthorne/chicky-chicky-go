package characters

import (
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"github.com/harrisonthorne/chicky-chicky-go/items"
)

// Controllable is a type that can be controlled using the
// mouse or keyboard. It implements both KeyboardListener
// and MouseListener.
type Controllable interface {
    Move(direction Direction, super bool)
	Down(super bool)       // Do something downward (fall or squat maybe?)
	Jump(super bool)       // Do something when the space bar is pressed
	Stop()      // Nothing is happening anymore

    // Initiates an attack by the Controllable. Returns
    // whatever the Controllable might be holding, the
    // attack power, and where the Controllable was aiming
	Attack() (with *items.Item, power float32, at maths.Vec2)
}

// TODO
// // KeyDown handles a Controllable action when a key
// // is pressed. Implemented from KeyboardListener.
// func (c *Controllable) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {
// 	fmt.Printf("Key down: %v\n", key)
//     if !c.inControl { return }

// 	switch key {
// 	case glfw.KeyA:
// 		c.Move(DirectionLeft, false)
// 	case glfw.KeyS:
// 		c.Down(false)
// 	case glfw.KeyD:
// 		c.Move(DirectionRight, false)
// 	case glfw.KeySpace:
//         c.Jump(false)
// 	}
// }

// // KeyUp handles a Controllable action when a key
// // is released. Implemented from KeyboardListener.
// func (c *Controllable) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {
// 	fmt.Printf("Key up: %v\n", key)
//     if !c.inControl { return }

// }
