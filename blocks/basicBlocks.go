package blocks

// GrassBlock implements Block
type GrassBlock struct {
	Block
}

// NewGrassBlock creates a new GrassBlock and returns it
func NewGrassBlock() GrassBlock {
	return GrassBlock{
		Block: Block{
			Killable: Killable{
				15,
				15,
			},
			spritePath: "grass.png",
		},
	}
}

// DirtBlock is a cubeeee of dirttttt
type DirtBlock struct {
	Block
}

// NewDirtBlock creates and returns a new DirtBlock
func NewDirtBlock() DirtBlock {
	return DirtBlock{
		Block: Block{
			Killable: Killable{
				10,
				10,
			},
			spritePath: "dirt.png",
		},
	}
}
