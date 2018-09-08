package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Manager contains KeyboardListeners and MouseListeners. KeyboardListeners and MouseListeners can be added to it, and triggered when the Manager's callback is called.
type Manager struct {
	keyboards []KeyboardListener
	// mice []MouseListener
}

// KeyCallback is a function for the key callback of Manager m. When this callback is called, respective functions of all children KeyboardListeners are called
func (m Manager) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch(action) {
	case glfw.Press:
		for _, kl := range m.keyboards {
			kl.KeyDown(key, scancode, mods)
		}
	case glfw.Release:
		for _, kl := range m.keyboards {
			kl.KeyUp(key, scancode, mods)
		}
	case glfw.Repeat:
		for _, kl := range m.keyboards {
			kl.KeyRepeat(key, scancode, mods)
		}
	}
}

// AddKeyboardListener adds a KeyboardListener to
// the Manager. The new KeyboardListener will be
// triggered when there is a key event.
func (m Manager) AddKeyboardListener(kl KeyboardListener) {
	m.keyboards = append(m.keyboards, kl)
}
