package assignments

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments"
	json_assignables "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables"
)

// NewAdapter creates a new adapter
func NewAdapter() assignments.Adapter {
	assignableAdapter := json_assignables.NewAdapter()
	builder := assignments.NewBuilder()
	return createAdapter(
		assignableAdapter.(*json_assignables.Adapter),
		builder,
	)
}
