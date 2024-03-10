package deletes

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithPath(path string) Builder
	WithIdentifier(identifier string) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Context() string
	Path() string
	Identifier() string
}
