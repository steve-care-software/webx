package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type commitAdapter struct {
	hashAdapter    hash.Adapter
	builder        CommitBuilder
	pointerAdapter PointerAdapter
}

func createCommitAdapter(
	hashAdapter hash.Adapter,
	builder CommitBuilder,
	pointerAdapter PointerAdapter,
) CommitAdapter {
	out := commitAdapter{
		hashAdapter:    hashAdapter,
		builder:        builder,
		pointerAdapter: pointerAdapter,
	}

	return &out
}

// ToContent converts a commit instance to content
func (app *commitAdapter) ToContent(ins Commit) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	pointerBytes, err := app.pointerAdapter.ToContent(ins.Pointer())
	if err != nil {
		return nil, err
	}

	createdOnBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(createdOnBytes, uint64(ins.CreatedOn().UnixNano()))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, pointerBytes...)
	output = append(output, createdOnBytes...)
	return output, nil
}

// ToCommit converts content to a commit instance
func (app *commitAdapter) ToCommit(content []byte) (Commit, error) {
	contentLength := len(content)
	if contentLength != commitSize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert to a Commit instance, %d provided", commitSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	pointerDelimiter := hash.Size + pointerSize
	pointer, err := app.pointerAdapter.ToPointer(content[hash.Size:pointerDelimiter])
	if err != nil {
		return nil, err
	}

	createdOnUnixNano := binary.LittleEndian.Uint64(content[pointerDelimiter : pointerDelimiter+8])
	createdOn := time.Unix(0, int64(createdOnUnixNano)).UTC()

	return app.builder.Create().
		WithHash(*pHash).
		WithPointer(pointer).
		CreatedOn(createdOn).
		Now()
}
