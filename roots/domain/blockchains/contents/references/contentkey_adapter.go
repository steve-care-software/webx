package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type contentKeyAdapter struct {
	hashAdapter    hash.Adapter
	pointerAdapter PointerAdapter
	builder        ContentKeyBuilder
}

func createContentKeyAdapter(
	hashAdapter hash.Adapter,
	pointerAdapter PointerAdapter,
	builder ContentKeyBuilder,
) ContentKeyAdapter {
	out := contentKeyAdapter{
		hashAdapter:    hashAdapter,
		pointerAdapter: pointerAdapter,
		builder:        builder,
	}

	return &out
}

// ToContent converts ContentKey to bytes
func (app *contentKeyAdapter) ToContent(ins ContentKey) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	kind := ins.Kind()
	contentBytes, err := app.pointerAdapter.ToContent(ins.Content())
	if err != nil {
		return nil, err
	}

	trxBytes := ins.Transaction().Bytes()
	createdOnBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(createdOnBytes, uint64(ins.CreatedOn().UnixNano()))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, kind)
	output = append(output, contentBytes...)
	output = append(output, trxBytes...)
	output = append(output, createdOnBytes...)
	return output, nil
}

// ToContentKey converts bytes to ContentKey instance
func (app *contentKeyAdapter) ToContentKey(content []byte) (ContentKey, error) {
	if len(content) != contentKeySize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert to a Pointer instance, %d provided", contentKeySize, len(content))
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	kindDelimiter := hash.Size + 1
	kind := content[hash.Size:kindDelimiter][0]

	contentDelimiter := kindDelimiter + pointerSize
	pointerContent, err := app.pointerAdapter.ToPointer(content[kindDelimiter:contentDelimiter])
	if err != nil {
		return nil, err
	}

	trxDelimiter := contentDelimiter + hash.Size
	pTrxHash, err := app.hashAdapter.FromBytes(content[contentDelimiter:trxDelimiter])
	if err != nil {
		return nil, err
	}

	createdOnUnixNano := binary.LittleEndian.Uint64(content[trxDelimiter : trxDelimiter+8])
	createdOn := time.Unix(0, int64(createdOnUnixNano)).UTC()

	return app.builder.Create().
		WithHash(*pHash).
		WithKind(kind).
		WithContent(pointerContent).
		WithTransaction(*pTrxHash).
		CreatedOn(createdOn).
		Now()
}
