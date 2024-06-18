package signs

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_creates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	json_validates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// NewAdapter creates a new adapter
func NewAdapter() signs.Adapter {
	createAdapterIns := json_creates.NewAdapter()
	validateAdapter := json_validates.NewAdapter()
	builder := signs.NewBuilder()
	return createAdapter(
		createAdapterIns.(*json_creates.Adapter),
		validateAdapter.(*json_validates.Adapter),
		builder,
	)
}
