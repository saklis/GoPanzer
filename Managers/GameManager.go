package Managers

import (
	"Components"
	"fmt"
)

const GameworldSize = 1000

// Holds game world (world's entities) and game logic.
type GameManager struct {
	// All entities in the game world
	Gameworld [GameworldSize]*Components.Entity
}

// Creates a new GameManager instance.
// Returns a pointer to the created GameManager.
func NewGameManager() *GameManager {
	return &GameManager{}
}

// Adds a new entity to the game world.
// - newEntity: A pointer to a newEntity of type Components.Entity to add to the world.
func (gm *GameManager) Spawn(newEntity *Components.Entity) {
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] == nil {
			gm.Gameworld[i] = newEntity
			return
		}
	}

	fmt.Println("ERROR: No more space to add extra entity to game world.")
}

// Update all entities in game world
// - deltaTime: Time in seconds since last frame
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

// Draw all entities in game world
func (gm *GameManager) Draw() {
	for i := 0; i < len(gm.Gameworld); i++ {
		if gm.Gameworld[i] != nil {
			gm.Gameworld[i].Draw()
		}
	}
}
