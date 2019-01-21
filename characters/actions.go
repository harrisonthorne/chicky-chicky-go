package characters

// CharacterAction specifies what a certain character is
// doing
type CharacterAction int8

// CharacterDirection is Right or Left, telling which direction a
// character is facing
type CharacterDirection int8

// Definitions for CharacterAction
const (
	ActionNothing CharacterAction = iota
	ActionWalk
	ActionRun
	ActionSquat
	ActionClimb
	ActionFall
	ActionAttack
	ActionHurt
	ActionDying
	ActionPush
	ActionSleep
	ActionEat
)

// Definitions for CharacterDirection
const (
	DirectionRight CharacterAction = iota
	DirectionLeft
)
