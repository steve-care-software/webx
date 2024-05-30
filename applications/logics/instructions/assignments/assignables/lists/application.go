package lists

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder  stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		assignableBuilder:  assignableBuilder,
		assignablesBuilder: assignablesBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable lists.List) (stacks.Assignable, *uint, error) {
	if assignable.IsFetch() {
		fetch := assignable.Fetch()
		indexName := fetch.Index()
		pIndex, err := frame.FetchUnsignedInt(indexName)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return nil, &code, nil
		}

		listName := fetch.List()
		list, err := frame.FetchList(listName)
		if err != nil {
			code := failures.CouldNotFetchListFromFrame
			return nil, &code, nil
		}

		index := *pIndex
		elements := list.List()
		if index >= uint(len(elements)) {
			code := failures.CouldNotFetchElementFromList
			return nil, &code, nil
		}

		return elements[index], nil, nil

	}

	if assignable.IsCreate() {
		elementName := assignable.Create()
		assignable, err := frame.Fetch(elementName)
		if err != nil {
			code := failures.CouldNotFetchFromFrame
			return nil, &code, nil
		}

		retAssignables, err := app.assignablesBuilder.Create().WithList([]stacks.Assignable{
			assignable,
		}).Now()
		if err != nil {
			return nil, nil, err
		}

		retAssignable, err := app.assignableBuilder.Create().WithList(retAssignables).Now()
		if err != nil {
			return nil, nil, err
		}

		return retAssignable, nil, nil
	}

	listVar := assignable.Length()
	list, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, nil
	}

	length := len(list.List())
	retAssignable, err := app.assignableBuilder.Create().WithUnsignedInt(uint(length)).Now()
	if err != nil {
		return nil, nil, err
	}

	return retAssignable, nil, nil
}
