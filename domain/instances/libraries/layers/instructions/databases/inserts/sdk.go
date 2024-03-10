package inserts

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithInstance(instance string) Builder
	WithPath(path string) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Context() string
	Instance() string
	Path() string
}
