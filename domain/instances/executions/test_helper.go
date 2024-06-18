package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
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
func NewExecutionForTests(link links.Link, database databases.Database) Execution {
	ins, err := NewExecutionBuilder().Create().WithLogic(link).WithDatabase(database).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
