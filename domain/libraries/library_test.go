package libraries

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/resources"
)

func TestLibrary_Success(t *testing.T) {
	layers := layers.NewLayersForTests([]layers.Layer{
		layers.NewLayerForTests(
			layers.NewInstructionsForTests([]layers.Instruction{
				layers.NewInstructionWithStopForTests(),
			}),
			layers.NewOutputForTests(
				"myVariable",
				layers.NewKindWithContinueForTests(),
			),
			"someInput",
		),
	})

	ins := NewLibraryForTests(layers)
	retLayers := ins.Layers()
	if !reflect.DeepEqual(layers, retLayers) {
		t.Errorf("the returned layers is invalid")
		return
	}

	if ins.HasLinks() {
		t.Errorf("the library was expected to NOT contain links")
		return
	}
}

func TestLibrary_withLinks_Success(t *testing.T) {
	layers := layers.NewLayersForTests([]layers.Layer{
		layers.NewLayerForTests(
			layers.NewInstructionsForTests([]layers.Instruction{
				layers.NewInstructionWithStopForTests(),
			}),
			layers.NewOutputForTests(
				"myVariable",
				layers.NewKindWithContinueForTests(),
			),
			"someInput",
		),
	})

	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	links := links.NewLinksForTests([]links.Link{
		links.NewLinkForTests(
			origins.NewOriginForTests(
				resources.NewResourceForTests(*pFirstLayer),
				operators.NewOperatorWithAndForTests(),
				origins.NewValueWithResourceForTests(
					resources.NewResourceForTests(*pSecondLayer),
				),
			),
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementForTests(*pLayer),
			}),
		),
	})

	ins := NewLibraryWithLinksForTests(layers, links)
	retLayers := ins.Layers()
	if !reflect.DeepEqual(layers, retLayers) {
		t.Errorf("the returned layers is invalid")
		return
	}

	if !ins.HasLinks() {
		t.Errorf("the library was expected to contain links")
		return
	}

	retLinks := ins.Links()
	if !reflect.DeepEqual(links, retLinks) {
		t.Errorf("the returned links is invalid")
		return
	}
}

func TestLibrary_withoutLayers_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
