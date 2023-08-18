package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ IComponent = (*RectComponent)(nil)

// RectComponent is a component that draws a rectangle.
type RectComponent struct {
	// Position in the game world.
	Width  int32
	Height int32

	// Color of rectangle.
	Color rl.Color

	// Owner of this component.
	Owner *Entity
}

// NewRectComponent creates a new RectComponent.
// - width: Width of rectangle.
// - height: Height of rectangle.
// - color: Color of rectangle.
// returns: Pointer to created RectComponent.
func NewRectComponent(width int32, height int32, color rl.Color) *RectComponent {
	return &RectComponent{
		Width:  width,
		Height: height,
		Color:  color,
	}
}

// Destroy destroys the component.
func (*RectComponent) Destroy() {
	// do nothing
}

// Draw draws the component.
func (rc *RectComponent) Draw() {
	var position rl.Vector2 = rc.Owner.Transform.Position
	rl.DrawRectangle(int32(position.X), int32(position.Y), rc.Width, rc.Height, rc.Color)
}

// Init initializes the component.
func (*RectComponent) Init() {
	// do nothing
}

// SetOwner sets the owner of this component.
// - owner: Entity that owns this component.
func (rc *RectComponent) SetOwner(owner *Entity) {
	rc.Owner = owner
}

// Update updates the component.
// - deltaTime: Time elapsed since last update.
func (*RectComponent) Update(deltaTime float32) {
	// do nothing
}
