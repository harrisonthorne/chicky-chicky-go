package blocks

import (
    "github.com/harrisonthorne/chicky-chicky-go/types"
    // "github.com/harrisonthorne/chicky-chicky-go/items"
    "github.com/harrisonthorne/chicky-chicky-go/render"
)

// DirtBlock is a cubeeee of dirttttt
type DirtBlock struct {
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
func (b *DirtBlock) Hit(with interface{}, power float32) []types.Item {

    return nil
}

// Kill is called when the DirtBlock is obliterated. Returns
// Dirt.
// items that might be dropped with the Killable dies.
func (b *DirtBlock) Kill() []types.Item {
    b.health = 0
    return []types.Item{items.Dirt}
}

// IsAlive returns true if the DirtBlock is still intact
func (b *DirtBlock) IsAlive() bool {

}

// HealthLeft returns the number of health points left on
// the DirtBlock
func (b *DirtBlock) HealthLeft() float32 {

}

// Lifespan returns the lifespan of health points on the
// DirtBlock
func (b *DirtBlock) Lifespan() float32 {

}

// Render renders the DirtBlock
func (b *DirtBlock) Render(c *render.Camera) {

}
