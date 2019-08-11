package world

import (
	"github.com/harrisonthorne/chicky-chicky-go/blocks"
)

// World contains a slice of Plots
type World struct {
	terrain [][][]blocks.Plot
}
