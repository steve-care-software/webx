package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(executable string, content Content) Execution {
	ins, err := NewBuilder().Create().WithExecutable(executable).WithContent(content).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithListForTests creates a new execution with list for tests
func NewContentWithListForTests() Content {
	ins, err := NewContentBuilder().Create().IsList().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithInitForTests creates a new execution with init for tests
func NewContentWithInitForTests(init inits.Init) Content {
	ins, err := NewContentBuilder().Create().WithInit(init).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithBeginForTests creates a new execution with begin for tests
func NewContentWithBeginForTests(begin begins.Begin) Content {
	ins, err := NewContentBuilder().Create().WithBegin(begin).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithExecuteForTests creates a new execution with execute for tests
func NewContentWithExecuteForTests(execute executes.Execute) Content {
	ins, err := NewContentBuilder().Create().WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithRetrieveForTests creates a new execution with retrieve for tests
func NewContentWithRetrieveForTests(retrieve retrieves.Retrieve) Content {
	ins, err := NewContentBuilder().Create().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithAmountForTests creates a new execution with amount for tests
func NewContentWithAmountForTests(amount amounts.Amount) Content {
	ins, err := NewContentBuilder().Create().WithAmount(amount).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithHeadForTests creates a new execution with head for tests
func NewContentWithHeadForTests(head heads.Head) Content {
	ins, err := NewContentBuilder().Create().WithHead(head).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
