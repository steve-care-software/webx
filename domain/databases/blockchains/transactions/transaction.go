package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type transaction struct {
	entity entities.Entity
	asset  hash.Hash
	proof  big.Int
}

func createTransaction(
	entity entities.Entity,
	asset hash.Hash,
	proof big.Int,
) Transaction {
	out := transaction{
		entity: entity,
		asset:  asset,
		proof:  proof,
	}

	return &out
}

// Entity returns the entity
func (obj *transaction) Entity() entities.Entity {
	return obj.entity
}

// Asset returns the asset
func (obj *transaction) Asset() hash.Hash {
	return obj.asset
}

// Proof returns the proof
func (obj *transaction) Proof() big.Int {
	return obj.proof
}
