package namespaces

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewNamespaceBuilder creates a new namespace builder
func NewNamespaceBuilder() NamespaceBuilder {
	return createNamespaceBuilder()
}

const namespaceDoesNotExistsErrPattern = "the namespace (name: %s) does not exists"

// Adapter represents a namespace adapter
type Adapter interface {
	InstancesToBytes(ins Namespaces) ([]byte, error)
	BytesToInstances(data []byte) (Namespaces, []byte, error)
	InstanceToBytes(ins Namespace) ([]byte, error)
	BytesToInstance(data []byte) (Namespace, []byte, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithList(list []Namespace) Builder
	Now() (Namespaces, error)
}

// Namespaces represents namespaces
type Namespaces interface {
	List() []Namespace
	Names() []string
	Fetch(name string) (Namespace, error)
	Index(name string) (*uint, error)
}

// NamespaceBuilder represents a namespace builder
type NamespaceBuilder interface {
	Create() NamespaceBuilder
	WithName(name string) NamespaceBuilder
	WithDescription(description string) NamespaceBuilder
	WithIterations(iterations delimiters.Delimiter) NamespaceBuilder
	IsDeleted() NamespaceBuilder
	Now() (Namespace, error)
}

// Namespace represents a namespace
type Namespace interface {
	Name() string
	Description() string
	IsDeleted() bool
	HasIterations() bool
	Iterations() delimiters.Delimiter
}
