package votes

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	json_creates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	json_validates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
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
