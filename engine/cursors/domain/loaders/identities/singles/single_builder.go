package singles

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/profiles"
)

type singleBuilder struct {
	profile profiles.Profile
	key     keys.Key
}

func createSingleBuilder() SingleBuilder {
	out := singleBuilder{
		profile: nil,
		key:     nil,
	}

	return &out
}

// Create initializes the singleBuilder
func (app *singleBuilder) Create() SingleBuilder {
	return createSingleBuilder()
}

// WithProfile adds a profile to the singleBuilder
func (app *singleBuilder) WithProfile(profile profiles.Profile) SingleBuilder {
	app.profile = profile
	return app
}

// WithKey adds a key to the singleBuilder
func (app *singleBuilder) WithKey(key keys.Key) SingleBuilder {
	app.key = key
	return app
}

// Now builds a new Single instance
func (app *singleBuilder) Now() (Single, error) {
	if app.profile == nil {
		return nil, errors.New("the profile is mandatory in order to build a Profile instance")
	}

	if app.key == nil {
		return nil, errors.New("the key is mandatory in order to build a Profile instance")
	}

	return createSingle(
		app.profile,
		app.key,
	), nil
}
