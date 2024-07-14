package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// NewExecutionWithListForTests creates a new execution with list for tests
func NewExecutionWithListForTests(list string) Execution {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithInitForTests creates a new execution with init for tests
func NewExecutionWithInitForTests(init inits.Init) Execution {
	ins, err := NewBuilder().Create().WithInit(init).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithBeginForTests creates a new execution with begin for tests
func NewExecutionWithBeginForTests(begin begins.Begin) Execution {
	ins, err := NewBuilder().Create().WithBegin(begin).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithExecuteForTests creates a new execution with execute for tests
func NewExecutionWithExecuteForTests(execute executes.Execute) Execution {
	ins, err := NewBuilder().Create().WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithRetrieveForTests creates a new execution with retrieve for tests
func NewExecutionWithRetrieveForTests(retrieve retrieves.Retrieve) Execution {
	ins, err := NewBuilder().Create().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithAmountForTests creates a new execution with amount for tests
func NewExecutionWithAmountForTests(amount amounts.Amount) Execution {
	ins, err := NewBuilder().Create().WithAmount(amount).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithHeadForTests creates a new execution with head for tests
func NewExecutionWithHeadForTests(head heads.Head) Execution {
	ins, err := NewBuilder().Create().WithHead(head).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
