package layers

import (
	"reflect"
	"testing"
)

func TestExecution_Success(t *testing.T) {
	input := "myInput"
	ins := NewExecutionForTests(input)

	if ins.HasLayer() {
		t.Errorf("the execution was expected to NOT contain a compile")
		return
	}

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the input is invalid")
		return
	}
}

func TestExecution_withLayer_Success(t *testing.T) {
	input := "myInput"
	layer := "myLayer"
	ins := NewExecutionWithLayerForTests(input, layer)

	if !ins.HasLayer() {
		t.Errorf("the execution was expected to contain a compile")
		return
	}

	retLayer := ins.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the layer is invalid")
		return
	}

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the input is invalid")
		return
	}
}

func TestExecution_withoutLayer_returnsError(t *testing.T) {
	_, err := NewExecutionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
