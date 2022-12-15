package applications

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// Application represents a tree application
type Application interface {
	New(name string, grammar hash.Hash) error
	Database
	Software
}

// Software represents the tree software application
type Software interface {
}

// Database represents the tree database application
type Database interface {
}

/*
	Migrate
	Insert
	Update
	Delete
	Retrieve
	Search (by selector/program)
*/
