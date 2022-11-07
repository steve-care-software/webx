package blockchains

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type blockchain struct {
	entity    entities.Entity
	reference hash.Hash
	head      entities.Identifier
}

func createBlockchain(
	entity entities.Entity,
	reference hash.Hash,
	head entities.Identifier,
) Blockchain {
	out := blockchain{
		entity:    entity,
		reference: reference,
		head:      head,
	}

	return &out
}

// Entity returns the entity
func (obj *blockchain) Entity() entities.Entity {
	return obj.entity
}

// Reference returns the reference
func (obj *blockchain) Reference() hash.Hash {
	return obj.reference
}

// Head returns the head block
func (obj *blockchain) Head() entities.Identifier {
	return obj.head
}
