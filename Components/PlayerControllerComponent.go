package Components

import (
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Check if PlayerControllerComponent implements IComponent.
var _ IComponent = (*PlayerControllerComponent)(nil)

// PlayerControllerComponent is a component that allows player to control the entity.
type PlayerControllerComponent struct {
	// PlayerNumber is the number of the player that controls this entity.
	PlayerNumber int32

	// MoveSpeed is the speed at which the entity moves.
	MoveSpeed float32

	// PlayerInput is the input action that controls this entity.
	PlayerInput *Structs.InputAction

	// Owner is the entity that owns this component.
	Owner *Entity
}

// Destroy is called when component is destroyed.
func (*PlayerControllerComponent) Destroy() {
	// do nothing
}

// Draw is called when component should be drawn.
func (*PlayerControllerComponent) Draw() {
	// do nothing
}

// Init is called when component is initialized.
func (*PlayerControllerComponent) Init() {
	// do nothing
}

// SetOwner sets the owner of this component.
// - owner: Entity that owns this component.
func (pcc *PlayerControllerComponent) SetOwner(owner *Entity) {
	pcc.Owner = owner
}

// Update is called when component should be updated.
// - deltaTime: Time elapsed since last update.
func (pcc *PlayerControllerComponent) Update(deltaTime float32) {
	// movement vector for this update
	var movement rl.Vector2 = rl.NewVector2(0, 0)

	// read action states
	if pcc.PlayerInput.Up {
		movement.Y -= 1
	}

	if pcc.PlayerInput.Down {
		movement.Y += 1
	}

	if pcc.PlayerInput.Left {
		movement.X -= 1
	}

	if pcc.PlayerInput.Right {
		movement.X += 1
	}

	// if there is movement, normalize and apply speed
	if rl.Vector2Length(movement) > 0 {
		// normalize movement vector
		movement = rl.Vector2Normalize(movement)

		// set movement direction
		var transform *TransformComponent = pcc.Owner.Transform
		transform.IsMoving = true
		transform.MovementDirection = movement

		// apply movement to Entity's transform
		transform.Position.X += movement.X * pcc.MoveSpeed * deltaTime
		transform.Position.Y += movement.Y * pcc.MoveSpeed * deltaTime
		transform.Rotation = rl.Vector2Angle(rl.Vector2Zero(), movement) * rl.Rad2deg
	} else {
		pcc.Owner.Transform.IsMoving = false
	}
}

// NewPlayerControllerComponent creates a new PlayerControllerComponent.
// - playerNumber: Number of the player that controls this entity.
// - action: Input action that controls this entity.
// - moveSpeed: Speed at which the entity moves.
// returns: Pointer to new PlayerControllerComponent.
func NewPlayerControllerComponent(playerNumber int32, action *Structs.InputAction, moveSpeed float32) *PlayerControllerComponent {
	pcc := PlayerControllerComponent{
		PlayerNumber: playerNumber,
		MoveSpeed:    moveSpeed,
		PlayerInput:  action,
	}

	return &pcc
}
