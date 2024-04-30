package links

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	condition_resources "github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

func TestAdapter_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondHash, err := hash.NewAdapter().FromBytes([]byte("this is another layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := links.NewLinksForTests([]links.Link{
		links.NewLinkForTests(
			origins.NewOriginForTests(
				resources.NewResourceForTests(*pHash),
				operators.NewOperatorWithAndForTests(),
				origins.NewValueWithResourceForTests(
					resources.NewResourceWithIsMandatoryForTests(*pSecondHash),
				),
			),
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementForTests(
					layers.NewLayerForTests(
						instructions.NewInstructionsForTests([]instructions.Instruction{
							instructions.NewInstructionWithAssignmentForTests(
								assignments.NewAssignmentForTests(
									"anotherName",
									assignables.NewAssignableWithBytesForTests(
										bytes_domain.NewBytesWithHashBytesForTests(
											"anotherInput",
										),
									),
								),
							),
							instructions.NewInstructionWithRaiseErrorForTests(22),
							instructions.NewInstructionWithStopForTests(),
						}),
						outputs.NewOutputForTests(
							"myVariable",
							kinds.NewKindWithContinueForTests(),
						),
						"myInput",
					),
				),
			}),
		),
		links.NewLinkWithReferecesForTests(
			origins.NewOriginForTests(
				resources.NewResourceForTests(*pHash),
				operators.NewOperatorWithAndForTests(),
				origins.NewValueWithResourceForTests(
					resources.NewResourceWithIsMandatoryForTests(*pSecondHash),
				),
			),
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementWithConditionForTests(
					layers.NewLayerForTests(
						instructions.NewInstructionsForTests([]instructions.Instruction{
							instructions.NewInstructionWithAssignmentForTests(
								assignments.NewAssignmentForTests(
									"anotherName",
									assignables.NewAssignableWithBytesForTests(
										bytes_domain.NewBytesWithHashBytesForTests(
											"anotherInput",
										),
									),
								),
							),
							instructions.NewInstructionWithRaiseErrorForTests(22),
							instructions.NewInstructionWithStopForTests(),
						}),
						outputs.NewOutputForTests(
							"myVariable",
							kinds.NewKindWithContinueForTests(),
						),
						"myInput",
					),
					conditions.NewConditionForTests(
						condition_resources.NewResourceForTests(uint(45)),
					),
				),
			}),
			references.NewReferencesForTests([]references.Reference{
				references.NewReferenceForTests("myVariable", *pHash),
			}),
		),
	})

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
