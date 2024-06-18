package modifications

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_withInsert_Success(t *testing.T) {
	insert := []byte("this is some bytes to insert")
	insertVar := "myInsert"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				insertVar,
				stacks.NewAssignableWithBytesForTests(
					insert,
				),
			),
		}),
	)

	instruction := modifications.NewModificationWithInsertForTests(insertVar)
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

	if !retAssignable.IsModification() {
		t.Errorf("the assignable was expected to contain a modification")
		return
	}

	modification := retAssignable.Modification()
	if !modification.IsInsert() {
		t.Errorf("the assignable was expected to contain an insert modification")
		return
	}

	retInsert := modification.Insert()
	if !bytes.Equal(insert, retInsert) {
		t.Errorf("the returned insert is invalid")
		return
	}
}

func TestExecute_withDelete_Success(t *testing.T) {
	delete := deletes.NewDeleteForTests(uint(23), uint(44))
	deleteVar := "myInsert"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				deleteVar,
				stacks.NewAssignableWithDeleteForTests(
					delete,
				),
			),
		}),
	)

	instruction := modifications.NewModificationWithDeleteForTests(deleteVar)
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

	if !retAssignable.IsModification() {
		t.Errorf("the assignable was expected to contain a modification")
		return
	}

	modification := retAssignable.Modification()
	if !modification.IsDelete() {
		t.Errorf("the assignable was expected to contain a delete modification")
		return
	}

	retDelete := modification.Delete()
	if !reflect.DeepEqual(delete, retDelete) {
		t.Errorf("the returned delete is invalid")
		return
	}
}

func TestExecute_withInsert_insertNotInFrame_returnsError(t *testing.T) {
	insertVar := "myInsert"

	frame := stacks.NewFrameForTests()

	instruction := modifications.NewModificationWithInsertForTests(insertVar)
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

	if *pCode != failures.CouldNotFetchBytesFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchBytesFromFrame)
		return
	}
}

func TestExecute_withDelete_deleteNotInFrame_returnsError(t *testing.T) {
	deleteVar := "myInsert"

	frame := stacks.NewFrameForTests()

	instruction := modifications.NewModificationWithDeleteForTests(deleteVar)
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

	if *pCode != failures.CouldNotFetchDeleteFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchDeleteFromFrame)
		return
	}
}
