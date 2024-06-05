package actions

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {

	modificationsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
		),
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		),
	})

	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				modifVar,
				stacks.NewAssignableWithListForTests(
					modificationsList,
				),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := actions.NewActionForTests(pathVar, modifVar)
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

	if !retAssignable.IsAction() {
		t.Errorf("the assignable was expected to contain an action")
		return
	}

	modifications := retAssignable.Action().Modifications().List()
	if len(modifications) != 2 {
		t.Errorf("the modifications list was expected to contain 2 elements")
		return
	}

}

func TestExecute_modificationsNotInFrame_returnsError(t *testing.T) {
	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := actions.NewActionForTests(pathVar, modifVar)
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

func TestExecute_pathNotInFrame_returnsError(t *testing.T) {

	modificationsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
		),
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				modifVar,
				stacks.NewAssignableWithListForTests(
					modificationsList,
				),
			),
		}),
	)

	instruction := actions.NewActionForTests(pathVar, modifVar)
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

func TestExecute_modificationsContainsInvalidElement_returnsError(t *testing.T) {

	modificationsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
		),
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		),
		stacks.NewAssignableWithStringForTests(
			"invalid",
		),
	})

	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				modifVar,
				stacks.NewAssignableWithListForTests(
					modificationsList,
				),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := actions.NewActionForTests(pathVar, modifVar)
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

	if *pCode != failures.CouldNotFetchModificationFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchModificationFromList)
		return
	}

}

func TestExecute_pathContainsInvalidElement_returnsError(t *testing.T) {

	modificationsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
		),
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		),
	})

	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
		stacks.NewAssignableWithBytesForTests(
			[]byte("invalid"),
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				modifVar,
				stacks.NewAssignableWithListForTests(
					modificationsList,
				),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := actions.NewActionForTests(pathVar, modifVar)
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

	if *pCode != failures.CouldNotFetchStringFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchStringFromList)
		return
	}

}
