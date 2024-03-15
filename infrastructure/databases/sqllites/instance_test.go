package sqllites

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances"
    "github.com/steve-care-software/datastencil/domain/accounts/signers"
	layers_bytes "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
)

type testInstance struct {
	path     []string
	instance instances.Instance
}

func TestInstance_Success(t *testing.T) {
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
	instances := map[string][]testInstance{
		"link_layer_instruction_assignment_assignable_bytes": {
			{
				path: []string{
					"link",
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers_bytes.NewBytesWithHashBytesForTests("myInput"),
			},
			{
				path: []string{
					"link",
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers_bytes.NewBytesWithCompareForTests([]string{
					"first",
					"second",
				}),
			},
			{
				path: []string{
					"link",
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers_bytes.NewBytesWithJoinForTests([]string{
					"first",
					"second",
				}),
			},
		},
	}

	pDB, err := sql.Open("sqlite3", basePath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

    signer:=  signers.NewFactory().Create()
	repository := NewInstanceRepository(skeleton, pDB)
	service := NewInstanceService(repository, signer, skeleton, pDB)

	// init our service:
	err = service.Init()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	for name, oneSection := range instances {
		for idx, oneInstance := range oneSection {
            // begin the context:
            pContext, err := service.Begin()
            if err != nil {
                t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
            }

			// insert instance:
        	err = service.Insert(*pContext, oneInstance.instance, oneInstance.path)
        	if err != nil {
        		t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, fmt.Sprintln(err.Error()))
                return
        	}

			// commit:
			err = service.Commit(*pContext)
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			insHash := oneInstance.instance.Hash()
			retInstance, err := repository.RetrieveByPathAndHash(oneInstance.path, insHash)
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			if !bytes.Equal(retInstance.Hash().Bytes(), insHash.Bytes()) {
				t.Errorf("section: %s: index: %d, the returned instance is invalid", name, idx)
				return
			}

			// revert:
			err = service.Revert()
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}
		}
	}
}
