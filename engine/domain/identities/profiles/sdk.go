package profiles

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the profile adapter
type Adapter interface {
	ToBytes(ins Profile) ([]byte, error)
	ToInstance(data []byte) (Profile, error)
}

// Builder represents the profile builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithPackages(packages []string) Builder
	Now() (Profile, error)
}

// Profile represents a profile
type Profile interface {
	Name() string
	Description() string
	HasPackages() bool
	Packages() []string
}
