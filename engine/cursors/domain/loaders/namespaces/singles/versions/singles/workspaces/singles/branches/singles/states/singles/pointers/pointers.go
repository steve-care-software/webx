package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type pointersIns struct {
	list []Pointer
}

func createPointers(
	list []Pointer,
) Pointers {
	out := pointersIns{
		list: list,
	}

	return &out
}

// List returns the list of pointers
func (obj *pointersIns) List() []Pointer {
	return obj.list
}

// NextIndex returns the next index
func (obj *pointersIns) NextIndex() uint64 {
	var delimiterWithBiggestIndex delimiters.Delimiter
	for _, onePointer := range obj.list {
		delimiter := onePointer.Storage().Delimiter()
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

// FindIndex finds the index of the pointer associated with the provided delimiter
func (obj *pointersIns) FindIndex(delimiter delimiters.Delimiter) (*uint, error) {
	for idx, onePointer := range obj.list {
		currentDelimiter := onePointer.Storage().Delimiter()
		if delimiter.Index() == currentDelimiter.Index() && delimiter.Length() == currentDelimiter.Length() {
			pIndex := uint(idx)
			return &pIndex, nil
		}
	}

	return nil, errors.New("there is no match for the provided delimiter")
}
