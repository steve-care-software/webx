package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type commitAdapter struct {
	hashAdapter     hash.Adapter
	hashTreeAdapter hashtrees.Adapter
	builder         CommitBuilder
}

func createCommitAdapter(
	hashAdapter hash.Adapter,
	hashTreeAdapter hashtrees.Adapter,
	builder CommitBuilder,
) CommitAdapter {
	out := commitAdapter{
		hashAdapter:     hashAdapter,
		hashTreeAdapter: hashTreeAdapter,
		builder:         builder,
	}

	return &out
}

// ToContent converts a Commit instance to content
func (app *commitAdapter) ToContent(ins Commit) ([]byte, error) {
	createdOnBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(createdOnBytes, uint64(ins.CreatedOn().UnixNano()))

	valuesBytes, err := app.hashTreeAdapter.ToContent(ins.Values())
	if err != nil {
		return nil, err
	}

	valuesBytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuesBytesAmount, uint64(len(valuesBytes)))

	output := []byte{}
	output = append(output, createdOnBytes...)
	output = append(output, valuesBytesAmount...)
	output = append(output, valuesBytes...)
	if ins.HasMine() {
		proofBytes := ins.Mine().Proof().Bytes()

		proofBytesLength := make([]byte, 8)
		binary.LittleEndian.PutUint64(proofBytesLength, uint64(len(proofBytes)))

		output = append(output, 0)
		output = append(output, proofBytesLength...)
		output = append(output, proofBytes...)
	}

	if ins.HasParent() {
		parentHashBytes := ins.Parent().Bytes()

		output = append(output, 1)
		output = append(output, parentHashBytes...)
	}

	return output, nil
}

// ToCommit converts content to a Commit instance
func (app *commitAdapter) ToCommit(content []byte) (Commit, error) {
	contentLength := len(content)
	if contentLength < commitMinSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Commit instance, %d provided", commitMinSize, contentLength)
		return nil, errors.New(str)
	}

	createdOnDelimiter := 8
	createdOnUnixNano := binary.LittleEndian.Uint64(content[0:createdOnDelimiter])
	createdOn := time.Unix(0, int64(createdOnUnixNano)).UTC()

	valuesBytesAmountDelimiter := createdOnDelimiter + 8
	valueBytesAmount := binary.LittleEndian.Uint64(content[createdOnDelimiter:valuesBytesAmountDelimiter])

	valueBytesDelimiter := valuesBytesAmountDelimiter + int(valueBytesAmount)
	values, err := app.hashTreeAdapter.ToHashTree(content[valuesBytesAmountDelimiter:valueBytesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[valueBytesDelimiter:]
	builder := app.builder.Create().WithValues(values).CreatedOn(createdOn)
	if len(remaining) > 0 {
		if remaining[0:1][0] == 0 {
			remaining = remaining[1:]
			proofBytesLengthDelimiter := 8
			proofBytesLength := binary.LittleEndian.Uint64(remaining[:proofBytesLengthDelimiter])

			proofBytesDelimiter := proofBytesLengthDelimiter + int(proofBytesLength)
			pProof := big.NewInt(int64(0)).SetBytes(remaining[proofBytesLengthDelimiter:proofBytesDelimiter])
			builder.WithProof(pProof)

			// reset the remaining:
			remaining = remaining[proofBytesDelimiter:]
		}
	}

	if len(remaining) > 0 {
		pParentHash, err := app.hashAdapter.FromBytes(remaining[1:])
		if err != nil {
			return nil, err
		}

		builder.WithParent(*pParentHash)
	}

	return builder.Now()
}
