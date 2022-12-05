package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type blockchainKeyAdapter struct {
	hashAdapter    hash.Adapter
	pointerAdapter PointerAdapter
	builder        BlockchainKeyBuilder
}

func createBlockchainKeyAdapter(
	hashAdapter hash.Adapter,
	pointerAdapter PointerAdapter,
	builder BlockchainKeyBuilder,
) BlockchainKeyAdapter {
	out := blockchainKeyAdapter{
		hashAdapter:    hashAdapter,
		pointerAdapter: pointerAdapter,
		builder:        builder,
	}

	return &out
}

// ToContent converts a BlockchainKey to bytes
func (app *blockchainKeyAdapter) ToContent(ins BlockchainKey) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	contentBytes, err := app.pointerAdapter.ToContent(ins.Content())
	if err != nil {
		return nil, err
	}

	createdOnBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(createdOnBytes, uint64(ins.CreatedOn().UnixNano()))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, contentBytes...)
	output = append(output, createdOnBytes...)
	return output, nil
}

// ToBlockchainKey converts bytes to a BlockchainKey instance
func (app *blockchainKeyAdapter) ToBlockchainKey(content []byte) (BlockchainKey, error) {
	if len(content) != blockchainKeySize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert to a Pointer instance, %d provided", blockchainKeySize, len(content))
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	contentDelimiter := hash.Size + pointerSize
	pointerContent, err := app.pointerAdapter.ToPointer(content[hash.Size:contentDelimiter])
	if err != nil {
		return nil, err
	}

	createdOnUnixNano := binary.LittleEndian.Uint64(content[contentDelimiter : contentDelimiter+8])
	createdOn := time.Unix(0, int64(createdOnUnixNano)).UTC()

	return app.builder.Create().
		WithHash(*pHash).
		WithContent(pointerContent).
		CreatedOn(createdOn).
		Now()
}
