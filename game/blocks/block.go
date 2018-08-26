package blocks

import "chicky-chicky-go/interfaces"

// Block block block block block block block block
type Block struct {
	Killable
	spritePath string
}

// Hit decrases the health of Block b by points
func (b *Block) Hit(points float32) {
	b.HealthLeft -= points;
}

// Kill returns a slice of Items when the Block
// is killed
func (b *Block) Kill() []Item {
	return make([]Item, 1)
}	


// Item TODO: build this
type Item struct {
}
