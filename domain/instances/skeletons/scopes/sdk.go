package scopes

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewScopeBuilder creates a new scope builder
func NewScopeBuilder() ScopeBuilder {
	hashAdapter := hash.NewAdapter()
	return createScopeBuilder(hashAdapter)
}

// Builder represents a scopes builder
type Builder interface {
	Create() Builder
	WithList(list []Scope) Builder
	Now() (Scopes, error)
}

// Scopes represent a scopes
type Scopes interface {
	Hash() hash.Hash
	List() []Scope
	Contains(path []string) bool
}

// ScopeBuilder represents a scope builder
type ScopeBuilder interface {
	Create() ScopeBuilder
	WithPrefix(prefix []string) ScopeBuilder
	Now() (Scope, error)
}

// Scope represents a scope
type Scope interface {
	Hash() hash.Hash
	Prefix() []string
	Contains(path []string) bool
}
