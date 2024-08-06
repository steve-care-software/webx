package bytes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
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
