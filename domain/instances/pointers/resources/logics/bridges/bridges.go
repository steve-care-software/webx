package bridges

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type bridges struct {
	hash hash.Hash
	list []Bridge
	mp   map[string]Bridge
}

func createBridges(
	hash hash.Hash,
	list []Bridge,
	mp map[string]Bridge,
) Bridges {
	out := bridges{
		hash: hash,
		list: list,
		mp:   mp,
	}

	return &out
}

// Hash returns the hash
func (obj *bridges) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *bridges) List() []Bridge {
	return obj.list
}

// Fetch fetches a bridge by layer path
func (obj *bridges) Fetch(path []string) (Bridge, error) {
	keyname := filepath.Join(path...)
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Bridge at path: %s", keyname)
	return nil, errors.New(str)
}
