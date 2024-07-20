package outputs

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/outputs/kinds"
)

func TestAdapter_Success(t *testing.T) {
	ins := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)

	adapter := NewAdapter()

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

func TestAdapter_withExecute_Success(t *testing.T) {
	ins := outputs.NewOutputWithExecuteForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
		[]string{
			"firstCommand",
			"secondCommand",
		},
	)

	adapter := NewAdapter()

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
