package pointers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type pointers struct {
	list []Pointer
	mp   map[string]Pointer
}

func createPointers(
	list []Pointer,
	mp map[string]Pointer,
) Pointers {
	out := pointers{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of pointers
func (obj *pointers) List() []Pointer {
	return obj.list
}

// Retrieve retrieves a pointer by hash
func (obj *pointers) Retrieve(hash hash.Hash) (Pointer, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Pointer that matches the requested hash: %s", keyname)
	return nil, errors.New(str)
}

// RetrieveAll retrieves all pointers related to the hashes
func (obj *pointers) RetrieveAll(hashes []hash.Hash) ([]Pointer, error) {
	output := []Pointer{}
	for _, oneHash := range hashes {
		ins, err := obj.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return output, nil
}
