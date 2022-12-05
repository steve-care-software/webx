package blockchains

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/blockchains/blocks"
)

type blockchain struct {
	reference hash.Hash
	head      blocks.Block
}

func createBlockchain(
	reference hash.Hash,
	head blocks.Block,
) Blockchain {
	return createBlockchainInternally(reference, head)
}

func createBlockchainInternally(
	reference hash.Hash,
	head blocks.Block,
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

// Head returns the head block
func (obj *blockchain) Head() blocks.Block {
	return obj.head
}
