package Managers

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ResourceManager struct {
	Images Images
}

type Images struct {
	Tank rl.Texture2D
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
}

func (i *Images) UnloadImages() {
	rl.UnloadTexture(i.Tank)
}
