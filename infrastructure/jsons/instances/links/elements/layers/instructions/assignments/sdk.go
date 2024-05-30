package assignments

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments"
	json_assignables "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables"
)

// NewAdapter creates a new adapter
func NewAdapter() assignments.Adapter {
	assignableAdapter := json_assignables.NewAdapter()
	builder := assignments.NewBuilder()
	return createAdapter(
		assignableAdapter.(*json_assignables.Adapter),
		builder,
	)
}
