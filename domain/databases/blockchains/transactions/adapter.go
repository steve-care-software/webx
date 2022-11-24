package transactions

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type adapter struct {
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToContent converts transaction to bytes
func (app *adapter) ToContent(ins Transaction) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	assetBytes := ins.Asset().Bytes()
	proofBytes := ins.Proof().Bytes()

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, assetBytes...)
	output = append(output, proofBytes...)
	return output, nil
}

// ToTransaction converts bytes to transaction
func (app *adapter) ToTransaction(content []byte) (Transaction, error) {
	contentLength := len(content)
	if contentLength < minTransactionSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the chain size of the Blockchain instance, %d provided", minTransactionSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	assetDelimiter := hash.Size * 2
	pAsset, err := app.hashAdapter.FromBytes(content[hash.Size:assetDelimiter])
	if err != nil {
		return nil, err
	}

	proof := big.NewInt(0).SetBytes(content[assetDelimiter:])
	return app.builder.Create().
		WithHash(*pHash).
		WithAsset(*pAsset).
		WithProof(*proof).
		Now()
}
