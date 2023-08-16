package Managers

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ResourceManager is a struct that holds all the resources
type ResourceManager struct {
	Images Images
}

// Images is a struct that holds all the images. For internal use (has no factory)
type Images struct {
	Tank      rl.Texture2D
	Gun_a_01  rl.Texture2D
	Hull_a_01 rl.Texture2D
	Track_01  rl.Texture2D
}

// Creates a new resource manager. Loads all defined resources
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// Loads all resources
func (rm *ResourceManager) LoadAll() {
	rm.Images.LoadImages()
}

// Unloads all the resources
func (rm *ResourceManager) UnloadAll() {
	rm.Images.UnloadImages()
}

// Loads all the images
func (i *Images) LoadImages() {
	i.Tank = rl.LoadTexture("assets/Textures/tank.png")
	i.Gun_a_01 = rl.LoadTexture("assets/Textures/gun_a_01.png")
	i.Hull_a_01 = rl.LoadTexture("assets/Textures/hull_a_01.png")
	i.Track_01 = rl.LoadTexture("assets/Textures/track_01.png")
}

// Unloads all the images
func (i *Images) UnloadImages() {
	rl.UnloadTexture(i.Tank)
	rl.UnloadTexture(i.Gun_a_01)
	rl.UnloadTexture(i.Hull_a_01)
	rl.UnloadTexture(i.Track_01)
}
