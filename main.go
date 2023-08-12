package main

import (
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

func main() {
	// init game window
	rl.InitWindow(1920, 1080, "GoPanzer")
	defer rl.CloseWindow()

	// raylib initial config
	rl.SetTargetFPS(60)

	// main game loop
	for !rl.WindowShouldClose() {
		// input

		Update()

		// physics update

		Draw()
	}
}

// Game logic update
func Update() {
}

// Game drawing
func Draw() {
	rl.BeginDrawing()
	// draw current frame
	rl.EndDrawing()
}
