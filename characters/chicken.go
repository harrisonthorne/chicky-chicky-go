package characters

import (
    "github.com/harrisonthorne/chicky-chicky-go/types"
    "github.com/harrisonthorne/chicky-chicky-go/world"
    "github.com/harrisonthorne/chicky-chicky-go/maths"
    "github.com/harrisonthorne/chicky-chicky-go/sprite"
    "github.com/harrisonthorne/chicky-chicky-go/render"
)

var chickenTextures = make(map[CharacterAction]*sprite.Sprite)

// InitGL initializes chicken sprites
func InitGL() {
    chickenTextures[ActionNothing] = sprite.MustNew("assets/photos/chicken/stand.png", 1, 0)
    chickenTextures[ActionRun] = sprite.MustNew("assets/photos/chicken/sprint.png", 4, 0.25)
    chickenTextures[ActionWalk] = sprite.MustNew("assets/photos/chicken/walk.png", 4, 0.5)
    chickenTextures[ActionSquat] = sprite.MustNew("assets/photos/chicken/squat.png", 1, 0)
    chickenTextures[ActionPush] = sprite.MustNew("assets/photos/chicken/push.png", 4, 1)
    chickenTextures[ActionFall] = sprite.MustNew("assets/photos/chicken/fall.png", 2, 0.1)
}

// Chicken is the main character of this game. we ain't
// callin it chicky chicky for nothing folks
type Chicken struct {
	*world.PhysicalObject
	Character

	backpack Backpack
	action CharacterAction
	direction Direction
}

// NewChicken creates and initializes a new Chicken
func NewChicken() *Chicken {
	return &Chicken{action: ActionNothing}
}

// Move moves the chicken!
func (c *Chicken) Move(direction Direction, super bool)  {
    if super {
        c.action = ActionRun
    } else {
        c.action = ActionWalk
    }
    c.direction = direction
}

// Jump jumps the chicken
func (c *Chicken) Jump(super bool) {
	if c.Hitbox != nil {
		c.ApplyForce(maths.Vec2{X: 0, Y: 6})
	}
}

// Down squats the chicken
func (c *Chicken) Down(super bool) {
    c.Stop()
    c.action = ActionSquat
}

// Stop stops the chicken's movement
func (c *Chicken) Stop() {
    c.action = ActionNothing
}

// Hit hits the chicken with the Item and power specified.
func (c *Chicken) Hit(with types.Item, power float32) []types.Item {
    return []types.Item{}
}

// Kill kills the chicken, dropping its inventory
func (c *Chicken) Kill() []types.Item {
    tmp := c.backpack
	c.backpack = make([]types.Item, 1)
    return []types.Item(tmp)
}

// IsAlive returns true if the chicken is alive
func (c *Chicken) IsAlive() bool {
	return true
}

// Animate moves and calculates sprite frames for the
// Chicken
func (c *Chicken) Animate(delta float32) {
	chickenTextures[c.action].Animate(delta)
}

// Render renders the chicken onto the screen
func (c *Chicken) Render(cam *render.Camera) {
	chickenTextures[c.action].Render(cam)
}
