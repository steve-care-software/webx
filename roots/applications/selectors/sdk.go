package selectors

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/trees"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// Application represents a selector application
type Application interface {
	Retrieve(context uint, hash hash.Hash) (selectors.Selector, error)
	Scan(context uint, input trees.Tree, output interface{}) (selectors.Selector, error)
	Insert(context uint, selector selectors.Selector) error
	InsertAll(context uint, selectors []selectors.Selector) error
}
