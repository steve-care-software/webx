package deletes

import (
	"testing"

	instruction_deletes "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	secondBytes := []byte("second bytes")
	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first bytes")),
		stacks.NewAssignableWithBytesForTests(secondBytes),
	})

	index := uint(1)

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

	instruction := instruction_deletes.NewDeleteForTests(listVar, indexVar)
	application := NewApplication()

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

func TestExecute_indexExceedsTopDelimiter_returnsError(t *testing.T) {
	secondBytes := []byte("second bytes")
	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first bytes")),
		stacks.NewAssignableWithBytesForTests(secondBytes),
	})

	index := uint(2)

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

	instruction := instruction_deletes.NewDeleteForTests(listVar, indexVar)
	application := NewApplication()

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotFetchElementFromList {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchElementFromList, *pCode)
		return
	}

}

func TestExecute_listNotInFrame_returnsError(t *testing.T) {
	index := uint(1)

	listVar := "myList"
	indexVar := "myIndex"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := instruction_deletes.NewDeleteForTests(listVar, indexVar)
	application := NewApplication()

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotFetchListFromFrame {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchListFromFrame, *pCode)
		return
	}

}

func TestExecute_indexNotInFrame_returnsError(t *testing.T) {
	secondBytes := []byte("second bytes")
	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first bytes")),
		stacks.NewAssignableWithBytesForTests(secondBytes),
	})

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
		}),
	)

	instruction := instruction_deletes.NewDeleteForTests(listVar, indexVar)
	application := NewApplication()

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotFetchUnsignedIntegerFromFrame {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchUnsignedIntegerFromFrame, *pCode)
		return
	}

}
