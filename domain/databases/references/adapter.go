package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type adapter struct {
	contentAdapter    ContentAdapter
	blockchainAdapter BlockchainAdapter
	builder           Builder
}

func createAdapter(
	contentAdapter ContentAdapter,
	blockchainAdapter BlockchainAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		contentAdapter:    contentAdapter,
		blockchainAdapter: blockchainAdapter,
		builder:           builder,
	}
	return &out
}

// ToContent converts reference to bytes
func (app *adapter) ToContent(ins Reference) ([]byte, error) {
	contentBytes, err := app.contentAdapter.ToContent(ins.Content())
	if err != nil {
		return nil, err
	}

	contentLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(contentLengthBytes, uint64(len(contentBytes)))

	output := []byte{}
	output = append(output, contentLengthBytes...)
	output = append(output, contentBytes...)

	if ins.HasBlockchain() {
		blockchainBytes, err := app.blockchainAdapter.ToContent(ins.Blockchain())
		if err != nil {
			return nil, err
		}

		output = append(output, blockchainBytes...)
	}

	return output, nil
}

// ToReference converts bytes to reference
func (app *adapter) ToReference(content []byte) (Reference, error) {
	contentLength := len(content)
	if contentLength < 8 {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the content size of the Reference instance, %d provided", 8, contentLength)
		return nil, errors.New(str)
	}

	contentBytesLength := binary.LittleEndian.Uint64(content[:8])
	contentBytesDelimiter := int(contentBytesLength + 8)
	if contentLength < contentBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the content size of the Reference instance, %d provided", contentBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	contentIns, err := app.contentAdapter.ToInstance(content[8:contentBytesDelimiter])
	if err != nil {
		return nil, err
	}

	blockchainBytes := content[contentBytesDelimiter:]
	builder := app.builder.Create().WithContent(contentIns)
	if len(blockchainBytes) > 0 {
		blockchain, err := app.blockchainAdapter.ToBlockchain(blockchainBytes)
		if err != nil {
			return nil, err
		}

		builder.WithBlockchain(blockchain)
	}

	return builder.Now()
}
