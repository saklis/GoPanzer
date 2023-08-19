package Managers

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ResourceManager is a struct that holds all the resources
type ResourceManager struct {
	// Images holds all the images
	Images Images
	// Maps holds all the maps
	Maps Maps
	// TileSets holds all the tilesets
	TileSets TileSets
}

// Images is a struct that holds all the images. For internal use (has no factory)
type Images struct {
	Tank      rl.Texture2D
	Gun_a_01  rl.Texture2D
	Hull_a_01 rl.Texture2D
	Track_01  rl.Texture2D
}

// Maps is a struct that holds all the maps. For internal use (has no factory)
type Maps struct {
	Level01 string
}

// TileSets is a struct that holds all the tilesets. For internal use (has no factory)
type TileSets struct {
	Assets rl.Texture2D
}

// NewResourceManager creates a new ResourceManager
// returns: Pointer to created ResourceManager.
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// LoadAll loads all the resources
func (rm *ResourceManager) LoadAll() {
	rm.Images.LoadImages()
	rm.Maps.LoadMaps()
	rm.TileSets.LoadTileSets()
}

// UnloadAll unloads all the resources
func (rm *ResourceManager) UnloadAll() {
	rm.Images.UnloadImages()
}

// LoadImages loads all the images
func (i *Images) LoadImages() {
	i.Tank = rl.LoadTexture("Assets/Textures/tank.png")
	i.Gun_a_01 = rl.LoadTexture("Assets/Textures/gun_a_01.png")
	i.Hull_a_01 = rl.LoadTexture("Assets/Textures/hull_a_01.png")
	i.Track_01 = rl.LoadTexture("Assets/Textures/track_01.png")
}

// UnloadImages unloads all the images
func (i *Images) UnloadImages() {
	rl.UnloadTexture(i.Tank)
	rl.UnloadTexture(i.Gun_a_01)
	rl.UnloadTexture(i.Hull_a_01)
	rl.UnloadTexture(i.Track_01)
}

// LoadMaps loads all the maps
func (m *Maps) LoadMaps() {
	m.Level01 = "Assets/Maps/level01.ldtk"
}

// LoadTileSets loads all the tilesets
func (t *TileSets) LoadTileSets() {
	t.Assets = rl.LoadTexture("Assets/Tiles/assets.png")
}

// UnloadTileSets unloads all the tilesets
func (t *TileSets) UnloadTileSets() {
	rl.UnloadTexture(t.Assets)
}
