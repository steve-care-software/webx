package lists

import (
	"testing"

	application_delete "github.com/steve-care-software/datastencil/applications/logics/instructions/lists/deletes"
	applications_inserts "github.com/steve-care-software/datastencil/applications/logics/instructions/lists/inserts"
	instruction_list "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists"
	instruction_deletes "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
	instruction_inserts "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withDelete_Success(t *testing.T) {
	secondBytes := []byte("second bytes")
	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first bytes")),
		stacks.NewAssignableWithBytesForTests(secondBytes),
	})

	index := uint(0)

	listVar := "myList"
	indexVar := "myIndex"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				listVar,
				stacks.NewAssignableWithListForTests(
					list,
				),
			),
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := instruction_list.NewListWithDeleteForTests(
		instruction_deletes.NewDeleteForTests(listVar, indexVar),
	)

	application := NewApplication(
		applications_inserts.NewApplication(),
		application_delete.NewApplication(),
	)

	retAssignment, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	retName := retAssignment.Name()
	if listVar != retName {
		t.Errorf("the name was expected to be '%s', '%s' returned", listVar, retName)
		return
	}

}

func TestExecute_withInsert_Success(t *testing.T) {
	secondBytes := []byte("second bytes")
	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first bytes")),
	})

	listVar := "myList"
	elementVar := "myElement"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				listVar,
				stacks.NewAssignableWithListForTests(
					list,
				),
			),
			stacks.NewAssignmentForTests(
				elementVar,
				stacks.NewAssignableWithBytesForTests(
					secondBytes,
				),
			),
		}),
	)

	instruction := instruction_list.NewListWithInsertForTests(
		instruction_inserts.NewInsertForTests(listVar, elementVar),
	)

	application := NewApplication(
		applications_inserts.NewApplication(),
		application_delete.NewApplication(),
	)

	retAssignment, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	retName := retAssignment.Name()
	if listVar != retName {
		t.Errorf("the name was expected to be '%s', '%s' returned", listVar, retName)
		return
	}

}
