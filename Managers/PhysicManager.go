package Managers

import (
	"Components"
	"Structs"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp"
)

type PhysicManager struct {
	// Space is a pointer to the space managed by physics engine.
	Space *cp.Space

	// CollisionHandler is a pointer to the collision handler.
	CollisionHandler *cp.CollisionHandler
}

// NewPhysicManager creates a new PhysicManager.
// returns: pointer to new PhysicManager.
func NewPhysicManager() *PhysicManager {
	return &PhysicManager{}
}

// AddNewBody creates a new physics body.
// - position: position of the body.
// - size: size of the body.
// returns: pointer to new physics body.
func (p *PhysicManager) AddNewBody(position rl.Vector2, size rl.Vector2, collisionType cp.CollisionType) *cp.Body {
	var body *cp.Body
	var shape *cp.Shape

	body = cp.NewKinematicBody()
	body.SetPosition(cp.Vector{X: float64(position.X), Y: float64(position.Y)})

	shape = cp.NewBox(body, float64(size.X), float64(size.Y), 0)
	shape.SetSensor(true)
	shape.SetCollisionType(collisionType)

	p.Space.AddBody(shape.Body())
	p.Space.AddShape(shape)

	return body
}

// Init initializes the physics manager.
func (pm *PhysicManager) Init() {
	// create a new physical world
	pm.Space = cp.NewSpace()
	pm.Space.Iterations = 10

	pm.CollisionHandler = pm.Space.NewCollisionHandler(Structs.COLLISIONTYPE_TANK, Structs.COLLISIONTYPE_WALL)
	pm.CollisionHandler.BeginFunc = BeginCollision
}

func BeginCollision(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
	body_a, body_b := arb.Bodies()
	var collider_a *Components.CollisionComponent = body_a.UserData.(*Components.CollisionComponent)
	var collider_b *Components.CollisionComponent = body_b.UserData.(*Components.CollisionComponent)

	collider_a.CollidesWith(collider_b)

	return true
}

// Close closes the physics manager.
func (p *PhysicManager) Close() {

}

// Update updates the physics manager.
// - deltaTime: time since last update.
func (pm *PhysicManager) Update(deltaTime float32) {
	pm.Space.Step(float64(deltaTime))
}
