package signatures

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	json_signs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_votes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

// NewAdapter creates a new adapter
func NewAdapter() signatures.Adapter {
	signAdapter := json_signs.NewAdapter()
	voteAdapter := json_votes.NewAdapter()
	builder := signatures.NewBuilder()
	return createAdapter(
		signAdapter.(*json_signs.Adapter),
		voteAdapter.(*json_votes.Adapter),
		builder,
	)
}
