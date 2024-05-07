package databases

import "github.com/steve-care-software/datastencil/domain/hash"

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Path() []string
	Description() string
	Head() string
	IsActive() bool
}
