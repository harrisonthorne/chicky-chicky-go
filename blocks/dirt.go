package blocks

import (
    "github.com/harrisonthorne/chicky-chicky-go/items"
    "github.com/harrisonthorne/chicky-chicky-go/render"
)

// DirtBlock is a cubeeee of dirttttt
type DirtBlock struct {
    Block
    lifespan, health float32 
}

// NewDirtBlock creates and returns a new DirtBlock
func NewDirtBlock() *DirtBlock {
	return &DirtBlock {
        lifespan: 15,
        health: 15,
    }
}

// Hit is called with the DirtBlock is hit
func (b *DirtBlock) Hit(with interface{}, power float32) []items.Item {
    return nil
}

// Kill is called when the DirtBlock is obliterated. Returns
// Dirt.
// items that might be dropped with the Killable dies.
func (b *DirtBlock) Kill() []items.Item {
    b.health = 0
    return nil //[]items.Item{items.Dirt}
}

// IsAlive returns true if the DirtBlock is still intact
func (b *DirtBlock) IsAlive() bool {
    return b.health > 0
}

// HealthLeft returns the number of health points left on
// the DirtBlock
func (b *DirtBlock) HealthLeft() float32 {
    return b.health
}

// Lifespan returns the lifespan of health points on the
// DirtBlock
func (b *DirtBlock) Lifespan() float32 {
    return b.lifespan
}

// Render renders the DirtBlock
func (b *DirtBlock) Render(c *render.Camera) {

}
