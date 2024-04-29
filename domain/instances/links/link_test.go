package links

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestLink_Success(t *testing.T) {
	elements := elements.NewElementsForTests([]elements.Element{
		elements.NewElementForTests(
			layers.NewLayerForTests(
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
				outputs.NewOutputForTests(
					"myVariable",
					kinds.NewKindWithContinueForTests(),
				),
				"myInput",
			),
		),
	})

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	origin := origins.NewOriginForTests(
		resources.NewResourceForTests(*pFirstLayer),
		operators.NewOperatorWithAndForTests(),
		origins.NewValueWithResourceForTests(
			resources.NewResourceForTests(*pSecondLayer),
		),
	)

	link := NewLinkForTests(origin, elements)

	retOrigin := link.Origin()
	if !reflect.DeepEqual(origin, retOrigin) {
		t.Errorf("the origin is invalid")
		return
	}

	retElements := link.Elements()
	if !reflect.DeepEqual(elements, retElements) {
		t.Errorf("the elements is invalid")
		return
	}
}

func TestLink_withoutElements_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	origin := origins.NewOriginForTests(
		resources.NewResourceForTests(*pFirstLayer),
		operators.NewOperatorWithAndForTests(),
		origins.NewValueWithResourceForTests(
			resources.NewResourceForTests(*pSecondLayer),
		),
	)

	_, err := NewLinkBuilder().Create().WithOrigin(origin).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLink_withoutOrigin_returnsError(t *testing.T) {
	elements := elements.NewElementsForTests([]elements.Element{
		elements.NewElementForTests(
			layers.NewLayerForTests(
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
				outputs.NewOutputForTests(
					"myVariable",
					kinds.NewKindWithContinueForTests(),
				),
				"myInput",
			),
		),
	})

	_, err := NewLinkBuilder().Create().WithElements(elements).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
