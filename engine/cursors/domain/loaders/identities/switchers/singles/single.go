package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/profiles"
)

type single struct {
	profile    profiles.Profile
	key        keys.Key
	namespaces namespaces.Namespaces
}

func createSingle(
	profile profiles.Profile,
	key keys.Key,
) Single {
	return createSingleInternally(profile, key, nil)
}

func createSingleWithNamespaces(
	profile profiles.Profile,
	key keys.Key,
	namespaces namespaces.Namespaces,
) Single {
	return createSingleInternally(profile, key, namespaces)
}

func createSingleInternally(
	profile profiles.Profile,
	key keys.Key,
	namespaces namespaces.Namespaces,
) Single {
	out := single{
		profile:    profile,
		key:        key,
		namespaces: namespaces,
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

// HasNamespaces returns true if there is namespaces, false otherwise
func (obj *single) HasNamespaces() bool {
	return obj.namespaces != nil
}

// Namespaces returns the namespaces, if any
func (obj *single) Namespaces() namespaces.Namespaces {
	return obj.namespaces
}
