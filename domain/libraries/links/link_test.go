package links

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/resources"
)

func TestLink_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	elements := NewElementsForTests([]Element{
		NewElementForTests(*pLayer),
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
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	elements := NewElementsForTests([]Element{
		NewElementForTests(*pLayer),
	})

	_, err := NewLinkBuilder().Create().WithElements(elements).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
