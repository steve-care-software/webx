package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/applications"
)

// Application represents a tree application
type Application interface {
	New(name string, grammar hash.Hash) error
	applications.Database
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
	Search (by selector)
*/
