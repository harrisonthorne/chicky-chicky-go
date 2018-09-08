package characters

import (
	"chicky-chicky-go/interfaces"
	"chicky-chicky-go/game/space"
)

// Character is a living, breathing object in the game of Chicky
// Chicky. They can eat, sleep, run, jump, live and die. hopefully
// they don't die unless they're bad
type Character struct {
	interfaces.Killable
	Position space.Point
	Velocity space.Velocity
}

// SimpleCharacter has an AABB hitbox
type SimpleCharacter struct {
	Character
}

// ISimpleCharacter is an interface to be
// implemented on SimpleCharacters 
type ISimpleCharacter interface {
	GetHitbox() AABB
}
