package success

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs/kinds"
)

func TestSuccess_Success(t *testing.T) {
	value := outputs.NewOutputForTests([]byte("this is some bytes"))
	kind := kinds.NewKindWithPromptForTests()
	ins := NewSuccessForTests(value, kind)
	retOutput := ins.Output()
	if !reflect.DeepEqual(value, retOutput) {
		t.Errorf("the returned output are invalid")
		return
	}

	retKind := ins.Kind()
	if !reflect.DeepEqual(kind, retKind) {
		t.Errorf("the returned kind is invalid")
		return
	}
}

func TestSuccess_withoutOutput_returnsError(t *testing.T) {
	kind := kinds.NewKindWithPromptForTests()
	_, err := NewBuilder().Create().WithKind(kind).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestSuccess_withoutKind_returnsError(t *testing.T) {
	value := outputs.NewOutputForTests([]byte("this is some bytes"))
	_, err := NewBuilder().Create().WithOutput(value).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
