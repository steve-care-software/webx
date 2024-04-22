package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
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

// NewTestInstanceAdapter creates a new test instance adapter
func NewTestInstanceAdapter(
	pointerAdapter pointers.Adapter,
) instances.Adapter {
	return createTestInstanceAdapter(
		pointerAdapter,
	)
}
