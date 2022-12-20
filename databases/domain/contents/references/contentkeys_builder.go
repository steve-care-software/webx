package references

import "errors"

type contentContentKeysBuilder struct {
	list []ContentKey
}

func createContentKeysBuilder() ContentKeysBuilder {
	out := contentContentKeysBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentContentKeysBuilder) Create() ContentKeysBuilder {
	return createContentKeysBuilder()
}

// WithList add contentContentKeys to the builder
func (app *contentContentKeysBuilder) WithList(list []ContentKey) ContentKeysBuilder {
	app.list = list
	return app
}

// Now builds ContentKeys instance
func (app *contentContentKeysBuilder) Now() (ContentKeys, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ContentKey in order to build a ContentKeys instance")
	}

	mp := map[string]ContentKey{}
	for _, oneContentKey := range app.list {
		contentContentKeyname := oneContentKey.Hash().String()
		mp[contentContentKeyname] = oneContentKey
	}

	return createContentKeys(mp, app.list), nil
}
