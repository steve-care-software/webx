package programs

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type valueBuilder struct {
	hashAdapter hash.Adapter
	pInput      *uint
	value       Value
	constant    []byte
	execution   Application
	program     Program
}

func createValueBuilder(
	hashAdapter hash.Adapter,
) ValueBuilder {
	out := valueBuilder{
		hashAdapter: hashAdapter,
		pInput:      nil,
		value:       nil,
		constant:    nil,
		execution:   nil,
		program:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *valueBuilder) WithInput(input uint) ValueBuilder {
	app.pInput = &input
	return app
}

// WithValue adds a value to the builder
func (app *valueBuilder) WithValue(value Value) ValueBuilder {
	app.value = value
	return app
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant []byte) ValueBuilder {
	app.constant = constant
	return app
}

// WithExecution adds an execution to the builder
func (app *valueBuilder) WithExecution(execution Application) ValueBuilder {
	app.execution = execution
	return app
}

// WithProgram adds a program to the builder
func (app *valueBuilder) WithProgram(program Program) ValueBuilder {
	app.program = program
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	data := []byte{}
	if app.pInput != nil {
		data = append(data, []byte(fmt.Sprintf("%d", *app.pInput))...)
	}

	if app.value != nil {
		data = append(data, app.value.Hash().Bytes()...)
	}

	if app.constant != nil {
		data = append(data, app.constant...)
	}

	if app.execution != nil {
		data = append(data, app.execution.Hash().Bytes()...)
	}

	if app.program != nil {
		data = append(data, app.program.Hash().Bytes()...)
	}

	pHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pInput != nil {
		content := createContentWithInput(app.pInput)
		return createValue(*pHash, content), nil
	}

	if app.value != nil {
		content := createContentWithValue(app.value)
		return createValue(*pHash, content), nil
	}

	if app.constant != nil {
		content := createContentWithConstant(app.constant)
		return createValue(*pHash, content), nil
	}

	if app.execution != nil {
		content := createContentWithExecution(app.execution)
		return createValue(*pHash, content), nil
	}

	if app.program != nil {
		content := createContentWithProgram(app.program)
		return createValue(*pHash, content), nil
	}

	return nil, errors.New("the Value is invalid")
}
