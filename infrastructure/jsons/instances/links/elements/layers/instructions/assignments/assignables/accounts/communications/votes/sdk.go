package votes

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"

// NewAdapter creates a new adapter
func NewAdapter() votes.Adapter {
	builder := votes.NewBuilder()
	return createAdapter(
		builder,
	)
}
