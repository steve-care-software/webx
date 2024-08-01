package pointers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
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

// Fetch fetches the pointer that matches the delimiter
func (obj *pointers) Fetch(delimiter delimiters.Delimiter) (Pointer, error) {
	for _, onePointer := range obj.list {
		currentDelimiter := onePointer.Delimiter()
		if delimiter.Index() == currentDelimiter.Index() && delimiter.Length() == currentDelimiter.Length() {
			return onePointer, nil
		}
	}

	return nil, errors.New("there is no match for the provided delimiter")
}

// Subset returns a subset of the pointer's list
func (obj *pointers) Subset(index uint64, length uint64) ([]Pointer, error) {
	list := []Pointer{}
	for _, onePointer := range obj.list {
		if onePointer.IsDeleted() {
			continue
		}

		list = append(list, onePointer)
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

// Search searches for the delimiters that are present in the search
func (obj *pointers) Search(index uint64, length uint64) ([]delimiters.Delimiter, error) {
	toIndex := index + length
	list := []delimiters.Delimiter{}
	for _, onePointer := range obj.list {
		delimiter := onePointer.Delimiter()
		delIndex := delimiter.Index()
		if index <= delIndex && index < toIndex {
			list = append(list, delimiter)
			continue
		}

		break
	}

	if len(list) <= 0 {
		str := fmt.Sprintf("there is no pointer list of the provided index (%d) and length (%d)", index, length)
		return nil, errors.New(str)
	}

	return list, nil
}
