package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"
)

// NewAdapter creates a new adapter
func NewAdapter() resources.Adapter {
	builder := resources.NewBuilder()
	return createAdapter(
		builder,
	)
}
