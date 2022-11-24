package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type blockchainAdapter struct {
	blockchainKeyAdapter  BlockchainKeyAdapter
	blockchainKeysAdapter BlockchainKeysAdapter
	builder               BlockchainBuilder
}

func createBlockchainAdapter(
	blockchainKeyAdapter BlockchainKeyAdapter,
	blockchainKeysAdapter BlockchainKeysAdapter,
	builder BlockchainBuilder,
) BlockchainAdapter {
	out := blockchainAdapter{
		blockchainKeyAdapter:  blockchainKeyAdapter,
		blockchainKeysAdapter: blockchainKeysAdapter,
		builder:               builder,
	}

	return &out
}

// ToContent converts blockchain to bytes
func (app *blockchainAdapter) ToContent(ins Blockchain) ([]byte, error) {
	chainBytes, err := app.blockchainKeyAdapter.ToContent(ins.Chain())
	if err != nil {
		return nil, err
	}

	blockBytes, err := app.blockchainKeysAdapter.ToContent(ins.Blocks())
	if err != nil {
		return nil, err
	}

	trxBytes, err := app.blockchainKeysAdapter.ToContent(ins.Transactions())
	if err != nil {
		return nil, err
	}

	chainLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(chainLengthBytes, uint64(len(chainBytes)))

	blockLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(blockLengthBytes, uint64(len(blockBytes)))

	output := []byte{}
	output = append(output, chainLengthBytes...)
	output = append(output, chainBytes...)
	output = append(output, blockLengthBytes...)
	output = append(output, blockBytes...)
	output = append(output, trxBytes...)
	return output, nil
}

// ToBlockchain converts bytes to blockchain instance
func (app *blockchainAdapter) ToBlockchain(content []byte) (Blockchain, error) {
	contentLength := len(content)
	if contentLength < 8 {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the chain size of the Blockchain instance, %d provided", 8, contentLength)
		return nil, errors.New(str)
	}

	chainBytesLength := binary.LittleEndian.Uint64(content[:8])
	chainBytesDelimiter := int(chainBytesLength + 8)
	if contentLength < chainBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the chain size of the Blockchain instance, %d provided", chainBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	chain, err := app.blockchainKeyAdapter.ToBlockchainKey(content[8:chainBytesDelimiter])
	if err != nil {
		return nil, err
	}

	blockLengthDelimiter := int(chainBytesDelimiter + 8)
	if contentLength < blockLengthDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the blocks size of the Blockchain instance, %d provided", blockLengthDelimiter, contentLength)
		return nil, errors.New(str)
	}

	blocksBytesLength := int(binary.LittleEndian.Uint64(content[chainBytesDelimiter:blockLengthDelimiter]))
	blocksDelimiter := blockLengthDelimiter + blocksBytesLength
	blocks, err := app.blockchainKeysAdapter.ToBlockchainKeys(content[blockLengthDelimiter:blocksDelimiter])
	if err != nil {
		return nil, err
	}

	trx, err := app.blockchainKeysAdapter.ToBlockchainKeys(content[blocksDelimiter:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithChain(chain).
		WithBlocks(blocks).
		WithTransactions(trx).
		Now()
}
