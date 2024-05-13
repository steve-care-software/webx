package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents an action builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithInsert(insert string) Builder
	IsDelete() Builder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Path() []string
	Content() Content
}

// Content represents an action content
type Content interface {
	IsDelete() bool
	IsInsert() bool
	Insert() string
}
