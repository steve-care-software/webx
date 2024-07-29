package pointers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

type pointers struct {
	list []Pointer
}

func createPointers(
	list []Pointer,
) Pointers {
	out := pointers{
		list: list,
	}

	return &out
}

// List returns the list of pointers
func (obj *pointers) List() []Pointer {
	return obj.list
}

// Fetch fetches the retrievals
func (obj *pointers) Fetch(index uint64, length uint64) ([]retrievals.Retrieval, error) {
	list := []retrievals.Retrieval{}
	for _, onePointer := range obj.list {
		if onePointer.IsDeleted() {
			continue
		}
	}

	listLength := uint64(len(list))
	if index >= listLength {
		str := fmt.Sprintf("the provided index (%d) is larger or equal to the non-deleted retrieval's list (length: %d)", index, listLength)
		return nil, errors.New(str)
	}

	toIndex := index + listLength
	if index >= listLength {
		str := fmt.Sprintf("the provided index (%d) +length (%d) equals %d, which is larger than the non-deleted retrieval's list (length: %d)", index, length, toIndex, listLength)
		return nil, errors.New(str)
	}

	subList := list[index:toIndex]
	return subList, nil
}
