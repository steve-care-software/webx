package files

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

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
