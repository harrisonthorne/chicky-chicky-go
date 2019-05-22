package blocks

// StoneBlock is a block of rock
type StoneBlock struct {
    Block
    lifespan, health float32
}

// NewStoneBlock creates a new GrassBlock and returns it
func NewStoneBlock() *StoneBlock {
	return &StoneBlock{
        lifespan: 100,
        health: 100,
    }
}
