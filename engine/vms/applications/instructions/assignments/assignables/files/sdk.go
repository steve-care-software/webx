package files

import (
	assignable_files "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/files"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// Application represents a file application
type Application interface {
	Execute(frame stacks.Frame, assignment assignable_files.File) (stacks.Assignable, *uint, error)
}
