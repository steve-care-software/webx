package links

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
)

func TestLink_Success(t *testing.T) {
	elements := elements.NewElementsForTests([]elements.Element{
		elements.NewElementForTests(
			[]string{"this", "is", "my", "layer", "path"},
		),
	})

	link := NewLinkForTests(elements)

	retElements := link.Elements()
	if !reflect.DeepEqual(elements, retElements) {
		t.Errorf("the elements is invalid")
		return
	}
}

func TestLink_withoutElements_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
