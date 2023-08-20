package Managers

import (
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// InputManager is the manager that handles all input from the user
// It reacts to raw device input and set actions in an object assigned to device.
type InputManager struct {
	// Flag marking, that controls are locked for specific player.
	ControlsLocked bool

	// Reference to InputAction assigned to specific player.
	Player1 *Structs.InputAction
	Player2 *Structs.InputAction

	// Reference to InputAction assigned to specific device.
	Keyboard *Structs.InputAction
	Gamepad  *Structs.InputAction
	Touch    *Structs.InputAction
}

// NewInputManager creates a new InputManager
// returns: Pointer to created InputManager.
func NewInputManager() *InputManager {
	return &InputManager{
		Keyboard: &Structs.InputAction{},
		Gamepad:  &Structs.InputAction{},
		Touch:    &Structs.InputAction{},
	}
}

// Update updates input manager
func (im *InputManager) Update() {
	if im.Keyboard != nil {
		im.UpdateKeyboard()
	}
}

// UpdateKeyboard updates keyboard input
func (im *InputManager) UpdateKeyboard() {
	im.Keyboard.Up = rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)
	im.Keyboard.Down = rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)
	im.Keyboard.Left = rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)
	im.Keyboard.Right = rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
	im.Keyboard.Fire = rl.IsKeyDown(rl.KeySpace)
	im.Keyboard.Menu = rl.IsKeyDown(rl.KeyEscape)

	im.Keyboard.UpBegin = rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp)
	im.Keyboard.DownBegin = rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown)
	im.Keyboard.LeftBegin = rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft)
	im.Keyboard.RightBegin = rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight)
	im.Keyboard.FireBegin = rl.IsKeyPressed(rl.KeySpace)
	im.Keyboard.MenuBegin = rl.IsKeyPressed(rl.KeyEscape)
}

// DEBUG_AssignKeyboardToPlayer1 assigns keyboard to player 1
func (im *InputManager) DEBUG_AssignKeyboardToPlayer1() {
	im.Player1 = im.Keyboard
}
