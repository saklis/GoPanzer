package Components

import rl "github.com/gen2brain/raylib-go/raylib"

var _ IComponent = &SpriteComponent{}

// SpriteComponent is a component that contains a sprite.
type SpriteComponent struct {
	// Source is the source rectangle of the sprite.
	Source rl.Rectangle

	// Target is the target size of the sprite.
	TargetSize rl.Vector2

	// Texture is the pointer to the texture
	TileSet *rl.Texture2D

	// Owner is the pointer to the Entity that owns this component.
	Owner *Entity
}

// NewSpriteComponent creates a new SpriteComponent.
// - tileSet: the pointer to the texture
// - source: the source rectangle of the sprite
// - targetSize: the target size of the sprite
// returns: a pointer to the new SpriteComponent
func NewSpriteComponent(tileSet *rl.Texture2D, source rl.Rectangle, targetSize rl.Vector2) *SpriteComponent {
	var sc SpriteComponent = SpriteComponent{
		Source:     source,
		TargetSize: targetSize,
		TileSet:    tileSet,
	}

	return &sc
}

// Destroy implements IComponent.
func (*SpriteComponent) Destroy() {
	// do nothing
}

// Draw implements IComponent.
func (sc *SpriteComponent) Draw() {
	var transform *TransformComponent = sc.Owner.Transform

	var source rl.Rectangle = rl.NewRectangle(sc.Source.X, sc.Source.Y, sc.Source.Width, sc.Source.Height)
	var target rl.Rectangle = rl.NewRectangle(transform.Position.X, transform.Position.Y, sc.TargetSize.X*transform.Scale, sc.TargetSize.Y*transform.Scale)
	var origin rl.Vector2 = rl.NewVector2(sc.TargetSize.X/2, sc.TargetSize.Y/2)

	rl.DrawTexturePro(*sc.TileSet, source, target, origin, transform.Rotation, rl.White)
}

// Init implements IComponent.
func (*SpriteComponent) Init() {
	// do nothing
}

// SetOwner implements IComponent.
func (sc *SpriteComponent) SetOwner(owner *Entity) {
	sc.Owner = owner
}

// Update implements IComponent.
func (*SpriteComponent) Update(deltaTime float32) {
	// do nothing
}
