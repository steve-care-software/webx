package constants

import (
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/constants"
)

func TestExecute_withBool_Success(t *testing.T) {
	value := false
	constant := constants.NewConstantWithBoolForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool value")
		return
	}

	pValue := retAssignable.Bool()
	if *pValue != value {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestExecute_withString_Success(t *testing.T) {
	value := "this is some string"
	constant := constants.NewConstantWithStringForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsString() {
		t.Errorf("the assignable was expected to contain a string value")
		return
	}

	pString := retAssignable.String()
	if *pString != value {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestExecute_withInt_Success(t *testing.T) {
	value := 34
	constant := constants.NewConstantWithIntForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsInt() {
		t.Errorf("the assignable was expected to contain an int value")
		return
	}

	pInt := retAssignable.Int()
	if *pInt != value {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestExecute_withUint_Success(t *testing.T) {
	value := uint(22)
	constant := constants.NewConstantWithUintForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsUnsignedInt() {
		t.Errorf("the assignable was expected to contain an uint value")
		return
	}

	pInt := retAssignable.UnsignedInt()
	if *pInt != value {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestExecute_withFloat_Success(t *testing.T) {
	value := 23.456
	constant := constants.NewConstantWithFloatForTests(value)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsFloat() {
		t.Errorf("the assignable was expected to contain a float value")
		return
	}

	pFloat := retAssignable.Float()
	if *pFloat != value {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestExecute_withList_Success(t *testing.T) {
	list := []constants.Constant{
		constants.NewConstantWithBoolForTests(true),
		constants.NewConstantWithStringForTests(""),
		constants.NewConstantWithIntForTests(32),
		constants.NewConstantWithUintForTests(uint(21)),
		constants.NewConstantWithFloatForTests(23.43),
	}

	constant := constants.NewConstantWithListForTests(
		constants.NewConstantsForTests(list),
	)

	application := NewApplication()
	retAssignable, pCode, err := application.Execute(constant)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsList() {
		t.Errorf("the assignable was expected to contain a list")
		return
	}

	retList := retAssignable.List().List()
	if len(list) != len(retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}
