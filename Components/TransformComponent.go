package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// type assertion
var _ IComponent = (*TransformComponent)(nil)

// Component that stores position and rotation of entity.
type TransformComponent struct {
	// Position in the game world.
	Position rl.Vector2

	// Rotation in degrees.
	Rotation float32

	// Scale of entity.
	Scale float32

	// Is entity moving?
	IsMoving bool

	// Movement direction.
	MovementDirection rl.Vector2

	// Owner of this component.
	Owner *Entity
}

// Destroy implements Component.
func (tc *TransformComponent) Destroy() {
	tc.Owner = nil
}

// Draw implements Component.
func (tc *TransformComponent) Draw() {
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
func NewTransformComponent(pos rl.Vector2, rot float32, scale float32) *TransformComponent {
	var t TransformComponent = TransformComponent{}

	t.Position = pos
	t.Rotation = rot
	t.Scale = scale
	t.IsMoving = false
	t.MovementDirection = rl.Vector2{X: 0, Y: 0}

	return &t
}
