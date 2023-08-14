package Structs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	Texture          *rl.Texture2D
	CurrentFrame     int32
	FramesCount      int32
	TargetFPS        int32
	FrameWidth       int32
	FrameHeight      int32
	_timer           float32
	_targetFrameTime float32
	IsPlaying        bool
}

func NewAnimatedSprite(texture *rl.Texture2D, framesCount int32, targetFPS int32) *AnimatedSprite {
	as := AnimatedSprite{}
	as.Texture = texture
	as.FramesCount = framesCount
	as.TargetFPS = targetFPS

	as.CurrentFrame = 0
	as.FrameWidth = as.Texture.Width
	as.FrameHeight = int32(as.Texture.Height / as.FramesCount)
	as._timer = 0
	as._targetFrameTime = float32(1.0 / as.TargetFPS)
	as.IsPlaying = true

	return &as
}

func (as *AnimatedSprite) Update(deltaTime float32) {
	if as.IsPlaying {
		as._timer += deltaTime

		if as._timer >= as._targetFrameTime {
			as._timer = 0
			as.CurrentFrame++

			if as.CurrentFrame >= as.FramesCount {
				as.CurrentFrame = 0
			}
		}
	}
}

func (as *AnimatedSprite) Draw(position rl.Vector2, rotation float32, scale float32, tint rl.Color) {
	source := rl.Rectangle{
		X:      float32(0),
		Y:      float32(as.CurrentFrame * as.FrameHeight),
		Width:  float32(as.FrameWidth),
		Height: float32(as.FrameHeight),
	}

	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(
		*as.Texture,
		source,
		rl.Rectangle{
			X:      position.X,
			Y:      position.Y,
			Width:  float32(as.FrameWidth) * scale,
			Height: float32(as.FrameHeight) * scale,
		},
		origin,
		rotation,
		tint,
	)
}
