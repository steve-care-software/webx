package criterias

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents the criteria application
type Application interface {
	Execute(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
}
