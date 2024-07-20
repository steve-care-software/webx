package files

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/contexts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers"
)

// NewContextRepository creates a new context repository
func NewContextRepository() contexts.Repository {
	return createContextRepository()
}

// NewContextService creates a new context service
func NewContextService() contexts.Service {
	return createContextService()
}

// NewLayerRepositoryBuilder creates a new layer repository builder
func NewLayerRepositoryBuilder(
	//pointerRepositoryBuilder pointers.RepositoryBuilder,
	adapter layers.Adapter,
) layers.RepositoryBuilder {
	return createLayerRepositoryBuilder(
		//pointerRepositoryBuilder,
		adapter,
	)
}
