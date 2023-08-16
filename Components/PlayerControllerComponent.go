package Components

import (
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ IComponent = (*PlayerControllerComponent)(nil)

type PlayerControllerComponent struct {
	PlayerNumber int32
	MoveSpeed    float32

	PlayerInput *Structs.InputAction

	Owner *Entity
}

// Destroy implements IComponent.
func (*PlayerControllerComponent) Destroy() {
	// do nothing
}

// Draw implements IComponent.
func (*PlayerControllerComponent) Draw() {
	// do nothing
}

// Init implements IComponent.
func (*PlayerControllerComponent) Init() {
	// do nothing
}

// SetOwner implements IComponent.
func (pcc *PlayerControllerComponent) SetOwner(owner *Entity) {
	pcc.Owner = owner
}

// Update implements IComponent.
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

func NewPlayerControllerComponent(playerNumber int32, action *Structs.InputAction, moveSpeed float32) *PlayerControllerComponent {
	pcc := PlayerControllerComponent{
		PlayerNumber: playerNumber,
		MoveSpeed:    moveSpeed,
		PlayerInput:  action,
	}

	return &pcc
}
