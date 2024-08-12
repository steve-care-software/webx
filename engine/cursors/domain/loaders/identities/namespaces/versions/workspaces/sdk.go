package singles

// Adapter represents a workspace adapter
type Adapter interface {
	ToBytes(ins Workspace) ([]byte, error)
	ToInstance(data []byte) (Workspace, error)
}

// Builder represents a workspace builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	Now() (Workspace, error)
}

// Workspace represents a workspace
type Workspace interface {
	Description() string
}
