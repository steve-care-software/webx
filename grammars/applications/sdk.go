package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/coverages"
	"github.com/steve-care-software/webx/grammars/domain/trees"
)

// Application represents a grammar application
type Application interface {
	New(name string) error
	applications.Database
	Database
	Software
}

// Database represents the grammar database application
type Database interface {
	Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error)
	Search(context uint, suites grammars.Suites) (grammars.Grammar, error)
	Scan(context uint, suites grammars.Suites) error
	Insert(context uint, grammar grammars.Grammar) error
	InsertAll(context uint, grammars []grammars.Grammar) error
}

// Software represents the grammar software application
type Software interface {
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Coverages(grammar grammars.Grammar) (coverages.Coverages, error)
	Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error)
	Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error)
}
