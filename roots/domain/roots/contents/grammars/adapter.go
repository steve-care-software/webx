package grammars

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
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

// ToContent returns the grammar to content
func (app *adapter) ToContent(ins Grammar) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	nameBytes := []byte(ins.Name())

	nameLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nameLengthBytes, uint64(len(nameBytes)))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, nameLengthBytes...)
	output = append(output, nameBytes...)

	if ins.HasHistory() {
		historyBytes, err := app.hashTreeAdapter.ToContent(ins.History())
		if err != nil {
			return nil, err
		}

		output = append(output, historyBytes...)
	}

	return output, nil
}

// ToGrammar returns the content to a grammar instance
func (app *adapter) ToGrammar(content []byte) (Grammar, error) {
	contentLength := len(content)
	if contentLength < minGrammarLength {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Grammar instance, %d provided", minGrammarLength, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	nameDelimiterLength := hash.Size + 8
	nameLength := binary.LittleEndian.Uint64(content[hash.Size:nameDelimiterLength])

	nameDelimiter := nameDelimiterLength + int(nameLength)
	name := string(content[nameDelimiterLength:nameDelimiter])

	remaining := content[nameDelimiter:]
	builder := app.builder.Create().WithHash(*pHash).WithName(name)
	if len(remaining) > 0 {
		hashTreeIns, err := app.hashTreeAdapter.ToHashTree(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithHistory(hashTreeIns)
	}

	return builder.Now()
}
