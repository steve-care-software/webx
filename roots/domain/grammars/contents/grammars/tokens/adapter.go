package tokens

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type adapter struct {
	hashAdapter  hash.Adapter
	linesAdapter LinesAdapter
	builder      Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	linesAdapter LinesAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:  hashAdapter,
		linesAdapter: linesAdapter,
		builder:      builder,
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

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, linesBytes...)
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

	lines, err := app.linesAdapter.ToLines(content[hash.Size:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithHash(*pHash).WithLines(lines).Now()
}
