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
var Resources *Managers.ResourceManager

func main() {
	// init game window
	rl.InitWindow(1024, 768, "GoPanzer")
	defer rl.CloseWindow()

	// raylib initial config
	rl.SetTargetFPS(60)

	// init game
	Resources = Managers.NewResourceManager()
	Game = Managers.NewGameManager()

	// load resources
	Resources.LoadAll()

	//test - make a tank
	var tank *Components.Entity = Components.NewEntity("Tank")
	tank.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 1024 / 2, Y: 768 / 2}, 0, 1))
	tank.AddComponent(Components.NewTankSpriteComponent(&Resources.Images.Tank))
	Game.Spawn(tank)

	// main game loop
	for !rl.WindowShouldClose() {
		// input

		Update()

		// physics update

		Draw()
	}

	Resources.UnloadAll()
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
