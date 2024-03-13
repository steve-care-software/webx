package queries

import "github.com/steve-care-software/datastencil/domain/instances/queries/conditions"

// NewQueryWithFieldsForTests creates a new query with fields for tests
func NewQueryWithFieldsForTests(entity string, condition conditions.Condition, fields []string) Query {
	ins, err := NewBuilder().Create().WithEntity(entity).WithCondition(condition).WithFields(fields).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryForTests creates a new query for tests
func NewQueryForTests(entity string, condition conditions.Condition) Query {
	ins, err := NewBuilder().Create().WithEntity(entity).WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
