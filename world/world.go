package world

import (
	"github.com/harrisonthorne/chicky-chicky-go/blocks"
	"github.com/harrisonthorne/chicky-chicky-go/maths"
	"github.com/harrisonthorne/chicky-chicky-go/render"

	"math/rand"
)

const MaxWorldSize = 256 // in chunks. 256 yields thousands and thousands of blocks in each direction

// World contains a slice of Plots
type World struct {
	terrain [][][]blocks.Chunk
	seed    int64

	heightGradientVectors      [MaxWorldSize][MaxWorldSize]maths.Vec2
	humidityGradientVectors    [MaxWorldSize][MaxWorldSize]maths.Vec2
	temperatureGradientVectors [MaxWorldSize][MaxWorldSize]maths.Vec2

	caveGradientVectors [MaxWorldSize][MaxWorldSize][MaxWorldSize]maths.Vec3
	oreGradientVectors  [MaxWorldSize][MaxWorldSize][MaxWorldSize]maths.Vec3

	renderDistance int
}

func NewWorld(seed int64) *World {
	w := &World{
		seed: seed,
	}

	w.heightGradientVectors = generate2DGradientMap(seed)
	w.humidityGradientVectors = generate2DGradientMap(seed + 1)
	w.temperatureGradientVectors = generate2DGradientMap(seed + 2)

	w.caveGradientVectors = generate3DGradientMap(seed + 3)
	w.oreGradientVectors = generate3DGradientMap(seed + 4)

	return w
}

func generate2DGradientMap(seed int64) [MaxWorldSize][MaxWorldSize]maths.Vec2 {
	r := rand.New(rand.NewSource(int64(seed)))
	m := [MaxWorldSize][MaxWorldSize]maths.Vec2{}

	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			m[x][y] = maths.Vec2{r.Float32(), r.Float32()}
		}
	}

	return m
}

func generate3DGradientMap(seed int64) [MaxWorldSize][MaxWorldSize][MaxWorldSize]maths.Vec3 {
	r := rand.New(rand.NewSource(int64(seed)))
	m := [MaxWorldSize][MaxWorldSize][MaxWorldSize]maths.Vec3{}

	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			for z := 0; z < len(m[x][y]); z++ {
				m[x][y][z] = maths.Vec3{r.Float32(), r.Float32(), r.Float32()}
			}
		}
	}

	
	return m
}

// Stringify serializes the world into a big string that can be parsed to
// recreate the world.
func (w *World) Stringify() string {
	return ""
}

func (w *World) generateChunk(x, y, z int) {

}

func (w *World) Render(c *render.Camera) {
	eyeChunkX := int(c.Position().X / (blocks.BlockWidth * blocks.ChunkSize))
	eyeChunkY := int(c.Position().Y / (blocks.BlockWidth * blocks.ChunkSize))
	eyeChunkZ := int(c.Position().Z / (blocks.BlockWidth * blocks.ChunkSize))

	for x := eyeChunkX - w.renderDistance; x < eyeChunkX+w.renderDistance; x++ {
		if x < 0 || x >= len(w.terrain) {
			continue
		}
		for y := eyeChunkY - w.renderDistance; y < eyeChunkY+w.renderDistance; y++ {
			if y < 0 || y >= len(w.terrain[x]) {
				continue
			}
			for z := eyeChunkZ - w.renderDistance; z < eyeChunkZ+w.renderDistance; z++ {
				if z < 0 || z >= len(w.terrain[x][y]) {
					continue
				}
				w.terrain[x][y][z].Render(c)
			}
		}
	}
}
