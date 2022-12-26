package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type contentKeys struct {
	mp   map[string]ContentKey
	list []ContentKey
}

func createContentKeys(
	mp map[string]ContentKey,
	list []ContentKey,
) ContentKeys {
	out := contentKeys{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the contentKeys
func (obj *contentKeys) List() []ContentKey {
	return obj.list
}

// Next returns the next beginning index for a pointer
func (obj *contentKeys) Next() int64 {
	biggest := int64(0)
	for _, oneContentKey := range obj.list {
		pointer := oneContentKey.Content()
		next := int64(pointer.From() + pointer.Length())
		if biggest < next {
			biggest = next
		}
	}

	return biggest
}

// ListByKind returns the list by kind
func (obj *contentKeys) ListByKind(kind uint) []ContentKey {
	output := []ContentKey{}
	for _, oneContentKey := range obj.list {
		if oneContentKey.Kind() != kind {
			continue
		}

		output = append(output, oneContentKey)
	}

	return output
}

// Fetch fetches a contentKey by hash
func (obj *contentKeys) Fetch(hash hash.Hash) (ContentKey, error) {
	contentKeyname := hash.String()
	if ins, ok := obj.mp[contentKeyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the contentKey (hash: %s) is invalid", contentKeyname)
	return nil, errors.New(str)
}
