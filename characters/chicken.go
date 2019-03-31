package characters

import (
    "github.com/harrisonthorne/chicky-chicky-go/types"
    "github.com/harrisonthorne/chicky-chicky-go/textures"

    "os"
)

var chickenTextures map[CharacterAction]*textures.Sprite

func init() {
    chickenTextures[ActionNothing] = loadTexture(ActionNothing, "assets/photos/chicken/stand.png")
    chickenTextures[ActionRun] = loadTexture(ActionRun, "assets/photos/chicken/sprint.png")
    chickenTextures[ActionWalk] = loadTexture(ActionWalk, "assets/photos/chicken/walk.png")
    chickenTextures[ActionSquat] = loadTexture(ActionSquat, "assets/photos/chicken/squat.png")
    chickenTextures[ActionPush] = loadTexture(ActionPush, "assets/photos/chicken/push.png")
    chickenTextures[ActionFall] = loadTexture(ActionFall, "assets/photos/chicken/fall.png")
}

// returns the texture, panics if there's an error loading a
// texture
func loadTexture(action CharacterAction, path string) (texture uint32) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    texture, err = textures.NewTexture(file)
    if err != nil {
        panic(err)
    }

    return
}

// Chicken is the main character of this game. we ain't
// callin it chicky chicky for nothing folks
type Chicken struct {
    Character
    Controllable
    inventory []types.Item
}

// NewChicken creates and initializes a new Chicken
func NewChicken() *Chicken {
    return nil
}

// Move moves the chicken!
func (c *Chicken) Move(direction Direction, super bool)  {
    if super {
        c.action = ActionRun
    } else {
        c.action = ActionWalk
    }
    c.direction = DirectionLeft
}

// Jump jumps the chicken
func (c *Chicken) Jump() {
    // TODO
}

// Squat squats the chicken
func (c *Chicken) Squat() {
    c.Stop()
    c.action = ActionSquat
}

// Stop stops the chicken's movement
func (c *Chicken) Stop() {
    c.action = ActionNothing
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
