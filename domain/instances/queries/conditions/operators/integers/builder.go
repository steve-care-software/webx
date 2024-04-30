package integers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		isSmallerThan: false,
		isBiggerThan:  false,
		isEqual:       false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsSmallerThan adds a smaller than to the builder
func (app *builder) IsSmallerThan() Builder {
	app.isSmallerThan = true
	return app
}

// IsBiggerThan adds a bigger than to the builder
func (app *builder) IsBiggerThan() Builder {
	app.isBiggerThan = true
	return app
}

// IsEqual adds an equal to the builder
func (app *builder) IsEqual() Builder {
	app.isEqual = true
	return app
}

// Now builds a new Integer instance
func (app *builder) Now() (Integer, error) {
	isSmallerThan := "false"
	isBiggerThan := "false"
	isEqual := "false"
	isValid := false
	if app.isSmallerThan {
		isSmallerThan = "true"
		isValid = true
	}

	if app.isBiggerThan {
		isBiggerThan = "true"
		isValid = true
	}

	if app.isEqual {
		isEqual = "true"
		isValid = true
	}

	if !isValid {
		return nil, errors.New("the Integer is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(isSmallerThan),
		[]byte(isBiggerThan),
		[]byte(isEqual),
	})

	if err != nil {
		return nil, err
	}

	if app.isSmallerThan && app.isEqual {
		return createIntegerWithSmallerThanAndEqual(*pHash), nil
	}

	if app.isSmallerThan {
		return createIntegerWithSmallerThan(*pHash), nil
	}

	if app.isBiggerThan && app.isEqual {
		return createIntegerWithBiggerThanAndEqual(*pHash), nil
	}

	if app.isBiggerThan {
		return createIntegerWithBiggerThan(*pHash), nil
	}

	return createIntegerWithEqual(*pHash), nil
}
