package commits

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type adapter struct {
	hashAdapter     hash.Adapter
	hashTreeAdapter hashtrees.Adapter
	builder         Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	hashTreeAdapter hashtrees.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:     hashAdapter,
		hashTreeAdapter: hashTreeAdapter,
		builder:         builder,
	}

	return &out
}

// ToContent converts a Commit instance to content
func (app *adapter) ToContent(ins Commit) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	valuesBytes, err := app.hashTreeAdapter.ToContent(ins.Values())
	if err != nil {
		return nil, err
	}

	valuesBytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuesBytesAmount, uint64(len(valuesBytes)))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, valuesBytesAmount...)
	output = append(output, valuesBytes...)
	if ins.HasParent() {
		parentHashBytes := ins.Parent().Bytes()
		output = append(output, parentHashBytes...)
	}

	return output, nil
}

// ToCommit converts content to a Commit instance
func (app *adapter) ToCommit(content []byte) (Commit, error) {
	contentLength := len(content)
	if contentLength < commitMinSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Commit instance, %d provided", commitMinSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	valuesBytesAmountDelimiter := hash.Size + 8
	valueBytesAmount := binary.LittleEndian.Uint64(content[hash.Size:valuesBytesAmountDelimiter])

	valueBytesDelimiter := valuesBytesAmountDelimiter + int(valueBytesAmount)
	values, err := app.hashTreeAdapter.ToHashTree(content[valuesBytesAmountDelimiter:valueBytesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[valueBytesDelimiter:]
	builder := app.builder.Create().WithHash(*pHash).WithValues(values)
	if len(remaining) > 0 {
		pParentHash, err := app.hashAdapter.FromBytes(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithParent(*pParentHash)
	}

	return builder.Now()
}
