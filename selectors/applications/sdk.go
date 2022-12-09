package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// Application represents a selector application
type Application interface {
	New(name string) error
	Database
	Software
}

// Software represents the selector software application
type Software interface {
	Matches(grammar grammars.Grammar, selector selectors.Selector) (bool, error)
	Execute(selector selectors.Selector, tree trees.Tree) (interface{}, bool, error)
}

// Database represents the selector database application
type Database interface {
	Retrieve(context uint, hash hash.Hash) (selectors.Selector, error)
	Scan(context uint, input trees.Tree, output interface{}) (selectors.Selector, error)
	Insert(context uint, selector selectors.Selector) error
	InsertAll(context uint, selectors []selectors.Selector) error
}
