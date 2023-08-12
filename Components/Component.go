package Components

// Base interface of Components
// Represents component that can be added to Entity
type Component interface {
	Init()
	Update(deltaTime float32)
	Draw()
	Destroy()
	SetOwner(*Entity)
}
