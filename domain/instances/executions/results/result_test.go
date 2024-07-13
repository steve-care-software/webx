package results

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/success"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
)

func TestResult_withSuccess_Success(t *testing.T) {
	success := success.NewSuccessForTests(
		outputs.NewOutputForTests([]byte("this is some bytes")),
		kinds.NewKindWithPromptForTests(),
	)

	ins := NewResultWithSuccessForTests(success)

	if !ins.IsSuccess() {
		t.Errorf("the result was expected to be success")
		return
	}

	if ins.IsInterruption() {
		t.Errorf("the result was expected to NOT be an interruption")
		return
	}

	retSuccess := ins.Success()
	if !reflect.DeepEqual(success, retSuccess) {
		t.Errorf("the returned success is invalid")
		return
	}
}

func TestResult_withFailure_Success(t *testing.T) {
	interruption := interruptions.NewInterruptionWithStopForTests(45)
	ins := NewResultWithInterruptionForTests(interruption)

	if ins.IsSuccess() {
		t.Errorf("the result was expected to NOT be success")
		return
	}

	if !ins.IsInterruption() {
		t.Errorf("the result was expected to be an interruption")
		return
	}

	retInterruption := ins.Interruption()
	if !reflect.DeepEqual(interruption, retInterruption) {
		t.Errorf("the returned interruption is invalid")
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
