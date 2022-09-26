package assets

import "errors"

type builder struct {
	assets []Asset
}

func createBuilder() Builder {
	out := builder{
		assets: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add assets to the builder
func (app *builder) WithList(assets []Asset) Builder {
	app.assets = assets
	return app
}

// Now builds a new Assets instance
func (app *builder) Now() (Assets, error) {
	if app.assets != nil && len(app.assets) <= 0 {
		app.assets = nil
	}

	if app.assets == nil {
		return nil, errors.New("there must be at least 1 Asset in order to build an Assets instance")
	}

	return createAssets(app.assets), nil
}
