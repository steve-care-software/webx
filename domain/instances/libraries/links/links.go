package links

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type links struct {
	hash hash.Hash
	mp   map[string]Link
	list []Link
}

func createLinks(
	hash hash.Hash,
	mp map[string]Link,
	list []Link,
) Links {
	out := links{
		hash: hash,
		mp:   mp,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *links) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *links) List() []Link {
	return obj.list
}

// Fetch fetches a link by hash
func (obj *links) Fetch(hash hash.Hash) (Link, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the requested Link (hash: %s) does not exists", keyname)
	return nil, errors.New(str)
}

// FetchByExecutedLayers fetches link by executed layers
func (obj *links) FetchByExecutedLayers(layerHashes []hash.Hash) (Link, error) {
	return nil, nil
}
