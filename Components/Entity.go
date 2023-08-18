package Components

import (
	"fmt"
	"reflect"
	"strconv"
)

// Global counter for all entities in the game world.
var GlobalEntityId int = 1

// GetNextEntityId returns next unique entity id.
// returns: Next unique entity id.
func GetNextEntityId() int {
	GlobalEntityId += 1
	return GlobalEntityId
}

// Entity is a struct that represents a game object.
type Entity struct {
	// Entitie's unique identifier.
	Id int

	// Entitie's human-friendly name.
	Name string

	// Components of this entity.
	Components [10]IComponent

	// Reference to TransformComponent.
	// It's set in AddComponent()
	Transform *TransformComponent

	// A flag marking state of the entity in the game world.
	// Set to 'true' during spawning.
	Initialized bool

	// A flag informing that this entity is market for deletion.
	// Will be removed from game world by next global update.
	Delete bool
}

// NewEntity creates new entity with unique id and name.
// - name: Human-friendly name of the entity.
// returns: Pointer to newly created entity.
func NewEntity(name string) *Entity {
	var e Entity = Entity{}

	e.Id = GetNextEntityId()
	e.Name = name + "_" + strconv.Itoa(e.Id)
	e.Initialized = false
	e.Delete = false

	return &e
}

// AddComponent adds new component to the entity.
// - newComponent: Pointer to new component.
func (e *Entity) AddComponent(newComponent IComponent) {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] == nil {
			e.Components[i] = newComponent
			newComponent.SetOwner(e)

			// check if new component is a TransformComponent
			if tc, ok := newComponent.(*TransformComponent); ok {
				e.Transform = tc
			}

			return
		}
	}

	fmt.Println("ERROR: No more space to add extra component to ", e.Name, " entity!")
}

// GetComponent returns component of given type.
// - searchedType: Type of component to search for.
// returns: Pointer to component of given type or nil if not found.
func (e *Entity) GetComponent(searchedType reflect.Type) IComponent {
	for i := 0; i < len(e.Components); i++ {
		if reflect.TypeOf(e.Components[i]) == searchedType {
			return e.Components[i]
		}
	}

	return nil
}

// Init initializes all attached component
func (e *Entity) Init() {
	for i := 0; i < len(e.Components); i++ {
		e.Components[i].Init()
	}

	e.Initialized = true
}

// Update updates all attached component
// - deltaTime: Time elapsed since last update.
func (e *Entity) Update(deltaTime float32) {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Update(deltaTime)
		}
	}
}

// Draw draws all attached component
func (e *Entity) Draw() {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Draw()
		}
	}
}

// Destroy destroys all attached component
func (e *Entity) Destroy() {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Destroy()
			e.Components[i] = nil
		}
	}
}
