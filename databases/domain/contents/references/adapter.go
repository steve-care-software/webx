package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type adapter struct {
	contentKeysAdapter ContentKeysAdapter
	commitsAdapter     CommitsAdapter
	builder            Builder
}

func createAdapter(
	contentKeysAdapter ContentKeysAdapter,
	commitsAdapter CommitsAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		contentKeysAdapter: contentKeysAdapter,
		commitsAdapter:     commitsAdapter,
		builder:            builder,
	}
	return &out
}

// ToContent converts reference to bytes
func (app *adapter) ToContent(ins Reference) ([]byte, error) {
	contentBytes, err := app.contentKeysAdapter.ToContent(ins.ContentKeys())
	if err != nil {
		return nil, err
	}

	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(len(contentBytes)))

	commitsBytes, err := app.commitsAdapter.ToContent(ins.Commits())
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, lengthBytes...)
	output = append(output, contentBytes...)
	output = append(output, commitsBytes...)
	return output, nil
}

// ToReference converts bytes to reference
func (app *adapter) ToReference(content []byte) (Reference, error) {
	contentLength := len(content)
	if contentLength < minReferenceSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Reference instance, %d provided", minReferenceSize, contentLength)
		return nil, errors.New(str)
	}

	contentKeysBytesLengthDelimiter := uint64(8)
	contentKeysBytesLength := binary.LittleEndian.Uint64(content[:contentKeysBytesLengthDelimiter])
	contentKeysBytesDelimiter := int(contentKeysBytesLength + contentKeysBytesLengthDelimiter)
	if contentLength < contentKeysBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the contentKeys size of the Reference instance, %d provided", contentKeysBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	contentKeys, err := app.contentKeysAdapter.ToContentKeys(content[contentKeysBytesLengthDelimiter:contentKeysBytesDelimiter])
	if err != nil {
		return nil, err
	}

	commits, err := app.commitsAdapter.ToCommits(content[contentKeysBytesDelimiter:])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithContentKeys(contentKeys).WithCommits(commits).Now()
}
