package Managers

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ResourceManager is a struct that holds all the resources
type ResourceManager struct {
	Images Images
	Maps   Maps
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

// NewResourceManager creates a new ResourceManager
// returns: Pointer to created ResourceManager.
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// LoadAll loads all the resources
func (rm *ResourceManager) LoadAll() {
	rm.Images.LoadImages()
	rm.Maps.LoadMaps()
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
