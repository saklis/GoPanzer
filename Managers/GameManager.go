package Managers

import (
	"Components"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// GameworldSize is the maximum number of entities in the game world.
const GameworldSize = 1000

// GameManager is a struct that manages all entities in the game world.
type GameManager struct {
	// All entities in the game world
	Gameworld [GameworldSize]*Components.Entity

	// Input manager reference
	Input *InputManager

	// Resource manager reference
	Resources *ResourceManager

	// Map manager reference
	Map *MapManager
}

// NewGameManager creates a new GameManager.
// - resources: Pointer to ResourceManager.
// - input: Pointer to InputManager.
// returns: Pointer to created GameManager.
func NewGameManager(resources *ResourceManager, input *InputManager) *GameManager {
	return &GameManager{
		Resources: resources,
		Input:     input,
		Map:       NewMapManager(),
	}
}

// Init initializes the game world.
func (gm *GameManager) Init() {
	gm.Map.LoadMap(gm.Resources.Maps.Level01)

	for i := 0; i < len(gm.Map.Entities); i++ {
		if gm.Map.Entities[i] != nil {
			gm.Spawn(gm.Map.Entities[i])
		}
	}
}

// Spawn adds new entity to the game world.
// - newEntity: Pointer to new entity.
func (gm *GameManager) Spawn(newEntity *Components.Entity) {
	// find first empty slot in game world
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] == nil {
			gm.Gameworld[i] = newEntity
			return
		}
	}

	fmt.Println("ERROR: No more space to add extra entity to game world.")
}

// Update updates all entities in game world.
// - deltaTime: Time elapsed since last update.
func (gm *GameManager) Update(deltaTime float32) {
	// slice to hold indexes of entities to delete
	var entityIndexToDelete []int

	// update all entities
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] != nil {
			if gm.Gameworld[i].Delete {
				entityIndexToDelete = append(entityIndexToDelete, i)
			} else {
				gm.Gameworld[i].Update(deltaTime)
			}
		}
	}

	// delete entities
	for i := len(entityIndexToDelete) - 1; i >= 0; i-- {
		gm.Gameworld[entityIndexToDelete[i]].Destroy()
		gm.Gameworld[entityIndexToDelete[i]] = nil
	}
}

// Draw draws all entities in game world.
func (gm *GameManager) Draw() {
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] != nil {
			gm.Gameworld[i].Draw()
		}
	}
}

// DEBUG_SpawnTestPlayerEntity spawns a test player entity.
func (gm *GameManager) DEBUG_SpawnTestPlayerEntity() {
	var tank *Components.Entity = Components.NewEntity("Player 1")
	tank.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 100, Y: 100}, 45, 1))
	tank.AddComponent(Components.NewPlayerControllerComponent(1, gm.Input.Player1, 100))
	tank.AddComponent(Components.NewTankSpriteComponent(&gm.Resources.Images.Hull_a_01, &gm.Resources.Images.Gun_a_01, &gm.Resources.Images.Track_01, 23))
	gm.Spawn(tank)
}
