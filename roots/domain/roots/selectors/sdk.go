package selectors

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
)

// Builder represents a selectors builder
type Builder interface {
	Create() Builder
	WithList(list []Selector) Builder
	Now() (Selectors, error)
}

// Selectors represents selectors
type Selectors interface {
	Hash() hash.Hash
	List() []Selector
}

// SelectorBuilder represents a selector builder
type SelectorBuilder interface {
	Create() SelectorBuilder
	WithName(name string) SelectorBuilder
	WithGrammar(grammar grammars.Grammar) SelectorBuilder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Selector, error)
}

// Selector represents a selector database
type Selector interface {
	Hash() hash.Hash
	Name() string
	Grammar() grammars.Grammar
	HasHistory() bool
	History() hashtrees.HashTree
}
