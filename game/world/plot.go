package world

import b "chicky-chicky-go/game/blocks"

// Plot contains a three-dimensional array of blocks
type Plot struct {
	blocks [128][64][2]b.Block // height, width, depth
}

// IsOnScreen returns true if Plot p is visible on the screen
func (p Plot) IsOnScreen() bool {
	return false;
}
