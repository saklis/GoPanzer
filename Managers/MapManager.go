package Managers

import (
	"Components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// MapManager is a struct that manages the map - read map file, create obstackle entities and so on
type MapManager struct {
	Entities [1000]*Components.Entity
}

// NewMapManager creates a new MapManager
// Returns a pointer to the new MapManager
func NewMapManager() *MapManager {
	var mm MapManager = MapManager{}

	return &mm
}

// LoadMap loads a map from a file
// - mapFile: the path to the map file
func (mm *MapManager) LoadMap(mapFile string) {
	mm.AddGroundEntities()
}

// AddGroundEntity adds a ground entity to the map
func (mm *MapManager) AddGroundEntities() {
	var groundEntity *Components.Entity = Components.NewEntity("Ground")
	groundEntity.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 0, Y: 0}, 0, 1))
	groundEntity.AddComponent(Components.NewRectComponent(1280, 704, rl.DarkGreen))
	mm.Entities[0] = groundEntity
}
