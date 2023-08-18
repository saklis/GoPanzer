package Components

import (
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Check if TankSpriteComponent implements IComponent.
var _ IComponent = (*TankSpriteComponent)(nil)

// TankSpriteComponent is a component that draws tank sprite.
type TankSpriteComponent struct {
	// HullSprite is the sprite of the tank hull.
	HullSprite *rl.Texture2D

	// GunSprite is the sprite of the tank gun.
	GunSprite *rl.Texture2D

	// TrackSprite is the sprite of the tank tracks.
	TrackSprite Structs.AnimatedSprite

	// TrackOffset is the offset of the tracks from the center of the tank.
	TrackOffset float32

	// Owner of this component.
	Owner *Entity
}

// NewTankSpriteComponent creates a new TankSpriteComponent.
// - hull: Texture of the tank hull.
// - gun: Texture of the tank gun.
// - track: Texture of the tank tracks.
// - trackOffset: Offset of the tracks from the center of the tank.
// returns: Pointer to created TankSpriteComponent.
func NewTankSpriteComponent(hull *rl.Texture2D, gun *rl.Texture2D, track *rl.Texture2D, trackOffset float32) *TankSpriteComponent {
	return &TankSpriteComponent{
		HullSprite:  hull,
		GunSprite:   gun,
		TrackSprite: *Structs.NewAnimatedSprite(track, 2, 4),
		TrackOffset: trackOffset,
	}
}

// Destroy is called when component is destroyed.
func (tsc *TankSpriteComponent) Destroy() {
	tsc.Owner = nil
}

// Draw is called when component should be drawn.
func (tsc *TankSpriteComponent) Draw() {
	var transform *TransformComponent = tsc.Owner.Transform

	// create a vector that will be used to align tracks with the tank
	vOffset := rl.Vector2{X: 0, Y: tsc.TrackOffset}
	vOffset = Structs.RotateByAngle(vOffset, transform.Rotation)

	// draw tracks
	tsc.TrackSprite.IsPlaying = transform.IsMoving
	tsc.TrackSprite.Draw(rl.Vector2Add(transform.Position, vOffset), transform.Rotation, transform.Scale, rl.White)
	tsc.TrackSprite.Draw(rl.Vector2Subtract(transform.Position, vOffset), transform.Rotation, transform.Scale, rl.White)

	// draw hull
	rl.DrawTexturePro(
		*tsc.HullSprite,
		rl.Rectangle{X: 0, Y: 0, Width: float32(tsc.HullSprite.Width), Height: float32(tsc.HullSprite.Height)},
		rl.Rectangle{X: transform.Position.X, Y: transform.Position.Y, Width: float32(tsc.HullSprite.Width) * transform.Scale, Height: float32(tsc.HullSprite.Height) * transform.Scale},
		rl.Vector2{X: float32(tsc.HullSprite.Width) / 2.0 * transform.Scale, Y: float32(tsc.HullSprite.Height) / 2.0 * transform.Scale},
		transform.Rotation,
		rl.White)
	//draw gun
	rl.DrawTexturePro(
		*tsc.GunSprite,
		rl.Rectangle{X: 0, Y: 0, Width: float32(tsc.HullSprite.Width), Height: float32(tsc.HullSprite.Height)},
		rl.Rectangle{X: transform.Position.X, Y: transform.Position.Y, Width: float32(tsc.HullSprite.Width) * transform.Scale, Height: float32(tsc.HullSprite.Height) * transform.Scale},
		rl.Vector2{X: float32(tsc.HullSprite.Width) / 2.0 * transform.Scale, Y: float32(tsc.HullSprite.Height) / 2.0 * transform.Scale},
		transform.Rotation,
		rl.White)
}

// Init is called when component is initialized.
func (*TankSpriteComponent) Init() {
	// do nothing
}

// SetOwner sets the owner of this component.
// - owner: Entity that owns this component.
func (tsc *TankSpriteComponent) SetOwner(owner *Entity) {
	tsc.Owner = owner
}

// Update is called when component should be updated.
// - deltaTime: Time elapsed since last update.
func (tsc *TankSpriteComponent) Update(deltaTime float32) {
	tsc.TrackSprite.Update(deltaTime)
}
