package Structs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// AnimatedSprite is a struct that wraps a texture and provides a way to draw it as an animation
// It only supports vertical animations and same-size frames with constant frame rate
type AnimatedSprite struct {
	// The texture to use
	Texture *rl.Texture2D

	// The current frame of the animation
	CurrentFrame int32

	// The number of frames in the spritesheet
	FramesCount int32

	// The target FPS for the animation
	TargetFPS int32

	// The width of a frame
	FrameWidth int32

	// The height of a frame
	FrameHeight int32

	// The internal timer
	_timer float32

	// The target frame time
	_targetFrameTime float32

	// Whether the animation is playing or not
	IsPlaying bool
}

// Creates a new AnimatedSprite
// - texture: the texture to use, must be a vertical spritesheet
// - framesCount: the number of frames in the spritesheet
// - targetFPS: the target FPS for the animation
// Returns a pointer to the new AnimatedSprite
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

// Updates the current animation frame
// - deltaTime: the time elapsed since the last frame
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

// Draws the current animation frame
// - position: the position to draw the sprite at
// - rotation: the rotation of the sprite
// - scale: the scale of the sprite
// - tint: the tint of the sprite
func (as *AnimatedSprite) Draw(position rl.Vector2, rotation float32, scale float32, tint rl.Color) {
	source := rl.Rectangle{
		X:      float32(0),
		Y:      float32(as.CurrentFrame * as.FrameHeight),
		Width:  float32(as.FrameWidth),
		Height: float32(as.FrameHeight),
	}

	origin := rl.Vector2{X: 0, Y: 0}

	dest := rl.Rectangle{
		X:      position.X,
		Y:      position.Y,
		Width:  float32(as.FrameWidth) * scale,
		Height: float32(as.FrameHeight) * scale,
	}

	rl.DrawTexturePro(
		*as.Texture,
		source,
		dest,
		origin,
		rotation,
		tint,
	)
}
