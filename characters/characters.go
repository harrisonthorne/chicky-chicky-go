package characters

import (
	// "chicky-chicky-go/game/space"
	"chicky-chicky-go/types"
)

// Character is a living, breathing object in the game of Chicky
// Chicky Go. They can eat, sleep, run, jump, live and die. hopefully
// they don't die unless they're bad
type Character struct {
    types.PhysicalObject
    types.Renderable
	types.Killable // implements lifespan, health, Hit(), Kill()
}

// SimpleCharacter has an AABB hitbox
type SimpleCharacter struct {
	Character
}

// ISimpleCharacter is an interface to be implemented on
// SimpleCharacters
type ISimpleCharacter interface {
	Hitbox() types.AABB
}
