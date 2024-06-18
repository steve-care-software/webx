package files

import "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"

// NewLayerRepositoryBuilder creates a new layer repository builder
func NewLayerRepositoryBuilder(
	adapter layers.Adapter,
) layers.RepositoryBuilder {
	return createLayerRepositoryBuilder(
		adapter,
	)
}
