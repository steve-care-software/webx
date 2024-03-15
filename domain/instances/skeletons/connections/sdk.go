package connections

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConnectionBuilder(
		hashAdapter,
	)
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	hashAdapter := hash.NewAdapter()
	return createFieldBuilder(
		hashAdapter,
	)
}

// Builder represents the connections builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	Hash() hash.Hash
	List() []Connection
	Fetch(name string) (Connection, error)
	FetchByPaths(from []string, to []string) (Connection, error)
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithName(name string) ConnectionBuilder
	From(from Field) ConnectionBuilder
	To(to Field) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Hash() hash.Hash
	Name() string
	From() Field
	To() Field
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Create() FieldBuilder
	WithName(name string) FieldBuilder
	WithPath(path []string) FieldBuilder
	Now() (Field, error)
}

// Field represents a connection field
type Field interface {
	Hash() hash.Hash
	Name() string
	Path() []string
}
