package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
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

// Fetch fetches a contentKey by hash
func (obj *contentKeys) Fetch(hash hash.Hash) (ContentKey, error) {
	contentKeyname := hash.String()
	if ins, ok := obj.mp[contentKeyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the contentKey (hash: %s) is invalid", contentKeyname)
	return nil, errors.New(str)
}
