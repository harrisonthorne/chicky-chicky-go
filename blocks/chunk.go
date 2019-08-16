package blocks

import (
	"github.com/harrisonthorne/chicky-chicky-go/render"
	mgl "github.com/go-gl/mathgl/mgl32"
)

const ChunkSize = 64

// Chunk contains a three-dimensional array of blocks
type Chunk struct {
	blocks [ChunkSize][ChunkSize][ChunkSize]*Block
	ordinal int
    matrix *mgl.Mat4
}

// NewChunk returns a new Chunk with the ordinal
func NewChunk(ordinal int) *Chunk {
    p := new(Chunk)
    p.ordinal = ordinal
    for i := 0; i < len(p.blocks); i++ {
        for j := 0; j < len(p.blocks[i]); j++ {
            for k := 0; k < len(p.blocks[i][j]); k++ {
                block := NewGrassBlock()
                x, y, z := float32(i)*BlockWidth, float32(j)*BlockWidth, float32(-k)*BlockWidth
                block.SetMatrix(mgl.Ident4().Mul4(mgl.Translate3D(x, y, z)))
                p.Set(i, j, k, block)
            }
        }
    }
    return p
}

// At returns the block at the array position
func (p *Chunk) At(i, j, k int) *Block {
    return p.blocks[i][j][k]
}

// Set sets the block at the array index
func (p *Chunk) Set(i, j, k int, b Block) {
    p.blocks[i][j][k] = &b
}

// Render renders Chunk p
func (p *Chunk) Render(c *render.Camera) {
    for i := 0; i < len(p.blocks); i++ {
        for j := 0; j < len(p.blocks[i]); j++ {
            for k := 0; k < len(p.blocks[i][j]); k++ {
                block := *(p.blocks[i][j][k])
                if block == nil {
                    continue
                }
                block.Render(c)
            }
        }
    }
}
