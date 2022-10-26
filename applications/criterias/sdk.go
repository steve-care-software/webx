package criterias

import (
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/trees"
	"github.com/steve-care-software/webx/domain/trees/selections"
)

// Application represents the criteria application
type Application interface {
	Retrieve(criteria criterias.Criteria, tree trees.Tree) (selections.Selections, error)
	Execute(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
}
