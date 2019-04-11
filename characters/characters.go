package characters

import (
	// "chicky-chicky-go/game/space"
	"github.com/harrisonthorne/chicky-chicky-go/types"
	"github.com/harrisonthorne/chicky-chicky-go/world"
)

func init() {
}

// Character is a living, breathing object in the game of Chicky
// Chicky Go. They can eat, sleep, run, jump, live and die. hopefully
// they don't die unless they're bad. Character embeds
// PhysicalObject, Renderable, and Killable
type Character struct {
    world.PhysicalObject
	types.Killable // implements lifespan, health, Hit(), Kill()
}
