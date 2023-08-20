package Components

// IComponent is an interface for all components.
// All components must implement this interface.
type IComponent interface {
	Init()
	Update(deltaTime float32)
	Draw()
	Destroy()
	SetOwner(*Entity)
}
