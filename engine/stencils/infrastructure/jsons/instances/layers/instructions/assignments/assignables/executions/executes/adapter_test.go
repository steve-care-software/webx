package executes

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

func TestAdapter_Success(t *testing.T) {
	adapter := NewAdapter()
	ins := executes.NewExecuteForTests(
		"myContext",
		inputs.NewInputWithPathForTests("myPath"),
	)

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withLayer_Success(t *testing.T) {
	adapter := NewAdapter()
	ins := executes.NewExecuteWithLayerForTests(
		"myContext",
		inputs.NewInputWithPathForTests("myPath"),
		"myLayer",
	)

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
