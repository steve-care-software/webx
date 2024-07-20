package interruptions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/interruptions/failures"
)

func TestAdapter_withStop_Success(t *testing.T) {

	ins := interruptions.NewInterruptionWithStopForTests(
		23,
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

func TestAdapter_withFailure_Success(t *testing.T) {

	ins := interruptions.NewInterruptionWithFailureForTests(
		failures.NewFailureForTests(
			uint(34),
			uint(32),
			false,
		),
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
