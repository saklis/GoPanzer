package Managers

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ResourceManager struct {
	Images Images
}

type Images struct {
	Tank      rl.Texture2D
	Gun_a_01  rl.Texture2D
	Hull_a_01 rl.Texture2D
	Track_01  rl.Texture2D
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

func (rm *ResourceManager) LoadAll() {
	rm.Images.LoadImages()
}

func (rm *ResourceManager) UnloadAll() {
	rm.Images.UnloadImages()
}

func (i *Images) LoadImages() {
	i.Tank = rl.LoadTexture("assets/Textures/tank.png")
	i.Gun_a_01 = rl.LoadTexture("assets/Textures/gun_a_01.png")
	i.Hull_a_01 = rl.LoadTexture("assets/Textures/hull_a_01.png")
	i.Track_01 = rl.LoadTexture("assets/Textures/track_01.png")
}

func (i *Images) UnloadImages() {
	rl.UnloadTexture(i.Tank)
	rl.UnloadTexture(i.Gun_a_01)
	rl.UnloadTexture(i.Hull_a_01)
	rl.UnloadTexture(i.Track_01)
}
