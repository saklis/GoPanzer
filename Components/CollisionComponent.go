package Components

import (
	"Structs"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp"
)

// Check if CollisionComponent implements IComponent.
var _ IComponent = (*CollisionComponent)(nil)

// CollisionComponent is a component that allows entity to collide with other entities.
type CollisionComponent struct {
	Collider      *cp.Body
	CollisionType cp.CollisionType

	Owner *Entity
}

// NewCollisionComponent creates a new CollisionComponent.
// - collider: pointer to physics body that will be used for collision.
// returns: pointer to new CollisionComponent.
func NewCollisionComponent(collisionType cp.CollisionType, collider *cp.Body) *CollisionComponent {
	var component CollisionComponent = CollisionComponent{
		CollisionType: collisionType,
		Collider:      collider,
	}
	component.Collider.UserData = &component
	return &component
}

func (cc *CollisionComponent) CollidesWith(other *CollisionComponent) {
	// handle collision as a tank
	if cc.CollisionType == Structs.COLLISIONTYPE_TANK {
		var position rl.Vector2 = cc.Owner.Transform.Position
		var otherPosition rl.Vector2 = other.Owner.Transform.Position
		var speed float32 = cc.Owner.GetComponent(reflect.TypeOf((*PlayerControllerComponent)(nil))).(*PlayerControllerComponent).MoveSpeed

		var newPos = rl.Vector2Subtract(position, otherPosition)
		// shorten the vector, making 'bounce off' smaller
		// 1650 was found by trial and error, to "feel" best
		// too big of a value will cause phasing through walls - may require tweeking
		newPos = rl.Vector2{X: newPos.X * speed / 1650, Y: newPos.Y * speed / 1650}

		// move tank back to previous position
		cc.Owner.Transform.Position = rl.Vector2Add(position, newPos)
	}
}

// Destroy implements IComponent.
func (*CollisionComponent) Destroy() {
	// do nothing
}

// Draw implements IComponent.
func (cc *CollisionComponent) Draw() {
	// TEST: draw collider
	//cc.DEBUG_DrawColliders()
}

// Init implements IComponent.
func (*CollisionComponent) Init() {
	// do nothing
}

// SetOwner implements IComponent.
// - owner: pointer to the Entity that owns this component.
func (cc *CollisionComponent) SetOwner(owner *Entity) {
	cc.Owner = owner
}

// Update implements IComponent.
func (cc *CollisionComponent) Update(deltaTime float32) {
	var position rl.Vector2 = cc.Owner.Transform.Position

	cc.Collider.SetPosition(cp.Vector{X: float64(position.X), Y: float64(position.Y)})
}

// DEBUG_DrawColliders draws the colliders of the entity.
func (cc *CollisionComponent) DEBUG_DrawColliders() {
	// draw body position
	var pos cp.Vector = cc.Collider.Position()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), 5, rl.Green)

	// draw shape verts
	cc.Collider.EachShape(func(shape *cp.Shape) {
		var poly *cp.PolyShape = shape.Class.(*cp.PolyShape)
		var polyCount int = poly.Count()
		for i := 0; i < polyCount; i++ {
			var start cp.Vector = poly.Vert(i)
			var end cp.Vector = poly.Vert((i + 1) % polyCount) // Loop back to the first vertex for the last line

			rl.DrawLineEx(
				rl.Vector2{X: float32(pos.X + start.X), Y: float32(pos.Y + start.Y)},
				rl.Vector2{X: float32(pos.X + end.X), Y: float32(pos.Y + end.Y)},
				4.0,
				rl.Red,
			)
		}
	})
}
