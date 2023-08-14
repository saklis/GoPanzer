package Components

import (
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Component = (*TankSpriteComponent)(nil)

type TankSpriteComponent struct {
	HullSprite  *rl.Texture2D
	GunSprite   *rl.Texture2D
	TrackSprite Structs.AnimatedSprite

	TrackOffset float32

	// Owner of this component.
	Owner *Entity
}

// Creates a new TankSpriteComponent instance.
// Returns a pointer to the created TankSpriteComponent.
func NewTankSpriteComponent(hull *rl.Texture2D, gun *rl.Texture2D, track *rl.Texture2D, trackOffset float32) *TankSpriteComponent {
	return &TankSpriteComponent{
		HullSprite:  hull,
		GunSprite:   gun,
		TrackSprite: *Structs.NewAnimatedSprite(track, 2, 4),
		TrackOffset: trackOffset,
	}
}

// Destroy implements Component.
func (tsc *TankSpriteComponent) Destroy() {
	tsc.Owner = nil
}

// Draw implements Component.
func (tsc *TankSpriteComponent) Draw() {
	var transform *TransformComponent = tsc.Owner.Transform

	// correct position to draw from center
	var position rl.Vector2 = rl.Vector2{
		X: transform.Position.X - float32(tsc.HullSprite.Width/2),
		Y: transform.Position.Y - float32(tsc.HullSprite.Height/2),
	}

	vOffset := rl.Vector2{X: 0, Y: tsc.TrackOffset}
	vOffset = Structs.RotateByAngle(vOffset, transform.Rotation)

	// draw tracks
	// TODO: uncomment when movement is implemented
	//tsc.TrackSprite.IsPlaying = transform.IsMoving
	tsc.TrackSprite.Draw(rl.Vector2Add(position, vOffset), transform.Rotation, transform.Scale, rl.White)
	tsc.TrackSprite.Draw(rl.Vector2Subtract(position, vOffset), transform.Rotation, transform.Scale, rl.White)

	// draw hull
	rl.DrawTextureEx(*tsc.HullSprite, position, transform.Rotation, transform.Scale, rl.White)
	//draw gun
	rl.DrawTextureEx(*tsc.GunSprite, position, transform.Rotation, transform.Scale, rl.White)
}

// Init implements Component.
func (*TankSpriteComponent) Init() {
	// do nothing
}

// SetOwner implements Component.
func (tsc *TankSpriteComponent) SetOwner(owner *Entity) {
	tsc.Owner = owner
}

// Update implements Component.
func (tsc *TankSpriteComponent) Update(deltaTime float32) {
	tsc.TrackSprite.Update(deltaTime)
}
