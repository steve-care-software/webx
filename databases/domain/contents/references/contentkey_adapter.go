package references

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
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

	kindBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(kindBytes, uint64(ins.Kind()))

	contentBytes, err := app.pointerAdapter.ToContent(ins.Content())
	if err != nil {
		return nil, err
	}

	commitBytes := ins.Commit().Bytes()

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, kindBytes...)
	output = append(output, contentBytes...)
	output = append(output, commitBytes...)
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

	kindDelimiter := hash.Size + 8
	kind := binary.LittleEndian.Uint64(content[hash.Size:kindDelimiter])

	contentDelimiter := kindDelimiter + pointerSize
	pointerContent, err := app.pointerAdapter.ToPointer(content[kindDelimiter:contentDelimiter])
	if err != nil {
		return nil, err
	}

	commitDelimiter := contentDelimiter + hash.Size
	pCommitHash, err := app.hashAdapter.FromBytes(content[contentDelimiter:commitDelimiter])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHash(*pHash).
		WithKind(uint(kind)).
		WithContent(pointerContent).
		WithCommit(*pCommitHash).
		Now()
}
