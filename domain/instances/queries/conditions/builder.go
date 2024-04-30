package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
)

type builder struct {
	hashAdapter hash.Adapter
	pointer     pointers.Pointer
	operator    operators.Operator
	element     Element
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pointer:     nil,
		operator:    nil,
		element:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPointer adds a pointer to the builder
func (app *builder) WithPointer(pointer pointers.Pointer) Builder {
	app.pointer = pointer
	return app
}

// WithOperator adds an operator to the builder
func (app *builder) WithOperator(operator operators.Operator) Builder {
	app.operator = operator
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element Element) Builder {
	app.element = element
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Condition instance")
	}

	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build a Condition instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.pointer.Hash().Bytes(),
		app.operator.Hash().Bytes(),
		app.element.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(
		*pHash,
		app.pointer,
		app.operator,
		app.element,
	), nil
}
