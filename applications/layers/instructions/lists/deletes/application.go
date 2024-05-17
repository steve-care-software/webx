package deletes

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable deletes.Delete) ([]stacks.Assignable, *uint, error) {
	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	listVar := assignable.List()
	listAssignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	index := *pIndex
	list := listAssignables.List()
	amount := uint(len(list))
	if index >= amount {
		code := failures.CouldNotFetchElementFromList
		return nil, &code, nil
	}

	topDelimiter := index + amount
	if topDelimiter >= amount {
		code := failures.CouldNotFetchElementFromList
		return nil, &code, nil
	}

	remainingList := append(list[0:index], list[index+1:topDelimiter]...)
	return remainingList, nil, nil
}
