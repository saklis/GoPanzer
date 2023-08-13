package main

import (
	"Components"
	"Managers"

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

var Game *Managers.GameManager

func main() {
	// init game window
	rl.InitWindow(1920, 1080, "GoPanzer")
	defer rl.CloseWindow()

	// raylib initial config
	rl.SetTargetFPS(60)

	// init game
	Game = Managers.NewGameManager()

	//test - make a tank
	var image rl.Texture2D = rl.LoadTexture("assets/Textures/tank.png")

	var tank *Components.Entity = Components.NewEntity("Tank")
	tank.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 1920/2 - 128, Y: 1080/2 - 128}, 0, 1))
	tank.AddComponent(Components.NewTankSpriteComponent(&image))
	Game.Spawn(tank)

	// main game loop
	for !rl.WindowShouldClose() {
		// input

		Update()

		// physics update

		Draw()
	}

	rl.UnloadTexture(image)
}

// Game logic update
func Update() {
	if CurrentGameState == GameState_Game {
		Game.Update(rl.GetFrameTime())
	}
}

// Game drawing
func Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	if CurrentGameState == GameState_Game {
		Game.Draw()
	}
	rl.EndDrawing()
}
