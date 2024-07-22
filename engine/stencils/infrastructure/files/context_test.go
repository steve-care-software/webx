package files

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/contexts"
	json_contexts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/contexts"
)

func TestAdapter_withExecutions_Success(t *testing.T) {
	dbPath := []string{
		"test_files",
		"myDbFile.data",
	}

	endPath := []string{
		"contexts",
	}

	defer func() {
		os.RemoveAll(dbPath[0])
	}()

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

	firstIns := contexts.NewContextForTests(34, *pHead, executions)
	secondIns := contexts.NewContextForTests(12, *pHead, executions)
	adapter := json_contexts.NewAdapter()

	repository := NewContextRepository(
		adapter,
		endPath,
	)

	service := NewContextService(
		adapter,
		endPath,
	)

	_, err = repository.Retrieve(dbPath)
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}

	err = service.Save(dbPath, firstIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstIns, err := repository.Retrieve(dbPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(firstIns.Hash().Bytes(), retFirstIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}

	err = service.Save(dbPath, secondIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSecondIns, err := repository.Retrieve(dbPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(secondIns.Hash().Bytes(), retSecondIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}

	err = service.Delete(dbPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = repository.Retrieve(dbPath)
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}

	err = service.Delete(dbPath)
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
