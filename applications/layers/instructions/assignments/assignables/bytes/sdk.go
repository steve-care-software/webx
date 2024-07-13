package bytes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewApplication creates a new application
func NewApplication() Application {
	hashAdapter := hash.NewAdapter()
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		hashAdapter,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable bytes.Bytes) (stacks.Assignable, *uint, error)
}
