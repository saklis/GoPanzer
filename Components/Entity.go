package Components

import (
	"fmt"
	"reflect"
	"strconv"
)

// Global counter for created entities.
var GlobalEntityId int = 1

// Increase the counter and return it's current value.
func GetNextEntityId() int {
	GlobalEntityId += 1
	return GlobalEntityId
}

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

// Entity's factory - creates new Entities
// - name: Human-readable name for this entity. It'll have unique id appended to it.
func NewEntity(name string) *Entity {
	var e Entity = Entity{}

	e.Id = GetNextEntityId()
	e.Name = name + "_" + strconv.Itoa(e.Id)
	e.Initialized = false
	e.Delete = false

	return &e
}

// Add new component to entity.
// - newComponent: New component to add to this entity.
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

// Get component of particular class
// - searchedType: Type of component that needs to be found
// returns: Pointer to component. Returns 'nil' if component is not found.
func (e *Entity) GetComponent(searchedType reflect.Type) IComponent {
	for i := 0; i < len(e.Components); i++ {
		if reflect.TypeOf(e.Components[i]) == searchedType {
			return e.Components[i]
		}
	}

	return nil
}

// Initializes Entity and its Components. Called by Spawn() method in Game class
func (e *Entity) Init() {
	for i := 0; i < len(e.Components); i++ {
		e.Components[i].Init()
	}

	e.Initialized = true
}

// Update all attached component
// - deltaTime: Time in seconds since last frame
func (e *Entity) Update(deltaTime float32) {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Update(deltaTime)
		}
	}
}

// Draw all attached component
func (e *Entity) Draw() {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Draw()
		}
	}
}

// Destroy all attached component
func (e *Entity) Destroy() {
	for i := 0; i < len(e.Components); i++ {
		if e.Components[i] != nil {
			e.Components[i].Destroy()
			e.Components[i] = nil
		}
	}
}
