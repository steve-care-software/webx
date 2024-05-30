package links

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
)

func TestLinks_withList_Success(t *testing.T) {
	list := []Link{
		NewLinkForTests(
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementForTests(
					[]string{"this", "is", "my", "layer", "path"},
				),
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
