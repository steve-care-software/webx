package pointers

import (
	"errors"

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

// NextIndex returns the next index
func (obj *pointers) NextIndex() uint64 {
	var delimiterWithBiggestIndex delimiters.Delimiter
	for _, onePointer := range obj.list {
		delimiter := onePointer.Delimiter()
		if delimiterWithBiggestIndex == nil {
			delimiterWithBiggestIndex = delimiter
			continue
		}

		index := delimiter.Index()
		if index > delimiterWithBiggestIndex.Index() {
			delimiterWithBiggestIndex = delimiter
		}
	}

	return delimiterWithBiggestIndex.Index() + delimiterWithBiggestIndex.Length()
}
