package files

import (
	assignable_files "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

// Application represents a file application
type Application interface {
	Execute(frame stacks.Frame, assignment assignable_files.File) (stacks.Assignable, *uint, error)
}
