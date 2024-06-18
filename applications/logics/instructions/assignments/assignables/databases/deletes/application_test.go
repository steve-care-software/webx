package deletes

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	index := uint(0)
	length := uint(32)

	indexVar := "myIndex"
	lengthVar := "myLength"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
			stacks.NewAssignmentForTests(
				lengthVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					length,
				),
			),
		}),
	)

	instruction := deletes.NewDeleteForTests(indexVar, lengthVar)
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

	if !retAssignable.IsDelete() {
		t.Errorf("the assignable was expected to contain a delete")
		return
	}

	delete := retAssignable.Delete()
	if index != delete.Index() {
		t.Errorf("the index was expected to be %d, %d returned", index, delete.Index())
		return
	}

	if length != delete.Length() {
		t.Errorf("the length was expected to be %d, %d returned", length, delete.Length())
		return
	}
}

func TestExecute_indexNotInFrame_returnsError(t *testing.T) {
	length := uint(32)

	indexVar := "myIndex"
	lengthVar := "myLength"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				lengthVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					length,
				),
			),
		}),
	)

	instruction := deletes.NewDeleteForTests(indexVar, lengthVar)
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
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchUnsignedIntegerFromFrame)
		return
	}
}

func TestExecute_lengthNotInFrame_returnsError(t *testing.T) {
	index := uint(0)

	indexVar := "myIndex"
	lengthVar := "myLength"

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

	instruction := deletes.NewDeleteForTests(indexVar, lengthVar)
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
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchUnsignedIntegerFromFrame)
		return
	}
}
