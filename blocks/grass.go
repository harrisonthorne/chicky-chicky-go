package blocks

// GrassBlock is a block of dirt with a green topping of
// grass
type GrassBlock struct {
    lifespan, health float32
}

// NewGrassBlock creates a new GrassBlock and returns it
func NewGrassBlock() *GrassBlock {
	return &GrassBlock {
        lifespan: 15,
        health: 15,
    }
}
