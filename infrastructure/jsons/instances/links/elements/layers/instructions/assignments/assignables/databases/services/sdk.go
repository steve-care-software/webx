package services

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/services"

// NewAdapter creates a new adapter
func NewAdapter() services.Adapter {
	builder := services.NewBuilder()
	return createAdapter(
		builder,
	)
}
