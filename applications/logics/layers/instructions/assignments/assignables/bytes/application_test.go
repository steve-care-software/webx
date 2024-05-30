package bytes

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_withJoin_canFetchAllVariablesFromFrame_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithJoinForTests([]string{
		firstVariable,
		secondVariable,
	})

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
	expected := bytes.Join([][]byte{
		firstValue,
		secondValue,
	}, []byte{})

	if !bytes.Equal(expected, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestExecute_withJoin_withVariableNotInFrame_returnsError(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithJoinForTests([]string{
		firstVariable,
		secondVariable,
	})

	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchJoinVariableFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchJoinVariableFromFrame, code)
		return
	}
}

func TestExecute_withCompare_canFetchAllVariablesFromFrame_allEqual_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("someValue")
	secondVariable := "secondVar"
	secondValue := []byte("someValue")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithCompareForTests([]string{
		firstVariable,
		secondVariable,
	})

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

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool")
		return
	}

	pBool := retAssignable.Bool()
	retValue := *pBool
	if !retValue {
		t.Errorf("the returned value was expected to be true, false returned")
		return
	}
}

func TestExecute_withCompare_canFetchAllVariablesFromFrame_notAllEqual_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("someValue")
	secondVariable := "secondVar"
	secondValue := []byte("someValue")
	thirdVariable := "thirdVar"
	thirdValue := []byte("different value")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
			stacks.NewAssignmentForTests(
				thirdVariable,
				stacks.NewAssignableWithBytesForTests(
					thirdValue,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithCompareForTests([]string{
		firstVariable,
		secondVariable,
		thirdVariable,
	})

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

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool")
		return
	}

	pBool := retAssignable.Bool()
	retValue := *pBool
	if retValue {
		t.Errorf("the returned value was expected to be false, true returned")
		return
	}
}

func TestExecute_withCompare_variableDoesNotExistsInFrame_returnsError(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("someValue")
	secondVariable := "secondVar"
	secondValue := []byte("someValue")
	thirdVariable := "thirdVar"
	thirdValue := []byte("different value")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
			stacks.NewAssignmentForTests(
				thirdVariable,
				stacks.NewAssignableWithBytesForTests(
					thirdValue,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithCompareForTests([]string{
		firstVariable,
		secondVariable,
		"invalidVariable",
	})

	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchCompareVariableFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCompareVariableFromFrame, code)
		return
	}
}

func TestExecute_withHashBytes_variableExistsInFrame_Success(t *testing.T) {
	variable := "myVar"
	value := []byte("someValue")
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				variable,
				stacks.NewAssignableWithBytesForTests(
					value,
				),
			),
		}),
	)

	instruction := assignable_bytes.NewBytesWithHashBytesForTests(variable)

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

	if !retAssignable.IsHash() {
		t.Errorf("the assignable was expected to contain an hash")
		return
	}

	pExpectedHash, err := hash.NewAdapter().FromBytes(value)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(pExpectedHash.Bytes(), retAssignable.Hash().Bytes()) {
		t.Errorf("the returned hash is invalid")
		return
	}
}

func TestExecute_withHashBytes_variableDoesNotExistsInFrame_returnsError(t *testing.T) {
	variable := "myVar"
	frame := stacks.NewFrameForTests()
	instruction := assignable_bytes.NewBytesWithHashBytesForTests(variable)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchHashVariableFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchHashVariableFromFrame, code)
		return
	}
}
