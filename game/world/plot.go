package world

import b "chicky-chicky-go/game/blocks"

// Plot contains a three-dimensional array of blocks
type Plot struct {
	// blocks is two blocks deep. there can be one
	// block in the back and one in the front. blocks not
	// supported by a block behind it will fall.
	blocks [128][64][2]b.Block // height, width, depth
}

// IsOnScreen returns true if Plot p is visible on the screen
// TODO
func (p Plot) IsOnScreen() bool {
	return false;
}
