package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
	"github.com/steve-care-software/webx/roots/domain/grammars/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// Application represents a selector application
type Application interface {
	Database
	Software
}

// Software represents the selector software application
type Software interface {
	Matches(grammar grammars.Grammar, selector selectors.Selector) (bool, error)
	Execute(selector selectors.Selector, script []byte) (interface{}, bool, []byte, error)
}

// Database represents the selector database application
type Database interface {
	Retrieve(context uint, hash hash.Hash) (selectors.Selector, error)
	Scan(context uint, input trees.Tree, output interface{}) (selectors.Selector, error)
	Insert(context uint, selector selectors.Selector) error
	InsertAll(context uint, selectors []selectors.Selector) error
}
