package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers"
	"github.com/steve-care-software/historydb/domain/databases"
)

// NewExecutionsForTests creates a new executions for tests
func NewExecutionsForTests(list []Execution) Executions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(layer layers.Layer, database databases.Database) Execution {
	ins, err := NewExecutionBuilder().Create().WithLayer(layer).WithDatabase(database).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
