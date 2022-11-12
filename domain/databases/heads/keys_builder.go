package heads

import "errors"

type keysBuilder struct {
	mp   map[string]Key
	list []Key
}

func createKeysBuilder() KeysBuilder {
	out := keysBuilder{
		mp:   nil,
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *keysBuilder) Create() KeysBuilder {
	return createKeysBuilder()
}

// WithList add keys to the builder
func (app *keysBuilder) WithList(list []Key) KeysBuilder {
	app.list = list
	return app
}

// Now builds Keys instance
func (app *keysBuilder) Now() (Keys, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Key in order to build a Keys instance")
	}

	mp := map[string]Key{}
	for _, oneKey := range app.list {
		keyname := oneKey.Hash().String()
		mp[keyname] = oneKey
	}

	return createKeys(mp, app.list), nil
}
