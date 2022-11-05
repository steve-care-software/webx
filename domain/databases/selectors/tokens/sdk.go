package tokens

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithReverse(reverse entities.Identifier) Builder
	WithElement(element entities.Identifier) Builder
	WithElementIndex(elementIndex uint) Builder
	WithContentIndex(contentIndex uint) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Entity() entities.Entity
	Reverse() entities.Identifier
	Element() Element
	HasContent() bool
	Content() *uint
}

// Element represents an element
type Element interface {
	Element() entities.Identifier
	Index() uint
}
