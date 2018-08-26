package interfaces

import "chicky-chicky-go/blocks"

// Animatable is an interface that can be added to
// objects that animate. It calls its Animate(delta) method
// during every logical loop to compute whether or not
// the animation should advance to the next frame, or,
// to compute which frame the animation should be on.
type Animatable interface {
	Animate(delta float32)
}

// Burnable can be added to objects that can be burned.
// Its Burn() method will be called when an offending
// object is placed next to it.
type Burnable interface {
	Burn()
}

// Killable is here to serve as an embed in
// blocks or characters. in other words, anything that
// can be killed.
type Killable struct {
	IKillable
	Lifespan int
	HealthLeft float32
}

// IKillable is implemented in aforementioned objects
// that can be killed
type IKillable interface {
	Hit(points int) 
	Kill() []Item
}
