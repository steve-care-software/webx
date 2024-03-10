package deletes

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction deletes.Delete) (*uint, error)
}
