package elements

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
)

func TestElements_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	elements := NewElementsForTests([]Element{
		NewElementForTests(*pLayer),
	})

	retList := elements.List()
	if len(retList) != 1 {
		t.Errorf("the list was expected to contain 1 element")
		return
	}
}

func TestElements_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Element{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElements_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
