package characters

import (
    "chicky-chicky-go/space"
    "chicky-chicky-go/types"
)

const (
    defaultWalkVelocity = 1
    defaultRunVelocity = 2
    defaultJumpVelocity = 3
)

// Chicken is the main character of this game. we ain't
// callin it chicky chicky for nothing folks
type Chicken struct {
    SimpleCharacter
    ControllableCharacter
    inventory []types.Item
    position space.Point
    velocity space.Velocity
    hitBounds []types.AABB
}

// NewChicken creates and initializes a new Chicken
func NewChicken() *Chicken {
    return nil
}

// WalkLeft walks the chicken to the left
func (c *Chicken) WalkLeft()  {
    c.action = ActionWalk
    c.velocity.Dx = -defaultWalkVelocity
}

// WalkRight walks the chicken to the right
func (c *Chicken) WalkRight() {
    c.ControllableCharacter.action = ActionWalk
    c.velocity.Dx = defaultWalkVelocity
}

// RunLeft runs the chicken to the left
func (c *Chicken) RunLeft() {
    c.action = ActionRun
    c.velocity.Dx = -defaultRunVelocity
}

// RunRight runs the chicken to the right
func (c *Chicken) RunRight() {
    c.action = ActionRun
    c.velocity.Dx = defaultRunVelocity
}

// Jump jumps the chicken
func (c *Chicken) Jump() {
    c.velocity.Dy = defaultJumpVelocity
}

// Squat squats the chicken
func (c *Chicken) Squat() {
    c.Stop()
}

// Stop stops the chicken's movement
func (c *Chicken) Stop() {
    c.velocity.Dx = 0
}

// Hit hits the chicken with the Item and power specified.
func (c *Chicken) Hit(with types.Item, power int) []types.Item {
    return []types.Item{}
}

// Kill kills the chicken, dropping its inventory
func (c *Chicken) Kill() []types.Item {
    tmp := c.inventory
    c.inventory = []types.Item{}
    return tmp
}

// Animate moves and calculates sprite frames for the
// Chicken
func (c *Chicken) Animate(delta float32) {
    
}

// Render renders the chicken onto the screen
func (c *Chicken) Render() {

}
