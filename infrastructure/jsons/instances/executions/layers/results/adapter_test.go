package results

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/success"
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
)

func TestAdapter_withSuccess_Success(t *testing.T) {

	ins := results.NewResultWithSuccessForTests(
		success.NewSuccessForTests(
			outputs.NewOutputForTests(
				[]byte("this is an input"),
			),
			kinds.NewKindWithPromptForTests(),
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

func TestAdapter_withInterruption_Success(t *testing.T) {

	ins := results.NewResultWithInterruptionForTests(
		interruptions.NewInterruptionWithStopForTests(
			23,
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
