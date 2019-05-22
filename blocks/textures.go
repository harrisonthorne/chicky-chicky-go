package blocks

import (
	"github.com/harrisonthorne/chicky-chicky-go/textures"
)

const baseDir = "./assets/photos/blocks/"

var (
    dirtTexture = textures.NewLazyTexture(baseDir+"dirt.png")
    grassTexture = textures.NewLazyTexture(baseDir+"grass.png")
    stoneTexture = textures.NewLazyTexture(baseDir+"stone.png")
    sandTexture = textures.NewLazyTexture(baseDir+"sand.png")
)
