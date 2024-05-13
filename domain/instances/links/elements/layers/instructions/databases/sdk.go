package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases/updates"
)

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithInsert(insert string) Builder
	WithUpdate(update updates.Update) Builder
	WithDelete(delete string) Builder
	WithPurge(purge string) Builder
	Now() (Database, error)
}

// Database represents a database instruction
type Database interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() string
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() string
	IsPurge() bool
	Purge() string
}
