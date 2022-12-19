package configs

import (
	"errors"
	"fmt"
)

type builder struct {
	ratios        []float64
	pDefaultValue *uint
}

func createBuilder() Builder {
	out := builder{
		ratios:        nil,
		pDefaultValue: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRatios add ratios to the builder
func (app *builder) WithRatios(ratios []float64) Builder {
	app.ratios = ratios
	return app
}

// WithDefault adds a default value to the builder
func (app *builder) WithDefault(defValue uint) Builder {
	app.pDefaultValue = &defValue
	return app
}

// Now builds a new Config instance
func (app *builder) Now() (Config, error) {
	if app.ratios == nil {
		app.ratios = []float64{}
	}

	if app.pDefaultValue == nil {
		return nil, errors.New("the default value is mandatory in order to build a Config instance")
	}

	for idx, oneRatio := range app.ratios {
		if oneRatio < 0.0 {
			str := fmt.Sprintf("the ratio (index: %d) is smaller than zero (0), but ratio's must be a number between 0 and 1", idx)
			return nil, errors.New(str)
		}

		if oneRatio > 1.0 {
			str := fmt.Sprintf("the ratio (index: %d) is bigger than one (1), but ratio's must be a number between 0 and 1", idx)
			return nil, errors.New(str)
		}
	}

	return createConfig(app.ratios, *app.pDefaultValue), nil
}
