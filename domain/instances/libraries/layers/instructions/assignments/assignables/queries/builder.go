package queries

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries/conditions"
)

type builder struct {
	hashAdapter hash.Adapter
	entity      string
	condition   conditions.Condition
	fields      []string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		entity:      "",
		condition:   nil,
		fields:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEntity adds a entity to the builder
func (app *builder) WithEntity(entity string) Builder {
	app.entity = entity
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// WithFields add fields to the builder
func (app *builder) WithFields(fields []string) Builder {
	app.fields = fields
	return app
}

// Now builds a new Query instance
func (app *builder) Now() (Query, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Query instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build a Query instance")
	}

	data := [][]byte{
		[]byte(app.entity),
		app.condition.Hash().Bytes(),
	}

	if app.fields != nil && len(app.fields) > 0 {
		for _, oneField := range app.fields {
			data = append(data, []byte(oneField))
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.fields != nil && len(app.fields) > 0 {
		return createQueryWithFields(
			*pHash,
			app.entity,
			app.condition,
			app.fields,
		), nil
	}

	return createQuery(
		*pHash,
		app.entity,
		app.condition,
	), nil
}
