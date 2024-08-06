package namespaces

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"

// NewNamespacesForTests creates a new namespaces for tests
func NewNamespacesForTests(list []Namespace) Namespaces {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewNamespaceForTests creates a new namespace for tests
func NewNamespaceForTests(name string, description string, isDeleted bool) Namespace {
	builder := NewNamespaceBuilder().Create().
		WithName(name).
		WithDescription(description)

	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewNamespaceWithIterationsForTests creates a new namespace with iterations
func NewNamespaceWithIterationsForTests(name string, description string, isDeleted bool, iterations delimiters.Delimiter) Namespace {
	builder := NewNamespaceBuilder().Create().
		WithName(name).
		WithDescription(description).
		WithIterations(iterations)

	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
