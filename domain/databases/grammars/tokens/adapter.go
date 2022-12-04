package tokens

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type adapter struct {
	hashAdapter   hash.Adapter
	linesAdapter  LinesAdapter
	suitesAdapter SuitesAdapter
	builder       Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	linesAdapter LinesAdapter,
	suitesAdapter SuitesAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:   hashAdapter,
		linesAdapter:  linesAdapter,
		suitesAdapter: suitesAdapter,
		builder:       builder,
	}

	return &out
}

// ToContent converts token instance to content
func (app *adapter) ToContent(ins Token) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	linesBytes, err := app.linesAdapter.ToContent(ins.Lines())
	if err != nil {
		return nil, err
	}

	lineBytesLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(lineBytesLength, uint64(len(linesBytes)))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, lineBytesLength...)
	output = append(output, linesBytes...)

	if ins.HasSuites() {
		suitesBytes, err := app.suitesAdapter.ToContent(ins.Suites())
		if err != nil {
			return nil, err
		}

		output = append(output, suitesBytes...)
	}

	return output, nil
}

// ToToken converts content to token instance
func (app *adapter) ToToken(content []byte) (Token, error) {
	contentLength := len(content)
	if contentLength < minTokenSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Token instance, %d provided", minTokenSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	linesBytesLengthDelimiter := hash.Size + 8
	linesBytesLength := binary.LittleEndian.Uint64(content[hash.Size:linesBytesLengthDelimiter])

	linesDelimiter := linesBytesLengthDelimiter + int(linesBytesLength)
	lines, err := app.linesAdapter.ToLines(content[linesBytesLengthDelimiter:linesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[linesDelimiter:]
	builder := app.builder.Create().WithHash(*pHash).WithLines(lines)
	if len(remaining) > 0 {
		suites, err := app.suitesAdapter.ToSuites(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithSuites(suites)
	}

	return builder.Now()
}
