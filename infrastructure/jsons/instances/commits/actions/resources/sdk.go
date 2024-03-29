package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
)

// NewAdapter creates a new adapter
func NewAdapter(
	instanceAdapter instances.Adapter,
) resources.Adapter {
	builder := resources.NewBuilder()
	return createAdapter(
		builder,
		instanceAdapter,
	)
}
