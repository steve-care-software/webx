package fetches

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {

	index := uint(2)
	value := []byte("third")

	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first")),
		stacks.NewAssignableWithBytesForTests([]byte("second")),
		stacks.NewAssignableWithBytesForTests(value),
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
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := fetches.NewFetchForTests(listVar, indexVar)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retBytes := retAssignable.Bytes()
	if !bytes.Equal(value, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestExecute_indexExceedsLimit_returnsError(t *testing.T) {

	index := uint(3)
	value := []byte("third")

	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first")),
		stacks.NewAssignableWithBytesForTests([]byte("second")),
		stacks.NewAssignableWithBytesForTests(value),
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
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := fetches.NewFetchForTests(listVar, indexVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchElementFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchElementFromList)
		return
	}
}

func TestExecute_indexNotInFrame_returnsError(t *testing.T) {
	value := []byte("third")

	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first")),
		stacks.NewAssignableWithBytesForTests([]byte("second")),
		stacks.NewAssignableWithBytesForTests(value),
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

	instruction := fetches.NewFetchForTests(listVar, indexVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchUnsignedIntegerFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchElementFromList)
		return
	}
}

func TestExecute_listNotInFrame_returnsError(t *testing.T) {
	index := uint(3)

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

	instruction := fetches.NewFetchForTests(listVar, indexVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchListFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchListFromFrame)
		return
	}
}
