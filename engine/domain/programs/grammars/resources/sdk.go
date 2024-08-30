package resources

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	return createResourceBuilder()
}

// Builder represents the resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents the resources
type Resources interface {
	List() []Resource
	Fetch(name string) (Resource, error)
}

// ResourceBuilder represents an resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithName(name string) ResourceBuilder
	WithRelativePath(relativePath string) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents an resource
type Resource interface {
	Name() string
	RelativePath() string
}
