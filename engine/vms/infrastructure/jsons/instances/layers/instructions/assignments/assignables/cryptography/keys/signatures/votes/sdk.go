package votes

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	json_creates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	json_validates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// NewAdapter creates a new adapter
func NewAdapter() votes.Adapter {
	createAdapterIns := json_creates.NewAdapter()
	validateAdapter := json_validates.NewAdapter()
	builder := votes.NewBuilder()
	return createAdapter(
		createAdapterIns.(*json_creates.Adapter),
		validateAdapter.(*json_validates.Adapter),
		builder,
	)
}
