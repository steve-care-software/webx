package links

import (
	"testing"

	"github.com/steve-care-software/identity/domain/hash"
)

func TestLinks_withList_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	list := []Link{
		NewLinkForTests(
			NewOriginForTests(
				NewOriginResourceForTests(*pFirstLayer),
				NewOperatorWithAndForTests(),
				NewOriginValueWithResourceForTests(
					NewOriginResourceForTests(*pSecondLayer),
				),
			),
			NewElementsForTests([]Element{
				NewElementForTests(*pLayer),
			}),
		),
	}

	ins := NewLinksForTests(list)
	retList := ins.List()
	if len(list) != len(retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestLinks_withEmptyList_returnsError(t *testing.T) {
	list := []Link{}
	_, err := NewBuilder().Create().WithList(list).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}

func TestLinks_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
