package suites

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
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

// ToContent converts a Suite instance to content
func (app *adapter) ToContent(ins Suite) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	isValidByte := byte(0)
	if ins.IsValid() {
		isValidByte = byte(1)
	}

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, isValidByte)
	return append(output, ins.Content()...), nil
}

// ToSuite converts content to a Suite instance
func (app *adapter) ToSuite(content []byte) (Suite, error) {
	contentLength := len(content)
	if contentLength < minSuiteSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Suite instance, %d provided", minSuiteSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	delimiter := hash.Size + 1
	builder := app.builder.Create().WithContent(content[delimiter:]).WithHash(*pHash)
	if content[hash.Size:delimiter][0] != 0 {
		builder.IsValid()
	}

	return builder.Now()
}
