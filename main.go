package main

import (
	"Managers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// GameState is the enum of states of the game
type GameState int

const (
	GameState_InputSelect GameState = iota
	GameState_MainMenu
	GameState_Game
	GameState_Intermission
	GameState_Pause
)

// CurrentGameState is the current state of the game
var CurrentGameState GameState = GameState_Game

// Managers
var Resources *Managers.ResourceManager
var Input *Managers.InputManager
var GameWorld *Managers.GameManager

// RenderTarget is the render target used for scaling
var RenderTarget rl.RenderTexture2D

func main() {
	// init game window
	rl.InitWindow(1280*1.5, 704*1.5, "GoPanzer")
	defer rl.CloseWindow()

	// init render target
	RenderTarget = rl.LoadRenderTexture(1280, 704)

	// raylib initial config
	rl.SetTargetFPS(60)

	// init managers
	Resources = Managers.NewResourceManager()
	Input = Managers.NewInputManager()
	GameWorld = Managers.NewGameManager(Resources, Input)

	// load all resources
	Resources.LoadAll()

	// init game world
	GameWorld.Init()

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
	rl.UnloadRenderTexture(RenderTarget)
}

// Game drawing
func Draw() {
	// draw to render target
	rl.BeginTextureMode(RenderTarget)
	rl.ClearBackground(rl.Black)

	if CurrentGameState == GameState_Game {
		GameWorld.Draw()
	}

	rl.EndTextureMode()

	// draw render target to screen
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(
		RenderTarget.Texture,
		rl.NewRectangle(0, 0, float32(RenderTarget.Texture.Width), float32(-RenderTarget.Texture.Height)),
		rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())),
		rl.Vector2{},
		0,
		rl.White,
	)

	rl.EndDrawing()
}
