package commits

import (
	"time"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

type content struct {
	description string
	actions     actions.Actions
	createdOn   time.Time
}

func createContent(
	description string,
	actions actions.Actions,
	createdOn time.Time,
) Content {
	out := content{
		description: description,
		actions:     actions,
		createdOn:   createdOn,
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

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
