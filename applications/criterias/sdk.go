package criterias

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

// Application represents the criteria application
type Application interface {
	Grammar(input []byte) (grammars.Grammar, []byte, error)
	Extract(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
}
