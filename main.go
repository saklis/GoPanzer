package main

import (
	"main/Managers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// game state enum
type GameState int

const (
	GameState_InputSelect GameState = iota
	GameState_MainMenu
	GameState_Game
	GameState_Intermission
	GameState_Pause
)

var CurrentGameState GameState = GameState_Game

var Resources *Managers.ResourceManager
var Input *Managers.InputManager
var GameWorld *Managers.GameManager

func main() {
	// init game window
	rl.InitWindow(640, 480, "GoPanzer")
	defer rl.CloseWindow()

	// raylib initial config
	rl.SetTargetFPS(60)

	// init managers
	Resources = Managers.NewResourceManager()
	Input = Managers.NewInputManager()
	GameWorld = Managers.NewGameManager(Resources, Input)

	// load all resources
	Resources.LoadAll()

	//test - make a tank for player 1
	Input.DEBUG_AssignKeyboardToPlayer1()
	GameWorld.DEBUG_SpawnTestPlayerEntity()
	// end test

	// main game loop
	for !rl.WindowShouldClose() {
		// input update
		Input.Update()

		// Game logic update
		if CurrentGameState == GameState_Game {
			GameWorld.Update(rl.GetFrameTime())
		}

		// TODO: physics update

		// rendering
		Draw()
	}

	Resources.UnloadAll()
}

// Game drawing
func Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	if CurrentGameState == GameState_Game {
		GameWorld.Draw()
	}
	rl.EndDrawing()
}
