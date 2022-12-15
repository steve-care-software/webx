package everythings

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type adapter struct {
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToContent converts an everything instance to content
func (app *adapter) ToContent(ins Everything) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	exceptionBytes := ins.Exception().Bytes()

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, exceptionBytes...)

	if ins.HasEscape() {
		escapeBytes := ins.Escape().Bytes()
		output = append(output, escapeBytes...)
	}

	return output, nil
}

// ToEverything converts content to an everything instance
func (app *adapter) ToEverything(content []byte) (Everything, error) {
	contentLength := len(content)
	if contentLength < minEverythingSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to an Everything instance, %d provided", minEverythingSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	exceptionDelimiter := hash.Size + hash.Size
	pException, err := app.hashAdapter.FromBytes(content[hash.Size:exceptionDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[exceptionDelimiter:]
	builder := app.builder.Create().WithHash(*pHash).WithException(*pException)
	if len(remaining) > 0 {
		pEscape, err := app.hashAdapter.FromBytes(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithEscape(*pEscape)
	}

	return builder.Now()
}
