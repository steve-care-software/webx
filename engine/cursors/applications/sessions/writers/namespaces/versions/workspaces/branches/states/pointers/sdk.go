package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
)

// NewApplication creates a new application
func NewApplication(
	dbApp databases.Application,
) Application {
	return createApplication(
		dbApp,
	)
}

// Application represents a pointer application
type Application interface {
	Write(startAtIndex uint64, pointers pointers.Pointers) error
}
