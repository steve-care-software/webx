package sqllites

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
	"github.com/steve-care-software/datastencil/domain/orms"
)

type typeDependency struct {
	keyname string
	index   uint
}

type testInstance struct {
	path         []string
	instance     orms.Instance
	dependencies []typeDependency
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

	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// build resources:
	instances := map[string][]testInstance{
		"library": {
			{
				path: []string{
					"library",
				},
				instance: libraries.NewLibraryForTests(
					layers.NewLayersForTests([]layers.Layer{
						layers.NewLayerForTests(
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
					}),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_layer",
						index:   0,
					},
				},
			},
			{
				path: []string{
					"library",
				},
				instance: libraries.NewLibraryWithLinksForTests(
					layers.NewLayersForTests([]layers.Layer{
						layers.NewLayerForTests(
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
					}),
					links.NewLinksForTests([]links.Link{
						links.NewLinkForTests(
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
							links.NewElementsForTests([]links.Element{
								links.NewElementForTests(
									*pHash,
								),
							}),
						),
					}),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_layer",
						index:   0,
					},
					{
						keyname: "library_link",
						index:   0,
					},
				},
			},
		},
		"library_link": {
			{
				path: []string{
					"library",
					"link",
				},
				instance: links.NewLinkForTests(
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
					links.NewElementsForTests([]links.Element{
						links.NewElementForTests(
							*pHash,
						),
					}),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_origin",
						index:   0,
					},
					{
						keyname: "library_link_element",
						index:   0,
					},
				},
			},
		},
		"library_link_element": {
			{
				path: []string{
					"library",
					"link",
					"element",
				},
				instance: links.NewElementForTests(
					*pHash,
				),
			},
			{
				path: []string{
					"library",
					"link",
					"element",
				},
				instance: links.NewElementWithConditionForTests(
					*pHash,
					links.NewConditionForTests(
						links.NewConditionResourceForTests(22),
					),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_element_condition",
						index:   0,
					},
				},
			},
		},
		"library_link_element_condition": {
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
				},
				instance: links.NewConditionForTests(
					links.NewConditionResourceForTests(22),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_element_condition_resource",
						index:   0,
					},
				},
			},
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
				},
				instance: links.NewConditionWithNextForTests(
					links.NewConditionResourceForTests(22),
					links.NewConditionValueWithResourceForTests(
						links.NewConditionResourceForTests(22),
					),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_element_condition_resource",
						index:   0,
					},
					{
						keyname: "library_link_element_condition_value",
						index:   0,
					},
				},
			},
		},
		"library_link_element_condition_value": {
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
					"value",
				},
				instance: links.NewConditionValueWithResourceForTests(
					links.NewConditionResourceForTests(22),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_element_condition_resource",
						index:   0,
					},
				},
			},
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
					"value",
				},
				instance: links.NewConditionValueWithConditionForTests(
					links.NewConditionForTests(
						links.NewConditionResourceForTests(22),
					),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_element_condition",
						index:   0,
					},
				},
			},
		},
		"library_link_element_condition_resource": {
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				},
				instance: links.NewConditionResourceForTests(22),
			},
			{
				path: []string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				},
				instance: links.NewConditionResourceWithIsRaisedInLayerForTests(54),
			},
		},
		"library_link_origin": {
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
				dependencies: []typeDependency{
					{
						keyname: "library_link_origin_resource",
						index:   0,
					},
					{
						keyname: "library_link_origin_operator",
						index:   0,
					},
					{
						keyname: "library_link_origin_value",
						index:   0,
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
					"value",
				},
				instance: links.NewOriginValueWithResourceForTests(
					links.NewOriginResourceForTests(
						*pHash,
					),
				),
				dependencies: []typeDependency{
					{
						keyname: "library_link_origin_resource",
						index:   0,
					},
				},
			},
			{
				path: []string{
					"library",
					"link",
					"origin",
					"value",
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
				dependencies: []typeDependency{
					{
						keyname: "library_link_origin",
						index:   0,
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
		},
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_instruction",
						index:   0,
					},
					{
						keyname: "library_layer_output",
						index:   0,
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_output_kind",
						index:   0,
					},
				},
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_output_kind",
						index:   0,
					},
				},
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_instruction_assignment",
						index:   0,
					},
				},
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_instruction_assignment_assignable",
						index:   0,
					},
				},
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
				dependencies: []typeDependency{
					{
						keyname: "library_layer_instruction_assignment_assignable_bytes",
						index:   0,
					},
				},
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

			// insert instance:
			err = insertInstance(name, idx, oneInstance, instances, service, false)
			if err != nil {
				t.Errorf(err.Error())
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

func insertInstance(
	name string,
	idx int,
	currentInstance testInstance,
	allInstances map[string][]testInstance,
	service orms.Service,
	skipIfFails bool,
) error {
	// if there is dependencies, insert them:
	if currentInstance.dependencies != nil && len(currentInstance.dependencies) > 0 {
		err := insertDependencies(name, idx, currentInstance, allInstances, service)
		if err != nil {
			str := fmt.Sprintf("section: %s: index: %d, there was an error while inserting the dependencies: %s", name, idx, fmt.Sprintln(err.Error()))
			return errors.New(str)
		}
	}

	// insert instance:
	err := service.Insert(currentInstance.instance, currentInstance.path)
	if err != nil {
		if skipIfFails {
			return nil
		}

		str := fmt.Sprintf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, fmt.Sprintln(err.Error()))
		return errors.New(fmt.Sprintln(str))
	}

	return nil
}

func insertDependencies(
	name string,
	idx int,
	currentInstance testInstance,
	allInstances map[string][]testInstance,
	service orms.Service,
) error {
	for _, oneDependency := range currentInstance.dependencies {
		if dependencyList, ok := allInstances[oneDependency.keyname]; ok {
			if dependencyIns := dependencyList[oneDependency.index]; ok {
				err := insertInstance(name, idx, dependencyIns, allInstances, service, true)
				if err != nil {
					str := fmt.Sprintf("section: %s: index: %d, the error was expected to be nil, error returned: %s", name, idx, fmt.Sprintln(err.Error()))
					return errors.New(str)
				}

				continue
			}

			str := fmt.Sprintf("the dependency (keyname: %s, index: %d) is undeclared in the instances list, used in at keyname: %s, index: %d", oneDependency.keyname, oneDependency.index, name, idx)
			return errors.New(fmt.Sprintln(str))
		}
	}

	return nil
}
