package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases/updates"
)

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
