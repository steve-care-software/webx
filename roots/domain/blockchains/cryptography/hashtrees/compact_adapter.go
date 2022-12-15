package hashtrees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type compactAdapter struct {
	hashAdapter   hash.Adapter
	leavesAdapter LeavesAdapter
	builder       CompactBuilder
}

func createCompactAdapter(
	hashAdapter hash.Adapter,
	leavesAdapter LeavesAdapter,
	builder CompactBuilder,
) CompactAdapter {
	out := compactAdapter{
		hashAdapter:   hashAdapter,
		leavesAdapter: leavesAdapter,
		builder:       builder,
	}

	return &out
}

// ToContent converts a Compact to bytes
func (app *compactAdapter) ToContent(ins Compact) ([]byte, error) {
	headBytes := ins.Head().Bytes()
	leavesBytes, err := app.leavesAdapter.ToContent(ins.Leaves())
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, headBytes...)
	output = append(output, leavesBytes...)
	return output, nil
}

// ToCompact converts bytes to Compact
func (app *compactAdapter) ToCompact(content []byte) (Compact, error) {
	contentLength := len(content)
	if contentLength < minCompactSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Compact instance, %d provided", minCompactSize, contentLength)
		return nil, errors.New(str)
	}

	pHead, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	leaves, err := app.leavesAdapter.ToLeaves(content[hash.Size:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHead(*pHead).
		WithLeaves(leaves).
		Now()
}
