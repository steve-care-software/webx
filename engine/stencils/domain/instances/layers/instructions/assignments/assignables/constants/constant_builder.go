package constants

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type constantBuilder struct {
	hashAdapter hash.Adapter
	pBool       *bool
	pString     *string
	pInt        *int
	pUint       *uint
	pFloat      *float64
	list        Constants
}

func createConstantBuilder(
	hashAdapter hash.Adapter,
) ConstantBuilder {
	out := constantBuilder{
		hashAdapter: hashAdapter,
		pBool:       nil,
		pString:     nil,
		pInt:        nil,
		pUint:       nil,
		pFloat:      nil,
		list:        nil,
	}

	return &out
}

// Create initializes the constantBuilder
func (app *constantBuilder) Create() ConstantBuilder {
	return createConstantBuilder(
		app.hashAdapter,
	)
}

// WithBool adds a bool to the constantBuilder
func (app *constantBuilder) WithBool(boolValue bool) ConstantBuilder {
	app.pBool = &boolValue
	return app
}

// WithString adds a string to the constantBuilder
func (app *constantBuilder) WithString(strValue string) ConstantBuilder {
	app.pString = &strValue
	return app
}

// WithInt adds an int to the constantBuilder
func (app *constantBuilder) WithInt(intValue int) ConstantBuilder {
	app.pInt = &intValue
	return app
}

// WithUint adds a uint to the constantBuilder
func (app *constantBuilder) WithUint(uintValue uint) ConstantBuilder {
	app.pUint = &uintValue
	return app
}

// WithFloat adds a float to the constantBuilder
func (app *constantBuilder) WithFloat(floatVal float64) ConstantBuilder {
	app.pFloat = &floatVal
	return app
}

// WithList adds a list to the constantBuilder
func (app *constantBuilder) WithList(list Constants) ConstantBuilder {
	app.list = list
	return app
}

// Now builds a new Constant instance
func (app *constantBuilder) Now() (Constant, error) {
	data := [][]byte{}
	if app.pBool != nil {
		value := "false"
		if *app.pBool {
			value = "true"
		}

		data = append(data, []byte("bool"))
		data = append(data, []byte(value))
	}

	if app.pString != nil {
		data = append(data, []byte("string"))
		data = append(data, []byte(*app.pString))
	}

	if app.pInt != nil {
		data = append(data, []byte("int"))
		data = append(data, []byte(strconv.Itoa(*app.pInt)))
	}

	if app.pUint != nil {
		data = append(data, []byte("uint"))
		data = append(data, []byte(strconv.Itoa(int(*app.pUint))))
	}

	if app.pFloat != nil {
		data = append(data, []byte("float"))
		data = append(data, []byte(strconv.FormatFloat(*app.pFloat, 'f', 10, 64)))
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		data = append(data, app.list.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Constant is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pBool != nil {
		return createConstantWithBool(*pHash, app.pBool), nil
	}

	if app.pString != nil {
		return createConstantWithString(*pHash, app.pString), nil
	}

	if app.pInt != nil {
		return createConstantWithInt(*pHash, app.pInt), nil
	}

	if app.pUint != nil {
		return createConstantWithUint(*pHash, app.pUint), nil
	}

	if app.pFloat != nil {
		return createConstantWithFloat(*pHash, app.pFloat), nil
	}

	return createConstantWithList(*pHash, app.list), nil
}
