package inserts

import (
	"bytes"
	"testing"

	instruction_inserts "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
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

	instruction := instruction_inserts.NewInsertForTests(listVar, elementVar)
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

	retAssignable := retAssignment.Assignable()
	if !retAssignable.IsList() {
		t.Errorf("the returned assignable was expected to be a list")
		return
	}

	retList := retAssignable.List().List()
	if len(retList) != len(list.List())+1 {
		t.Errorf("the returned list is invalid, %d elements expected, %d returned", len(list.List())+1, len(retList))
		return
	}

	retBytes := retList[1].Bytes()
	if !bytes.Equal(secondBytes, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}

}

func TestExecute_elementNotInFrame_returnsError(t *testing.T) {
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
		}),
	)

	instruction := instruction_inserts.NewInsertForTests(listVar, elementVar)
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

	if *pCode != failures.CouldNotFetchFromFrame {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchFromFrame, *pCode)
		return
	}

}

func TestExecute_listNotInFrame_returnsError(t *testing.T) {
	secondBytes := []byte("second bytes")
	listVar := "myList"
	elementVar := "myElement"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				elementVar,
				stacks.NewAssignableWithBytesForTests(
					secondBytes,
				),
			),
		}),
	)

	instruction := instruction_inserts.NewInsertForTests(listVar, elementVar)
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
