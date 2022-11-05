package transactions

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type transactions struct {
	hash   hash.Hash
	list   []Transaction
	mp     map[string]Transaction
	pScore *big.Int
}

func createTransactions(
	hash hash.Hash,
	list []Transaction,
	mp map[string]Transaction,
	pScore *big.Int,
) Transactions {
	out := transactions{
		hash:   hash,
		list:   list,
		mp:     mp,
		pScore: pScore,
	}

	return &out
}

// Hash returns the hash
func (obj *transactions) Hash() hash.Hash {
	return obj.hash
}

// List returns the transactions list
func (obj *transactions) List() []Transaction {
	return obj.list
}

// Fetch fetches a transaction by hash
func (obj *transactions) Fetch(hash hash.Hash) (Transaction, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the Transaction (hash: %s) could not be found in the Transactions", keyname)
	return nil, errors.New(str)
}

// Score returns the score
func (obj *transactions) Score() *big.Int {
	return obj.pScore
}
