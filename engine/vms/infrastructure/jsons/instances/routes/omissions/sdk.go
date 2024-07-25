package omissions

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
)

// NewAdapter creates a new adapter
func NewAdapter() omissions.Adapter {
	elementAdapter := json_elements.NewAdapter()
	builder := omissions.NewBuilder()
	return createAdapter(
		elementAdapter.(*json_elements.Adapter),
		builder,
	)
}
