package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Component = (*TankSpriteComponent)(nil)

type TankSpriteComponent struct {

	// Image to draw.
	Image *rl.Texture2D

	// Owner of this component.
	Owner *Entity
}

// Creates a new TankSpriteComponent instance.
// Returns a pointer to the created TankSpriteComponent.
func NewTankSpriteComponent(image *rl.Texture2D) *TankSpriteComponent {
	return &TankSpriteComponent{
		Image: image,
	}
}

// Destroy implements Component.
func (tsc *TankSpriteComponent) Destroy() {
	tsc.Owner = nil
}

// Draw implements Component.
func (tsc *TankSpriteComponent) Draw() {
	var transform *TransformComponent = tsc.Owner.Transform

	// correct position to draw from center
	var position rl.Vector2 = rl.Vector2{
		X: transform.Position.X - float32(tsc.Image.Width/2),
		Y: transform.Position.Y - float32(tsc.Image.Height/2),
	}
	rl.DrawTextureEx(*tsc.Image, position, transform.Rotation, transform.Scale, rl.White)
}

// Init implements Component.
func (*TankSpriteComponent) Init() {
	// do nothing
}

// SetOwner implements Component.
func (tsc *TankSpriteComponent) SetOwner(owner *Entity) {
	tsc.Owner = owner
}

// Update implements Component.
func (*TankSpriteComponent) Update(deltaTime float32) {
	// do nothing
}
