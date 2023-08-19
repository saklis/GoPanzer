package Managers

import (
	"Components"
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// MapManager is a struct that manages the map - read map file, create obstackle entities and so on
type MapManager struct {
	// Entities holds all the entities in this map
	Entities [1000]*Components.Entity

	// Resources is a pointer to the ResourceManager
	Resources *ResourceManager
}

// NewMapManager creates a new MapManager
// - resources: Pointer to ResourceManager.
// Returns a pointer to the new MapManager
func NewMapManager(resources *ResourceManager) *MapManager {
	var mm MapManager = MapManager{
		Resources: resources,
	}

	return &mm
}

// LoadMap loads a map from a file
// - mapFile: the path to the map file
func (mm *MapManager) LoadMap(mapFile string) {
	mm.AddGroundEntities()
	mm.AddObstacleEntities(mapFile)
}

// AddObstacleEntities adds obstacle entities to the map
// - mapFile: the path to the map file
func (mm *MapManager) AddObstacleEntities(mapFile string) {
	// open map file
	file, err := os.Open(mapFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// decode map file
	var data JSONData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	// handle different layers
	for _, level := range data.Levels {
		if level.Identifier == "Level01" { // TODO: make this dynamic
			for _, layerInstance := range level.LayerInstances {

				// handle obstacles
				if layerInstance.Identifier == "Obstacles" {
					// get grid size
					var gridSize float32 = layerInstance.GridSize

					// add obstacle entities
					for i := 0; i < len(layerInstance.GridTiles); i++ {
						var tile *LdtkGridTile = &layerInstance.GridTiles[i]

						var obstacleEntity *Components.Entity = Components.NewEntity("Obstacle")
						obstacleEntity.AddComponent(Components.NewTransformComponent(
							rl.Vector2{X: tile.Px[0] + gridSize/2, Y: tile.Px[1] + gridSize/2},
							0,
							1,
						))
						obstacleEntity.AddComponent(Components.NewSpriteComponent(
							&mm.Resources.TileSets.Assets,
							rl.NewRectangle(tile.Src[0], tile.Src[1], gridSize, gridSize),
							rl.NewVector2(gridSize, gridSize),
						))
						mm.Entities[mm.GetFreeEntityIndex()] = obstacleEntity
					}
				}
			}
		}
	}
}

// AddGroundEntity adds a ground entity to the map
func (mm *MapManager) AddGroundEntities() {
	var groundEntity *Components.Entity = Components.NewEntity("Ground")
	groundEntity.AddComponent(Components.NewTransformComponent(rl.Vector2{X: 0, Y: 0}, 0, 1))
	groundEntity.AddComponent(Components.NewRectComponent(1280, 704, rl.DarkGreen))
	mm.Entities[0] = groundEntity

	// TODO: Add more ground entities with decorations
	// OR make ground editable in LDTK, too
}

func (mm *MapManager) GetFreeEntityIndex() int {
	for i := 0; i < len(mm.Entities); i++ {
		if mm.Entities[i] == nil {
			return i
		}
	}

	panic("No free entity index found!")
}

// LdtkGridTile is a struct that represents a map tile from LDTK
type LdtkGridTile struct {
	Px  [2]float32 `json:"px"`
	Src [2]float32 `json:"src"`
	F   float32    `json:"f"`
	T   float32    `json:"t"`
	D   [1]float32 `json:"d"`
	A   float32    `json:"a"`
}

// LayerInstance is a struct that represents a layer instance from LDTK
type LayerInstance struct {
	Identifier string         `json:"__identifier"`
	GridSize   float32        `json:"__gridSize"`
	GridTiles  []LdtkGridTile `json:"gridTiles"`
}

// Level is a struct that represents a level from LDTK
type Level struct {
	Identifier     string          `json:"identifier"`
	LayerInstances []LayerInstance `json:"layerInstances"`
}

// JSONData is a struct that represents the data from a LDTK map file
type JSONData struct {
	Levels []Level `json:"levels"`
}
