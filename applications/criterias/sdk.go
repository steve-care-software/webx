package criterias

import (
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/trees"
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents the criteria application
type Application interface {
	Execute(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
}
