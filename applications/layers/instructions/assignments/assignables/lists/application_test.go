package lists

import (
	"bytes"
	"testing"

	application_fetches "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/lists/fetches"
	instructions_list "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists"
	instructions_fetches "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_withFetch_Success(t *testing.T) {
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

	instruction := instructions_list.NewListWithFetchForTests(
		instructions_fetches.NewFetchForTests(listVar, indexVar),
	)

	application := NewApplication(
		application_fetches.NewApplication(),
	)

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

func TestExecute_withCreate_Success(t *testing.T) {
	value := []byte("third")
	createVar := "myElement"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				createVar,
				stacks.NewAssignableWithBytesForTests(
					value,
				),
			),
		}),
	)

	instruction := instructions_list.NewListWithCreateForTests(
		createVar,
	)

	application := NewApplication(
		application_fetches.NewApplication(),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsList() {
		t.Errorf("the assignable was expected to contains a list")
		return
	}

	retList := retAssignable.List().List()
	if len(retList) != 1 {
		t.Errorf("the returned list was expected to contain 1 element, %d returned", len(retList))
		return
	}

	if !bytes.Equal(retList[0].Bytes(), value) {
		t.Errorf("the returned list element is invalid")
		return
	}
}

func TestExecute_withCreate_firstElementNotInFrame_returnsError(t *testing.T) {
	createVar := "myElement"

	frame := stacks.NewFrameForTests()
	instruction := instructions_list.NewListWithCreateForTests(
		createVar,
	)

	application := NewApplication(
		application_fetches.NewApplication(),
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchFromFrame)
		return
	}
}

func TestExecute_withLength_Success(t *testing.T) {

	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("this is a value")),
		stacks.NewAssignableWithBytesForTests([]byte("this is a second value")),
	})

	listVar := "myList"

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

	instruction := instructions_list.NewListWithLengthForTests(
		listVar,
	)

	application := NewApplication(
		application_fetches.NewApplication(),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsUnsignedInt() {
		t.Errorf("the assignable was expected to contains an unsigned int")
		return
	}

	retInt := retAssignable.UnsignedInt()
	if *retInt != 2 {
		t.Errorf("the length was expected to be %d, %d returned", 2, *retInt)
		return
	}
}

func TestExecute_withLength_listhNotInFrame_returnsError(t *testing.T) {
	listVar := "myList"
	frame := stacks.NewFrameForTests()
	instruction := instructions_list.NewListWithLengthForTests(
		listVar,
	)

	application := NewApplication(
		application_fetches.NewApplication(),
	)

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
