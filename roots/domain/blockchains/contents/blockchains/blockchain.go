package blockchains

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type blockchain struct {
	reference hash.Hash
	head      hash.Hash
}

func createBlockchain(
	reference hash.Hash,
	head hash.Hash,
) Blockchain {
	out := blockchain{
		reference: reference,
		head:      head,
	}

	return &out
}

// Reference returns the reference
func (obj *blockchain) Reference() hash.Hash {
	return obj.reference
}

// Head returns the head block hash
func (obj *blockchain) Head() hash.Hash {
	return obj.head
}
