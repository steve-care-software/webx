package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
)

type action struct {
	hash          hash.Hash
	path          []string
	modifications modifications.Modifications
}

func createAction(
	hash hash.Hash,
	path []string,
	modifications modifications.Modifications,
) Action {
	out := action{
		hash:          hash,
		path:          path,
		modifications: modifications,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *action) Path() []string {
	return obj.path
}

// Modifications returns the modifications
func (obj *action) Modifications() modifications.Modifications {
	return obj.modifications
}
