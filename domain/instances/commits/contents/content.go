package contents

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/previous"
)

type content struct {
	hash     hash.Hash
	action   actions.Action
	previous previous.Previous
}

func createContent(
	hash hash.Hash,
	action actions.Action,
	previous previous.Previous,
) Content {
	out := content{
		hash:     hash,
		action:   action,
		previous: previous,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Action returns the action
func (obj *content) Action() actions.Action {
	return obj.action
}

// Previous returns the previous
func (obj *content) Previous() previous.Previous {
	return obj.previous
}
