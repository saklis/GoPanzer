package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Check if TransformComponent implements IComponent.
var _ IComponent = (*TransformComponent)(nil)

// TransformComponent is a component that holds position, rotation and scale of entity.
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

// Destroy is called when component is destroyed.
func (tc *TransformComponent) Destroy() {
	tc.Owner = nil
}

// Draw is called when component should be drawn.
func (tc *TransformComponent) Draw() {
}

// Init is called when component is initialized.
func (*TransformComponent) Init() {
	// do nothing
}

// SetOwner sets the owner of this component.
// - owner: Entity that owns this component.
func (t *TransformComponent) SetOwner(owner *Entity) {
	t.Owner = owner
}

// Update is called when component should be updated.
// - deltaTime: Time elapsed since last update.
func (*TransformComponent) Update(deltaTime float32) {
	// do nothing
}

// Transform component's factory - creates new TransformComponent
// - pos: Position in the game world.
// - rot: Rotation in degrees
// - scale: Scale of entity.
// returns: Pointer to newly created TransformComponent.
func NewTransformComponent(pos rl.Vector2, rot float32, scale float32) *TransformComponent {
	var t TransformComponent = TransformComponent{}

	t.Position = pos
	t.Rotation = rot
	t.Scale = scale
	t.IsMoving = false
	t.MovementDirection = rl.Vector2{X: 0, Y: 0}

	return &t
}
