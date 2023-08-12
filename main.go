package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	fmt.Println("Window initialized")

	// main game loop
	for !rl.WindowShouldClose() {
		// input

		// game logic update

		// physics update

		// draw
		rl.BeginDrawing()
		// draw current frame
		rl.EndDrawing()
	}
}
