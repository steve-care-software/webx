package files

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/contexts"
)

// NewContextRepository creates a new context repository
func NewContextRepository() contexts.Repository {
	return createContextRepository()
}

// NewContextService creates a new context service
func NewContextService() contexts.Service {
	return createContextService()
}
