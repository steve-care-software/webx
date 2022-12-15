package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type blockchainKeys struct {
	mp   map[string]BlockchainKey
	list []BlockchainKey
}

func createBlockchainKeys(
	mp map[string]BlockchainKey,
	list []BlockchainKey,
) BlockchainKeys {
	out := blockchainKeys{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the blockchainKeys
func (obj *blockchainKeys) List() []BlockchainKey {
	return obj.list
}

// Fetch fetches a blockchainKey by hash
func (obj *blockchainKeys) Fetch(hash hash.Hash) (BlockchainKey, error) {
	blockchainKeyname := hash.String()
	if ins, ok := obj.mp[blockchainKeyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the blockchainKey (hash: %s) is invalid", blockchainKeyname)
	return nil, errors.New(str)
}
