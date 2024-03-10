package updates

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction updates.Update) (*uint, error)
}
