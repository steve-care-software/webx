package decrypts

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"

// NewAdapter creates a new adapter
func NewAdapter() decrypts.Adapter {
	builder := decrypts.NewBuilder()
	return createAdapter(
		builder,
	)
}
