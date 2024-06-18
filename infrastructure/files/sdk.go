package files

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

// NewLayerRepositoryBuilder creates a new layer repository builder
func NewLayerRepositoryBuilder(
	adapter layers.Adapter,
) layers.RepositoryBuilder {
	return createLayerRepositoryBuilder(
		adapter,
	)
}

// NewLinkRepositoryBuilder creates a new link repository builder
func NewLinkRepositoryBuilder(
	adapter links.Adapter,
) links.RepositoryBuilder {
	return createLinkRepositoryBuilder(
		adapter,
	)
}
