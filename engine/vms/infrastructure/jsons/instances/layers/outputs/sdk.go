package outputs

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs"
	json_kinds "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/outputs/kinds"
)

// NewAdapter creates a new adapter
func NewAdapter() outputs.Adapter {
	kindAdapter := json_kinds.NewAdapter()
	builder := outputs.NewBuilder()
	return createAdapter(
		kindAdapter.(*json_kinds.Adapter),
		builder,
	)
}
