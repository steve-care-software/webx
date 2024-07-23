package contexts

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
)

func TestAdapter_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	executions := []hash.Hash{}
	for i := 0; i < 5; i++ {
		pHash, err := hashAdapter.FromBytes([]byte(fmt.Sprintf("this is execution %d", i)))
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		executions = append(executions, *pHash)
	}

	ins := contexts.NewContextForTests(34, executions)

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

func TestAdapter_withHead_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pHead, err := hashAdapter.FromBytes([]byte("this is the head hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := []hash.Hash{}
	for i := 0; i < 5; i++ {
		pHash, err := hashAdapter.FromBytes([]byte(fmt.Sprintf("this is execution %d", i)))
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		executions = append(executions, *pHash)
	}

	ins := contexts.NewContextWithHeadForTests(34, *pHead, executions)

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
