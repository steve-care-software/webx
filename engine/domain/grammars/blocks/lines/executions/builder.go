package executions

import (
	"errors"
	"fmt"
)

type builder struct {
	flags   []uint16
	tokens  []string
	pFnFLag *uint16
}

func createBuilder(
	flags []uint16,
) Builder {
	out := builder{
		flags:   flags,
		tokens:  nil,
		pFnFLag: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.flags,
	)
}

// WithTokens add tokens to the buiilder
func (app *builder) WithTokens(tokens []string) Builder {
	app.tokens = tokens
	return app
}

// WithFuncFlag add fn flag to the buiilder
func (app *builder) WithFuncFlag(fnFlag uint16) Builder {
	app.pFnFLag = &fnFlag
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.tokens == nil {
		app.tokens = []string{}
	}

	if app.pFnFLag == nil {
		return nil, errors.New("the func flag is mandatory in order to build an Execution instance")
	}

	isValid := false
	currentFlag := *app.pFnFLag
	for _, oneFlag := range app.flags {
		if currentFlag == oneFlag {
			isValid = true
			break
		}
	}

	if !isValid {
		str := fmt.Sprintf("the provided flag (%d) is invalid", currentFlag)
		return nil, errors.New(str)
	}

	return createExecution(app.tokens, currentFlag), nil
}
