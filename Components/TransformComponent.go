package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// type assertion
var _ Component = (*TransformComponent)(nil)

// Component that stores position and rotation of entity.
type TransformComponent struct {
	// Position in the game world.
	Position rl.Vector2

	// Rotation in degrees.
	Rotation float32

	// Scale of entity.
	Scale rl.Vector2

	// Is entity moving?
	IsMoving bool

	// Movement direction.
	MovementDirection rl.Vector2

	// Owner of this component.
	Owner *Entity
}

// Destroy implements Component.
func (*TransformComponent) Destroy() {
	// do nothing
}

// Draw implements Component.
func (*TransformComponent) Draw() {
	// do nothing
}

// Init implements Component.
func (*TransformComponent) Init() {
	// do nothing
}

// SetOwner implements Component.
func (t *TransformComponent) SetOwner(owner *Entity) {
	t.Owner = owner
}

// Update implements Component.
func (*TransformComponent) Update(deltaTime float32) {
	// do nothing
}

// Transform component's factory - creates new TransformComponent
// - pos: Position in the game world.
// - rot: Rotation in degrees.
// - scale: Scale of entity.
func NewTransformComponent(pos rl.Vector2, rot float32, scale rl.Vector2) *TransformComponent {
	var t TransformComponent = TransformComponent{}

	t.Position = pos
	t.Rotation = rot
	t.Scale = scale
	t.IsMoving = false
	t.MovementDirection = rl.Vector2{X: 0, Y: 0}

	return &t
}
