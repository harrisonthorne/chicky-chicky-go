package blocks

import (
    "github.com/harrisonthorne/chicky-chicky-go/render"
    "github.com/harrisonthorne/chicky-chicky-go/items"
)

// GrassBlock is a block of dirt with a green topping of
// grass
type GrassBlock struct {
	Block
	lifespan, health float32
    matrix *Mat4

	gridX, gridY, gridZ int
}

// NewGrassBlock creates a new GrassBlock and returns it
func NewGrassBlock() *GrassBlock {
	return &GrassBlock{
		lifespan: 15,
		health:   15,
	}
}

// Render renders the GrassBlock
func (b *GrassBlock) Render(c *render.Camera) {
	renderBlock(c, b.Matrix(), grassTexture.ID())
}

// Hit is called when the Killable is hit. Returns any items that
// the Killable might drop when hit.
func (b *GrassBlock) Hit(with interface{}, power float32) []items.Item {
    return nil
}

// Kill is called when the Killable should be killed.
// Returns grass and dirt.
func (b *GrassBlock) Kill() []items.Item {
    return nil
}

// IsAlive returns true if the GrassBlock is still intact
func (b *GrassBlock) IsAlive() bool {
    return b.health > 0
}

// HealthLeft returns the number of health points left on the
// block
func (b *GrassBlock) HealthLeft() float32 {
    return b.health
}

// Lifespan returns the lifespan of health points on the Killable
func (b *GrassBlock) Lifespan() float32 {
    return b.lifespan
}

// Matrix returns a pointer to the orientation matrix of
// this block
func (b *GrassBlock) Matrix() *Mat4 {
    return b.matrix
}

// SetMatrix sets the matrix of the block
func (b *GrassBlock) SetMatrix(mat Mat4) {
    b.matrix = &mat
}

func (b *GrassBlock) GridPos() (x, y, z int) {
	return b.gridX, b.gridY, b.gridZ
}

func (b *GrassBlock) SetGridPos(x, y, z int) {
	b.gridX, b.gridY, b.gridZ = x, y, z
}
