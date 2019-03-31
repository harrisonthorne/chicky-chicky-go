package types

// Animatable is an interface that can be added to objects
// that animate. It calls its Animate(delta) method during
// every logical loop to compute whether or not the
// animation should advance to the next frame, which frame
// the animation should be on, or the animation
// itself.
// type Animatable interface {
// 	Animate(delta float32)
// }

// Ignitable can be added to objects that can be ignited.
// Such objects can burn any objects with the Burnable
// interface attached.
type Ignitable interface {
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
	Hit(with Item, power float32) []Item

	// Called when the Killable should be killed. Returns any
	// items that might be dropped with the Killable dies.
	Kill() []Item

	// Returns true if the Killable is still alive
	IsAlive() bool
}
