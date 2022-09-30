package grammars

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages/uncovers"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

// Application represents a grammar application
type Application interface {
	Grammar(input []byte) (grammars.Grammar, []byte, error)
	Tree(grammar grammars.Grammar, values []byte) (trees.Tree, []byte, error)
	Test(grammar grammars.Grammar) (coverages.Coverage, error)
	Uncovered(coverage coverages.Coverage) (uncovers.Uncovers, error)
}
