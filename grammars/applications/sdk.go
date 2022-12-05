package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/coverages"
	"github.com/steve-care-software/webx/grammars/domain/trees"
)

// Application represents a grammar application
type Application interface {
	Open(name string, height int) (*uint, error)
	Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error)
	Search(context uint, suites grammars.Suites) (grammars.Grammar, error)
	Insert(context uint, grammar grammars.Grammar) error
	InsertAll(context uint, grammars []grammars.Grammar) error
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Coverages(grammar grammars.Grammar) (coverages.Coverages, error)
	Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error)
	Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error)
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint) error
	Close(context uint) error
}
