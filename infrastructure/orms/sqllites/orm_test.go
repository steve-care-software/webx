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

type testInstance struct {
	path         []string
	instance     orms.Instance
	dependencies []testInstance
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

	/*pHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}*/

	// build resources:
	instances := map[string][]testInstance{
		/*"library_link_origin": {
			{
				path: []string{
					"library",
					"link",
					"origin",
				},
				instance: links.NewOriginForTests(
					links.NewOriginResourceForTests(
						*pHash,
					),
					links.NewOperatorWithAndForTests(),
					links.NewOriginValueWithResourceForTests(
						links.NewOriginResourceForTests(
							*pHash,
						),
					),
				),
				dependencies: []testInstance{
					{
						path: []string{
							"library",
							"link",
							"origin",
							"resource",
						},
						instance: links.NewOriginResourceForTests(
							*pHash,
						),
					},
					{
						path: []string{
							"library",
							"link",
							"origin",
							"operator",
						},
						instance: links.NewOperatorWithAndForTests(),
					},
					{
						path: []string{
							"library",
							"link",
							"origin",
							"resource",
						},
						instance: links.NewOriginValueWithResourceForTests(
							links.NewOriginResourceForTests(
								*pHash,
							),
						),
						dependencies: []testInstance{
							{
								path: []string{
									"library",
									"link",
									"origin",
									"resource",
								},
								instance: links.NewOriginResourceForTests(
									*pHash,
								),
							},
						},
					},
				},
			},
		},
		"library_link_origin_value": {
			{
				path: []string{
					"library",
					"link",
					"origin",
					"resource",
				},
				instance: links.NewOriginValueWithResourceForTests(
					links.NewOriginResourceForTests(
						*pHash,
					),
				),
				dependencies: []testInstance{
					{
						path: []string{
							"library",
							"link",
							"origin",
							"resource",
						},
						instance: links.NewOriginResourceForTests(
							*pHash,
						),
					},
				},
			},
			{
				path: []string{
					"library",
					"link",
					"origin",
					"resource",
				},
				instance: links.NewOriginValueWithOriginForTests(
					links.NewOriginForTests(
						links.NewOriginResourceForTests(
							*pHash,
						),
						links.NewOperatorWithAndForTests(),
						links.NewOriginValueWithResourceForTests(
							links.NewOriginResourceForTests(
								*pHash,
							),
						),
					),
				),
				dependencies: []testInstance{
					{
						path: []string{
							"library",
							"link",
							"origin",
						},
						instance: links.NewOriginForTests(
							links.NewOriginResourceForTests(
								*pHash,
							),
							links.NewOperatorWithAndForTests(),
							links.NewOriginValueWithResourceForTests(
								links.NewOriginResourceForTests(
									*pHash,
								),
							),
						),
						dependencies: []testInstance{
							{
								path: []string{
									"library",
									"link",
									"origin",
									"resource",
								},
								instance: links.NewOriginResourceForTests(
									*pHash,
								),
							},
							{
								path: []string{
									"library",
									"link",
									"origin",
									"operator",
								},
								instance: links.NewOperatorWithAndForTests(),
							},
							{
								path: []string{
									"library",
									"link",
									"origin",
									"resource",
								},
								instance: links.NewOriginValueWithResourceForTests(
									links.NewOriginResourceForTests(
										*pHash,
									),
								),
								dependencies: []testInstance{
									{
										path: []string{
											"library",
											"link",
											"origin",
											"resource",
										},
										instance: links.NewOriginResourceForTests(
											*pHash,
										),
									},
								},
							},
						},
					},
				},
			},
		},
		"library_link_origin_resource": {
			{
				path: []string{
					"library",
					"link",
					"origin",
					"resource",
				},
				instance: links.NewOriginResourceForTests(
					*pHash,
				),
			},
			{
				path: []string{
					"library",
					"link",
					"origin",
					"resource",
				},
				instance: links.NewOriginResourceWithIsMandatoryForTests(
					*pHash,
				),
			},
		},
		"library_link_origin_operator": {
			{
				path: []string{
					"library",
					"link",
					"origin",
					"operator",
				},
				instance: links.NewOperatorWithAndForTests(),
			},
			{
				path: []string{
					"library",
					"link",
					"origin",
					"operator",
				},
				instance: links.NewOperatorWithOrForTests(),
			},
			{
				path: []string{
					"library",
					"link",
					"origin",
					"operator",
				},
				instance: links.NewOperatorWithXOrForTests(),
			},
		},*/
		"library_layer": {
			{
				path: []string{
					"library",
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
				dependencies: []testInstance{
					{
						path: []string{
							"library",
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
			},
		},
		"library_layer_output": {
			{
				path: []string{
					"library",
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
					"library",
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
		"library_layer_output_kind": {
			{
				path: []string{
					"library",
					"layer",
					"output",
					"kind",
				},
				instance: layers.NewKindWithContinueForTests(),
			},
			{
				path: []string{
					"library",
					"layer",
					"output",
					"kind",
				},
				instance: layers.NewKindWithPromptForTests(),
			},
		},
		"library_layer_instruction": {
			{
				path: []string{
					"library",
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
		"library_layer_instruction_assignment": {
			{
				path: []string{
					"library",
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
		"library_layer_instruction_assignment_assignable": {
			{
				path: []string{
					"library",
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
		"library_layer_instruction_assignment_assignable_bytes": {
			{
				path: []string{
					"library",
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
					"library",
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
					"library",
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

			// if there is dependencies, insert them:
			if oneInstance.dependencies != nil && len(oneInstance.dependencies) > 0 {
				for _, oneDependency := range oneInstance.dependencies {
					err = service.Insert(oneDependency.instance, oneDependency.path)
					if err != nil {
						t.Errorf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, err.Error())
						return
					}
				}
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
