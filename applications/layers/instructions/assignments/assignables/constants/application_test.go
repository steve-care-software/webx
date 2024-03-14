package constants

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/constants"
)

func TestExecute_withBool_isTrue_Success(t *testing.T) {
	instruction := constants.NewConstantWithBoolForTests(true)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain bool")
		return
	}

	pValue := retAssignable.Bool()
	if !*pValue {
		t.Errorf("the value was expected to be true")
		return
	}
}

func TestExecute_withBool_isFalse_Success(t *testing.T) {
	instruction := constants.NewConstantWithBoolForTests(false)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain bool")
		return
	}

	pValue := retAssignable.Bool()
	if *pValue {
		t.Errorf("the value was expected to be false")
		return
	}
}

func TestExecute_withBytes_Success(t *testing.T) {
	value := []byte("this is some bytes")
	instruction := constants.NewConstantWithBytesForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(instruction)
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

	retValue := retAssignable.Bytes()
	if !bytes.Equal(value, retValue) {
		t.Errorf("the returned value is invalid")
		return
	}
}
