package sqllites

import (
	"bytes"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers"
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
	instances := map[string][]testInstance{
		"layer": {
			{
				path: []string{
					"layer",
				},
				instance: layers.NewLayerForTests(
					layers.NewInstructionsForTests([]layers.Instruction{
						layers.NewInstructionWithAssignmentForTests(
							layers.NewAssignmentForTests(
								"myName",
								layers.NewAssignableWithBytesForTests(
									layers.NewBytesWithHashBytesForTests("myInput"),
								),
							),
						),
					}),
					layers.NewOutputForTests(
						"myVariable",
						layers.NewKindWithContinueForTests(),
					),
					"some input",
				),
			},
		},
		"output": {
			{
				path: []string{
					"layer",
					"output",
				},
				instance: layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithContinueForTests(),
				),
			},
			{
				path: []string{
					"layer",
					"output",
				},
				instance: layers.NewOutputWithExecuteForTests(
					"myVariable",
					layers.NewKindWithContinueForTests(),
					"some command to execute",
				),
			},
		},
		"kind": {
			{
				path: []string{
					"layer",
					"output",
					"kind",
				},
				instance: layers.NewKindWithContinueForTests(),
			},
			{
				path: []string{
					"layer",
					"output",
					"kind",
				},
				instance: layers.NewKindWithPromptForTests(),
			},
		},
		"instruction": {
			{
				path: []string{
					"layer",
					"instruction",
				},
				instance: layers.NewInstructionWithAssignmentForTests(
					layers.NewAssignmentForTests(
						"myName",
						layers.NewAssignableWithBytesForTests(
							layers.NewBytesWithHashBytesForTests("myInput"),
						),
					),
				),
			},
		},
		"assignment": {
			{
				path: []string{
					"layer",
					"instruction",
					"assignment",
				},
				instance: layers.NewAssignmentForTests(
					"myName",
					layers.NewAssignableWithBytesForTests(
						layers.NewBytesWithHashBytesForTests("myInput"),
					),
				),
			},
		},
		"assignable": {
			{
				path: []string{
					"layer",
					"instruction",
					"assignment",
					"assignable",
				},
				instance: layers.NewAssignableWithBytesForTests(
					layers.NewBytesWithHashBytesForTests("myInput"),
				),
			},
		},
		"bytes": {
			{
				path: []string{
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers.NewBytesWithHashBytesForTests("myInput"),
			},
			{
				path: []string{
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers.NewBytesWithCompareForTests([]string{
					"first",
					"second",
				}),
			},
			{
				path: []string{
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				},
				instance: layers.NewBytesWithJoinForTests([]string{
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

	repository := NewOrmRepository(skeleton, pDB)
	for name, oneSection := range instances {
		for idx, oneInstance := range oneSection {
			pTx, err := pDB.Begin()
			if err != nil {
				t.Errorf("section: %s: the error was expected to be nil, error returned: %s", name, err.Error())
				return
			}

			// create the service:
			service := NewOrmService(repository, skeleton, pDB, pTx)

			// init our service:
			err = service.Init()
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			// insert instance:
			err = service.Insert(oneInstance.instance, oneInstance.path)
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			// commit:
			err = pTx.Commit()
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			insHash := oneInstance.instance.Hash()
			retInstance, err := repository.Retrieve(oneInstance.path, insHash)
			if err != nil {
				t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
				return
			}

			if !bytes.Equal(retInstance.Hash().Bytes(), insHash.Bytes()) {
				t.Errorf("section: %s: index: %d, the returned instance is invalid", name, idx)
				return
			}
		}
	}
}
