package channels

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
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

// ToContent converts channel to bytes
func (app *adapter) ToContent(ins Channel) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	tokenBytes := ins.Token().Bytes()
	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, tokenBytes...)

	hasPreviousByte := byte(0)
	if ins.HasPrevious() {
		hasPreviousByte = byte(1)
	}

	output = append(output, hasPreviousByte)

	if ins.HasPrevious() {
		previousBytes := ins.Previous().Bytes()
		output = append(output, previousBytes...)
	}

	if ins.HasNext() {
		nextBytes := ins.Next().Bytes()
		output = append(output, nextBytes...)
	}

	return output, nil
}

// ToChannel converts bytes to a channel instance
func (app *adapter) ToChannel(content []byte) (Channel, error) {
	contentLength := len(content)
	if contentLength < minChannelSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Channel instance, %d provided", minChannelSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	tokenDelimiter := hash.Size + hash.Size
	pToken, err := app.hashAdapter.FromBytes(content[hash.Size:tokenDelimiter])
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHash(*pHash).WithToken(*pToken)
	hasPreviousDelimiter := tokenDelimiter + 1

	hasPrevious := false
	if content[tokenDelimiter:hasPreviousDelimiter][0] != 0 {
		hasPrevious = true
	}

	previousDelimiter := hasPreviousDelimiter
	if hasPrevious {
		previousDelimiter = previousDelimiter + hash.Size
		pPrevious, err := app.hashAdapter.FromBytes(content[hasPreviousDelimiter:previousDelimiter])
		if err != nil {
			return nil, err
		}

		builder.WithPrevious(*pPrevious)
	}

	remaining := content[previousDelimiter:]
	if len(remaining) > 0 {
		pNext, err := app.hashAdapter.FromBytes(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithNext(*pNext)
	}

	return builder.Now()
}
