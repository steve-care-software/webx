package compilers

import (
	"github.com/steve-care-software/syntax/domain/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewCompositionBuilder creates a new composition builder
func NewCompositionBuilder() CompositionBuilder {
	return createCompositionBuilder()
}

// NewReplacementsBuilder creates a new replacements builder
func NewReplacementsBuilder() ReplacementsBuilder {
	return createReplacementsBuilder()
}

// NewReplacementBuilder creates a new replacement builder
func NewReplacementBuilder() ReplacementBuilder {
	return createReplacementBuilder()
}

// Builder represents a compiler builder
type Builder interface {
	Create() Builder
	WithElements(elements []Element) Builder
	Now() (Compiler, error)
}

// Compiler represents a compiler
type Compiler interface {
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGrammar(grammar grammars.Grammar) ElementBuilder
	WithComposition(composition Composition) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Grammar() grammars.Grammar
	Composition() Composition
}

// CompositionBuilder represents a composition builder
type CompositionBuilder interface {
	Create() CompositionBuilder
	WithPrefix(prefix []byte) CompositionBuilder
	WithSuffix(suffix []byte) CompositionBuilder
	WithPattern(pattern []byte) CompositionBuilder
	WithReplacements(replacements Replacements) CompositionBuilder
	Now() (Composition, error)
}

// Composition represents a composition
type Composition interface {
	Prefix() []byte
	Suffix() []byte
	Pattern() []byte
	Replacements() Replacements
}

// ReplacementsBuilder represents a replacements builder
type ReplacementsBuilder interface {
	Create() ReplacementsBuilder
	WithList(list []Replacement) ReplacementsBuilder
	Now() (Replacements, error)
}

// Replacements represents replacements
type Replacements interface {
	List() []Replacement
}

// ReplacementBuilder represents a replacement builder
type ReplacementBuilder interface {
	Create() ReplacementBuilder
	WithName(name []byte) ReplacementBuilder
	WithCriteria(criteria criterias.Criteria) ReplacementBuilder
	Now() (Replacement, error)
}

// Replacement represents a replacement
type Replacement interface {
	Name() []byte
	Criteria() criterias.Criteria
}
