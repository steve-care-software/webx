package compilers

import (
	"github.com/steve-care-software/logics/domain/bytes/criterias"
	"github.com/steve-care-software/logics/domain/bytes/grammars"
)

// Compiler represents a compiler
type Compiler interface {
    Elements() []Element
}

// Element represents an element
type Element interface {
	Grammar() grammars.Grammar
    Composition() Composition
}

// Composition represents a composition
type Composition interface {
    Pattern() string
    Replacements() Replacements
}

// Replacements represents replacements
type Replacements interface {
	List() []Replacement
}

// Replacement represents a replacement
type Replacement interface {
	Name() string
	Criteria() criterias.Criteria
}
