package actions

import "github.com/steve-care-software/datastencil/domain/hash"

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Path() string
	IsDelete() bool
	IsInsert() bool
	Insert() string
}
