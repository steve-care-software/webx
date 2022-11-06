package entities

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type reference struct {
	identifier Identifier
	trx        hash.Hash
	block      hash.Hash
	chain      hash.Hash
}

func createReference(
	identifier Identifier,
	trx hash.Hash,
	block hash.Hash,
	chain hash.Hash,
) Reference {
	out := reference{
		identifier: identifier,
		trx:        trx,
		block:      block,
		chain:      chain,
	}

	return &out
}

// Identifier returns the identifier
func (obj *reference) Identifier() Identifier {
	return obj.identifier
}

// Transaction returns the transaction hash
func (obj *reference) Transaction() hash.Hash {
	return obj.trx
}

// Block returns the block hash
func (obj *reference) Block() hash.Hash {
	return obj.block
}

// Chain returns the chain hash
func (obj *reference) Chain() hash.Hash {
	return obj.chain
}
