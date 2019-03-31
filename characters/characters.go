package characters

import (
	// "chicky-chicky-go/game/space"
	"github.com/harrisonthorne/chicky-chicky-go/types"
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"github.com/harrisonthorne/chicky-chicky-go/world"
)

func init() {
}

// Character is a living, breathing object in the game of Chicky
// Chicky Go. They can eat, sleep, run, jump, live and die. hopefully
// they don't die unless they're bad
type Character struct {
    world.PhysicalObject
    types.Renderable
	types.Killable // implements lifespan, health, Hit(), Kill()
}

// ISimpleCharacter is an interface to be implemented on
// SimpleCharacters
type ISimpleCharacter interface {
	Hitbox() maths.AABB
}
