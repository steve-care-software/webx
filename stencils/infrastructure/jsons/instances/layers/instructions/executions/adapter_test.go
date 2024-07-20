package executions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions/merges"
)

func TestAdapter_withMerge_Success(t *testing.T) {
	ins := executions.NewExecutionForTests(
		"myExecutable",
		executions.NewContentWithMergeForTests(
			merges.NewMergeForTests("myBase", "myTop"),
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

func TestAdapter_withCommit_Success(t *testing.T) {
	ins := executions.NewExecutionForTests(
		"myExecutable",
		executions.NewContentWithCommitForTests("myCommit"),
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

func TestAdapter_withRollback_Success(t *testing.T) {
	ins := executions.NewExecutionForTests(
		"myExecutable",
		executions.NewContentWithRollbackForTests("myRollback"),
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

func TestAdapter_withCancel_Success(t *testing.T) {
	ins := executions.NewExecutionForTests(
		"myExecutable",
		executions.NewContentWithCancelForTests("myCancel"),
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
