package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        string
	init        inits.Init
	begin       begins.Begin
	execute     executes.Execute
	retrieve    retrieves.Retrieve
	amount      amounts.Amount
	head        heads.Head
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        "",
		init:        nil,
		begin:       nil,
		execute:     nil,
		retrieve:    nil,
		amount:      nil,
		head:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list string) Builder {
	app.list = list
	return app
}

// WithInit adds an init to the builder
func (app *builder) WithInit(init inits.Init) Builder {
	app.init = init
	return app
}

// WithBegin adds a begin to the builder
func (app *builder) WithBegin(begin begins.Begin) Builder {
	app.begin = begin
	return app
}

// WithExecute adds an execute to the builder
func (app *builder) WithExecute(execute executes.Execute) Builder {
	app.execute = execute
	return app
}

// WithRetrieve adds a retrieve to the builder
func (app *builder) WithRetrieve(retrieve retrieves.Retrieve) Builder {
	app.retrieve = retrieve
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount amounts.Amount) Builder {
	app.amount = amount
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	bytes := [][]byte{}
	if app.list != "" {
		bytes = append(bytes, []byte("list"))
		bytes = append(bytes, []byte(app.list))
	}

	if app.init != nil {
		bytes = append(bytes, []byte("init"))
		bytes = append(bytes, app.init.Hash().Bytes())
	}

	if app.begin != nil {
		bytes = append(bytes, []byte("begin"))
		bytes = append(bytes, app.begin.Hash().Bytes())
	}

	if app.execute != nil {
		bytes = append(bytes, []byte("execute"))
		bytes = append(bytes, app.execute.Hash().Bytes())
	}

	if app.retrieve != nil {
		bytes = append(bytes, []byte("retrieve"))
		bytes = append(bytes, app.retrieve.Hash().Bytes())
	}

	if app.amount != nil {
		bytes = append(bytes, []byte("amount"))
		bytes = append(bytes, app.amount.Hash().Bytes())
	}

	if app.head != nil {
		bytes = append(bytes, []byte("head"))
		bytes = append(bytes, app.head.Hash().Bytes())
	}

	if len(bytes) != 2 {
		return nil, errors.New("the Execution is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.list != "" {
		return createExecutionWithList(*pHash, app.list), nil
	}

	if app.init != nil {
		return createExecutionWithInit(*pHash, app.init), nil
	}

	if app.begin != nil {
		return createExecutionWithBegin(*pHash, app.begin), nil
	}

	if app.execute != nil {
		return createExecutionWithExecute(*pHash, app.execute), nil
	}

	if app.retrieve != nil {
		return createExecutionWithRetrieve(*pHash, app.retrieve), nil
	}

	if app.amount != nil {
		return createExecutionWithAmount(*pHash, app.amount), nil
	}

	return createExecutionWithHead(*pHash, app.head), nil
}
