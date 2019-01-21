package characters

import (
	"chicky-chicky-go/input"
	"chicky-chicky-go/types"
	"chicky-chicky-go/space"

    "fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// ControllableCharacter is a character that can controlled
// by the player
type ControllableCharacter struct {
	iControllableCharacter
	Character
	inControl bool
	action    CharacterAction // What is the character doing?
}

// IControllableCharacter is an interface that makes up
// ControllableCharacters
type iControllableCharacter interface {
	input.KeyboardListener // embed this interface so we don't have to explicitly define KeyRepeat
	WalkLeft()             // Make the character walk to the left
	WalkRight()            // Make the character walk to the right
	RunLeft()              // Make the character run to the left
	RunRight()             // Make the character run ro the right
	Jump()                 // Make the character Jump
	Squat()                // Squat the character
	Stop()                 // Stop the character's motion

	// Initiates an attack by the character. Returns whatever
	// the character is holding, the attack power, and where the
	// character was aiming
	Attack() (with types.Item, power float32, at space.Point)
}

// KeyDown handles a ControllableCharacter action when a key
// is pressed. Implemented from KeyboardListener.
func (c *ControllableCharacter) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {
    fmt.Printf("Key down: %v\n", key)

    switch key {
    case glfw.KeyA:
    case glfw.KeyS:
    case glfw.KeyD:
    case glfw.KeySpace:
    }
}

// KeyUp handles a ControllableCharacter action when a key
// is released. Implemented from KeyboardListener.
func (c *ControllableCharacter) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {
    fmt.Printf("Key up: %v\n", key)

}
