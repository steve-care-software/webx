package blockchains

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

// ToContent converts a blockchain to bytes
func (app *adapter) ToContent(ins Blockchain) ([]byte, error) {
	referenceBytes := ins.Reference().Bytes()
	headBytes := ins.Head().Bytes()

	output := []byte{}
	output = append(output, referenceBytes...)
	output = append(output, headBytes...)
	return output, nil
}

// ToBlockchain converts bytes to a Blockchain instance
func (app *adapter) ToBlockchain(content []byte) (Blockchain, error) {
	contentLength := len(content)
	if contentLength != blockchainSize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert data to a Blockchain instance, %d provided", blockchainSize, contentLength)
		return nil, errors.New(str)
	}

	pReference, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	pHead, err := app.hashAdapter.FromBytes(content[hash.Size:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithReference(*pReference).
		WithHead(*pHead).
		Now()
}
