package actions

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
	json_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/resources"
)

// NewAdapter creates a new adapter
func NewAdapter(
	instanceAdapter instances.Adapter,
) actions.Adapter {
	actionAdapter := NewActionAdapter(instanceAdapter)
	builder := actions.NewBuilder()
	return createAdapter(
		actionAdapter.(*ActionAdapter),
		builder,
	)
}

// NewActionAdapter represents a new action adapter
func NewActionAdapter(
	instanceAdapter instances.Adapter,
) actions.ActionAdapter {
	resourceAdapter := json_resources.NewAdapter(
		instanceAdapter,
	)

	pointerAdapter := json_pointers.NewAdapter()
	builder := actions.NewActionBuilder()
	return createActionAdapter(
		resourceAdapter.(*json_resources.Adapter),
		pointerAdapter.(*json_pointers.Adapter),
		builder,
	)
}
