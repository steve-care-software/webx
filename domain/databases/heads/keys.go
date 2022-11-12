package heads

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type keys struct {
	mp   map[string]Key
	list []Key
}

func createKeys(
	mp map[string]Key,
	list []Key,
) Keys {
	out := keys{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the keys
func (obj *keys) List() []Key {
	return obj.list
}

// Fetch fetches a key by hash
func (obj *keys) Fetch(hash hash.Hash) (Key, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the key (hash: %s) is invalid", keyname)
	return nil, errors.New(str)
}
