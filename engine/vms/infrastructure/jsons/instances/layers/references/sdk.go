package references

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/references"
)

// NewAdapter creates a new adapter
func NewAdapter() references.Adapter {
	referenceBuilder := references.NewReferenceBuilder()
	builder := references.NewBuilder()
	return createAdapter(
		referenceBuilder,
		builder,
	)
}
