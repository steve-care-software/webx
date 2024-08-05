package namespaces

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

type namespace struct {
	name        string
	description string
	isDeleted   bool
	iterations  delimiters.Delimiter
}

func createNamespace(
	name string,
	description string,
	isDeleted bool,
) Namespace {
	return createNamespaceIntrnally(name, description, isDeleted, nil)
}

func createNamespaceWithIterations(
	name string,
	description string,
	isDeleted bool,
	iterations delimiters.Delimiter,
) Namespace {
	return createNamespaceIntrnally(name, description, isDeleted, iterations)
}

func createNamespaceIntrnally(
	name string,
	description string,
	isDeleted bool,
	iterations delimiters.Delimiter,
) Namespace {
	out := namespace{
		name:        name,
		description: description,
		isDeleted:   isDeleted,
		iterations:  iterations,
	}

	return &out
}

// Name returns the name
func (obj *namespace) Name() string {
	return obj.name
}

// Description returns the description
func (obj *namespace) Description() string {
	return obj.description
}

// IsDeleted returns true if deleted, false otherwise
func (obj *namespace) IsDeleted() bool {
	return obj.isDeleted
}

// HasIterations returns true if there is iterations, false otherwise
func (obj *namespace) HasIterations() bool {
	return obj.iterations != nil
}

// Iterations returns the iterations, if any
func (obj *namespace) Iterations() delimiters.Delimiter {
	return obj.iterations
}
