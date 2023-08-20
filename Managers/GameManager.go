package Managers

import (
	"Components"
	"Structs"
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

	// Physics manager reference
	Physics *PhysicManager
}

// NewGameManager creates a new GameManager.
// - resources: Pointer to ResourceManager.
// - input: Pointer to InputManager.
// returns: Pointer to created GameManager.
func NewGameManager(resources *ResourceManager, input *InputManager) *GameManager {
	var gm GameManager = GameManager{
		Resources: resources,
		Input:     input,
	}

	gm.Physics = NewPhysicManager()
	gm.Map = NewMapManager(resources, gm.Physics)

	return &gm
}

// Init initializes the game world.
func (gm *GameManager) Init() {
	// init physics
	gm.Physics.Init()

	// load map
	gm.Map.LoadMap(gm.Resources.Maps.Level01)

	// spawn all entities from map
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
			newEntity.Init()
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

	// update physics
	gm.Physics.Update(deltaTime)

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
	// draw all entities
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] != nil {
			gm.Gameworld[i].Draw()
		}
	}
}

func (gm *GameManager) Destroy() {
	// destroy all entities
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] != nil {
			gm.Gameworld[i].Destroy()
			gm.Gameworld[i] = nil
		}
	}

	// destroy physics
	gm.Physics.Close()
}

// DEBUG_SpawnTestPlayerEntity spawns a test player entity.
func (gm *GameManager) DEBUG_SpawnTestPlayerEntity() {
	var tank *Components.Entity = Components.NewEntity("Player 1")
	tank.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 200, Y: 200}, 45, 1))
	tank.AddComponent(Components.NewPlayerControllerComponent(1, gm.Input.Player1, 100))
	tank.AddComponent(Components.NewTankSpriteComponent(&gm.Resources.Images.Hull_a_01, &gm.Resources.Images.Gun_a_01, &gm.Resources.Images.Track_01, 23))
	tank.AddComponent(Components.NewCollisionComponent(
		Structs.COLLISIONTYPE_TANK, gm.Physics.AddNewBody(rl.Vector2{X: 200, Y: 200}, rl.Vector2{X: 55, Y: 55}, Structs.COLLISIONTYPE_TANK),
	))
	gm.Spawn(tank)
}
