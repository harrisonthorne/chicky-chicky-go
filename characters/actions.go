package characters


// CharacterAction specifies what a certain character is
// doing
type CharacterAction int8

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


// Direction is Right or Left, telling which direction a
// character (or whatever) is facing
type Direction int8

// Definitions for CharacterDirection
const (
	DirectionRight Direction = iota
	DirectionLeft
)
