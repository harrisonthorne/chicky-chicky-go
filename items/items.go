package items

// Item is anything that can be held in a user's backpack
type Item interface {
	RenderIcon(x, y int)
	Name() string
}

// Tool is an item that can degrade as it is used
type Tool interface {
	Item
	ToolType() ToolType
	Degrade(by float32)
}

// Weapon is an item (more of a Tool) that may have
// additional special features in the future
type Weapon interface {
	Item
	Degrade(by float32)
}

// ToolType is an enum for determining types of tools
// (shovel, axe, other, etc)
type ToolType int

// Declarations for ToolType
const (
	Shovel ToolType = iota
	Axe
	Pick 
	Other
)
