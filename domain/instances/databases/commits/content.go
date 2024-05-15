package commits

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

type content struct {
	description string
	actions     actions.Actions
}

func createContent(
	description string,
	actions actions.Actions,
) Content {
	out := content{
		description: description,
		actions:     actions,
	}

	return &out
}

// Description returns the description
func (obj *content) Description() string {
	return obj.description
}

// Actions returns the actions
func (obj *content) Actions() actions.Actions {
	return obj.actions
}
