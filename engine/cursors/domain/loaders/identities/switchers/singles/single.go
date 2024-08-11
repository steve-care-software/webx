package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/profiles"
)

type single struct {
	profile profiles.Profile
	key     keys.Key
}

func createSingle(
	profile profiles.Profile,
	key keys.Key,
) Single {
	out := single{
		profile: profile,
		key:     key,
	}

	return &out
}

// Profile returns the profile
func (obj *single) Profile() profiles.Profile {
	return obj.profile
}

// Key returns the key
func (obj *single) Key() keys.Key {
	return obj.key
}
