package communications

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications"
	json_signs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	json_votes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// NewAdapter creates a new adapter
func NewAdapter() communications.Adapter {
	signAdapter := json_signs.NewAdapter()
	voteAdapter := json_votes.NewAdapter()
	builder := communications.NewBuilder()
	return createAdapter(
		signAdapter.(*json_signs.Adapter),
		voteAdapter.(*json_votes.Adapter),
		builder,
	)
}
