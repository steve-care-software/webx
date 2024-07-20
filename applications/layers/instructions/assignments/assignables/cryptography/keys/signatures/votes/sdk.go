package votes

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
	validateApp validates.Application,
) Application {
	return createApplication(
		createApp,
		validateApp,
	)
}

// Application represents a vote application
type Application interface {
	Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, *uint, error)
}
