package results

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers"
)

func TestResult_withSuccess_Success(t *testing.T) {
	success := NewSuccessForTests(
		[]byte("this is some bytes"),
		layers.NewKindWithPromptForTests(),
	)

	ins := NewResultWithSuccessForTests(success)

	if !ins.IsSuccess() {
		t.Errorf("the result was expected to be success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the result was expected to NOT be failure")
		return
	}

	retSuccess := ins.Success()
	if !reflect.DeepEqual(success, retSuccess) {
		t.Errorf("the returned success is invalid")
		return
	}
}

func TestResult_withFailure_Success(t *testing.T) {
	failure := NewFailureForTests(56, true)
	ins := NewResultWithFailureForTests(failure)

	if ins.IsSuccess() {
		t.Errorf("the result was expected to NOT be success")
		return
	}

	if !ins.IsFailure() {
		t.Errorf("the result was expected to be failure")
		return
	}

	retFailure := ins.Failure()
	if !reflect.DeepEqual(failure, retFailure) {
		t.Errorf("the returned failure is invalid")
		return
	}
}

func TestResult_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
