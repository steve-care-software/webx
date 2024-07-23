package files

import (
	instruction_files "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/files"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, file instruction_files.File) (*uint, error)
}
