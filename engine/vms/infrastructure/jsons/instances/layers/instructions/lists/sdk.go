package lists

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists"
	json_deletes "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/lists/deletes"
	json_inserts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/lists/inserts"
)

// NewAdapter creates a new adapter
func NewAdapter() lists.Adapter {
	deleteAdapter := json_deletes.NewAdapter()
	insertAdapter := json_inserts.NewAdapter()
	builder := lists.NewBuilder()
	return createAdapter(
		deleteAdapter.(*json_deletes.Adapter),
		insertAdapter.(*json_inserts.Adapter),
		builder,
	)
}
