package sqllites

import (
	"bytes"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/orms"
)

type instanceExec struct {
	name     string
	insatnce orms.Instance
}

type testInstance struct {
	path     []string
	instance orms.Instance
}

func TestOrm_Success(t *testing.T) {
	dbDir := "./test_files"
	dbName := "testdb"
	basePath := filepath.Join(dbDir, dbName)
	defer func() {
		os.Remove(basePath)
	}()

	skeleton, err := NewSkeletonFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// build resources:
	instances := []testInstance{
		{
			path: []string{
				"libraries",
				"layers",
				"links",
			},
			instance: links.NewOperatorWithAndForTests(),
		},
	}

	pDB, err := sql.Open("sqlite3", basePath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	repository := NewOrmRepository(skeleton, pDB)
	for idx, oneInstance := range instances {
		pTx, err := pDB.Begin()
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		// create the service:
		service := NewOrmService(skeleton, pDB, pTx)

		// init our service:
		err = service.Init()
		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		// insert instance:
		err = service.Insert(oneInstance.instance, oneInstance.path)

		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		// commit:
		err = pTx.Commit()
		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		insHash := oneInstance.instance.Hash()
		retInstance, err := repository.Retrieve(oneInstance.path, insHash)
		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		if !bytes.Equal(retInstance.Hash().Bytes(), insHash.Bytes()) {
			t.Errorf("index: %d, the returned insatnce is invalid", idx)
			return
		}
	}
}