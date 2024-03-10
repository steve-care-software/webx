package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type integerOperatorBuilder struct {
	hashAdapter   hash.Adapter
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createIntegerOperatorBuilder(
	hashAdapter hash.Adapter,
) IntegerOperatorBuilder {
	out := integerOperatorBuilder{
		hashAdapter:   hashAdapter,
		isSmallerThan: false,
		isBiggerThan:  false,
		isEqual:       false,
	}

	return &out
}

// Create initializes the builder
func (app *integerOperatorBuilder) Create() IntegerOperatorBuilder {
	return createIntegerOperatorBuilder(
		app.hashAdapter,
	)
}

// IsSmallerThan adds a smaller than to the builder
func (app *integerOperatorBuilder) IsSmallerThan() IntegerOperatorBuilder {
	app.isSmallerThan = true
	return app
}

// IsBiggerThan adds a bigger than to the builder
func (app *integerOperatorBuilder) IsBiggerThan() IntegerOperatorBuilder {
	app.isBiggerThan = true
	return app
}

// IsEqual adds an equal to the builder
func (app *integerOperatorBuilder) IsEqual() IntegerOperatorBuilder {
	app.isEqual = true
	return app
}

// Now builds a new IntegerOperator instance
func (app *integerOperatorBuilder) Now() (IntegerOperator, error) {
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
		return nil, errors.New("the IntegerOperator is invalid")
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
		return createIntegerOperatorWithSmallerThanAndEqual(*pHash), nil
	}

	if app.isSmallerThan {
		return createIntegerOperatorWithSmallerThan(*pHash), nil
	}

	if app.isBiggerThan && app.isEqual {
		return createIntegerOperatorWithBiggerThanAndEqual(*pHash), nil
	}

	if app.isBiggerThan {
		return createIntegerOperatorWithBiggerThan(*pHash), nil
	}

	return createIntegerOperatorWithEqual(*pHash), nil
}
