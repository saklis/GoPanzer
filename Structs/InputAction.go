package Structs

// InputAction is a struct that represents an action that can be taken by the user
type InputAction struct {
	Up         bool
	Down       bool
	Left       bool
	Right      bool
	Fire       bool
	Menu       bool
	UpBegin    bool
	DownBegin  bool
	LeftBegin  bool
	RightBegin bool
	FireBegin  bool
	MenuBegin  bool
}

// NewInputAction creates a new InputAction struct
func NewInputAction() *InputAction {
	return &InputAction{Up: false, Down: false, Left: false, Right: false, Fire: false, Menu: false, UpBegin: false, DownBegin: false, LeftBegin: false, RightBegin: false, FireBegin: false, MenuBegin: false}
}

// Reset resets the InputAction struct
func (ia *InputAction) Reset() {
	ia.Up = false
	ia.Down = false
	ia.Left = false
	ia.Right = false
	ia.Fire = false
	ia.Menu = false
	ia.UpBegin = false
	ia.DownBegin = false
	ia.LeftBegin = false
	ia.RightBegin = false
	ia.FireBegin = false
	ia.MenuBegin = false
}
