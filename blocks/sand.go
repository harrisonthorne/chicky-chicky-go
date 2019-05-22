package blocks

// SandBlock is a block of sand, easy to destroy/collect by
// hand
type SandBlock struct {
    Block
    lifespan, health float32
}

// NewSandBlock returns a new SandBlock
func NewSandBlock() *SandBlock {
    return &SandBlock{
        lifespan: 5,
        health: 5,
    }
}
