package fetches

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable fetches.Fetch) (stacks.Assignable, *uint, error) {
	listVar := assignable.List()
	assignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, nil
	}

	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	index := *pIndex
	list := assignables.List()
	if index >= uint(len(list)) {
		code := failures.CouldNotFetchElementFromList
		return nil, &code, nil
	}

	return list[index], nil, nil
}
