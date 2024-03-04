package connections

import "github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	return createConnectionBuilder()
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return createFieldBuilder()
}

// Builder represents the connections builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
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
	Name() string
	From() Field
	To() Field
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithField(field Field) ElementBuilder
	WithResource(resource resources.Resource) ElementBuilder
	Now() (Element, error)
}

// Element represents an elemnt
type Element interface {
	Field() Field
	//Resource() resources.Resource
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
	Name() string
	Path() []string
}
