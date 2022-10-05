package coverages

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages/uncovers"
)

// Application represents the coverage application
type Application interface {
	Coverage(grammar grammars.Grammar) (coverages.Coverage, error)
	Uncovered(coverage coverages.Coverage) (uncovers.Uncovers, error)
}
