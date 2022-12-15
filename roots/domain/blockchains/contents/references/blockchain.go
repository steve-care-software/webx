package references

type blockchain struct {
	chain  BlockchainKey
	blocks BlockchainKeys
	trx    BlockchainKeys
}

func createBlockchain(
	chain BlockchainKey,
	blocks BlockchainKeys,
	trx BlockchainKeys,
) Blockchain {
	out := blockchain{
		chain:  chain,
		blocks: blocks,
		trx:    trx,
	}

	return &out
}

// Chain returns the chain
func (obj *blockchain) Chain() BlockchainKey {
	return obj.chain
}

// Blocks returns the blocks
func (obj *blockchain) Blocks() BlockchainKeys {
	return obj.blocks
}

// Transactions returns the transactions
func (obj *blockchain) Transactions() BlockchainKeys {
	return obj.trx
}
