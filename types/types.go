package types

import (
	"github.com/harrisonthorne/chicky-chicky-go/render"
	"github.com/harrisonthorne/chicky-chicky-go/items"
)

// Animatable is an interface that can be added to objects
// that animate. It calls its Animate(delta) method during
// every logical loop to compute whether or not the
// animation should advance to the next frame, which frame
// the animation should be on, or the animation
// itself.
// type Animatable interface {
// 	Animate(delta float32)
// }

// Flammable can be added to objects that can be ignited.
// Such objects can burn any objects with the Burnable
// interface attached.
type Flammable interface {
    Ignite()
	Ignited() bool
}

// Burnable can be added to objects that can be burned.  Its
// Burn() method will be called when an offending object is
// placed next to it.
type Burnable interface {
	Burn()
}

// Killable is here to serve as an embed in blocks or
// characters. in other words, anything that can be killed.
type Killable interface {
	// Called when the Killable is hit. Returns any items that
	// the Killable might drop when hit.
	Hit(with interface{}, power float32) []items.Item

	// Called when the Killable should be killed. Returns any
	// items that might be dropped with the Killable dies.
	Kill() []items.Item

	// Returns true if the Killable is still alive. A
	// Killable can still be alive even if it has no health
	// left. Any Killables determined to be dead are removed
	// from the world
	IsAlive() bool
// Returns the number of health points left on the
	// Killable
	HealthLeft() float32

	// Returns the lifespan of health points on the Killable
	Lifespan() float32
}

// Renderable is implemented by anything that can be
// rendered.
type Renderable interface {
	Render(c *render.Camera)
}

// Logicable is on objects that should have logic calculated for them
type Logicable interface {
	Logic(delta float32)
}
