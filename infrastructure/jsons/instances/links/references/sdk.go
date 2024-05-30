package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/references"
)

// NewAdapter creates a new adapter
func NewAdapter() references.Adapter {
	hashAdapter := hash.NewAdapter()
	referenceBuilder := references.NewReferenceBuilder()
	builder := references.NewBuilder()
	return createAdapter(
		hashAdapter,
		referenceBuilder,
		builder,
	)
}
