package votes

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, error)
}
