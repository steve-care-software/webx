package fetches

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks/failures"
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
		return nil, &code, err
	}

	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, err
	}

	index := *pIndex
	list := assignables.List()
	limit := uint(len(list) - 1)
	if index > limit {
		code := failures.CouldNotFetchElementFromList
		str := fmt.Sprintf("the provided index (%d) exceeds the limit index (%d) of the list (name: %s)", index, limit, listVar)
		return nil, &code, errors.New(str)
	}

	return list[index], nil, nil
}
