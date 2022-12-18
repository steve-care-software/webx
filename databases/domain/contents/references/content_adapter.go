package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type contentAdapter struct {
	contentKeysAdapter ContentKeysAdapter
	builder            ContentBuilder
}

func createContentAdapter(
	contentKeysAdapter ContentKeysAdapter,
	builder ContentBuilder,
) ContentAdapter {
	out := contentAdapter{
		contentKeysAdapter: contentKeysAdapter,
		builder:            builder,
	}

	return &out
}

// ToContent converts Content to bytes
func (app *contentAdapter) ToContent(ins Content) ([]byte, error) {
	active := []byte{}
	if ins.HasActive() {
		data, err := app.contentKeysAdapter.ToContent(ins.Active())
		if err != nil {
			return nil, err
		}

		active = data
	}

	pendings := []byte{}
	if ins.HasPendings() {
		data, err := app.contentKeysAdapter.ToContent(ins.Pendings())
		if err != nil {
			return nil, err
		}

		pendings = data
	}

	deleted := []byte{}
	if ins.HasDeleted() {
		data, err := app.contentKeysAdapter.ToContent(ins.Deleted())
		if err != nil {
			return nil, err
		}

		deleted = data
	}

	activeLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(activeLengthBytes, uint64(len(active)))

	pendingLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(pendingLengthBytes, uint64(len(pendings)))

	output := []byte{}
	output = append(output, activeLengthBytes...)
	output = append(output, active...)
	output = append(output, pendingLengthBytes...)
	output = append(output, pendings...)
	output = append(output, deleted...)
	return output, nil
}

// ToInstance converts bytes to Content
func (app *contentAdapter) ToInstance(content []byte) (Content, error) {
	contentLength := len(content)
	if contentLength < 8 {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the active ContentKeys size of the Content instance, %d provided", 8, contentLength)
		return nil, errors.New(str)
	}

	activeBytesLength := binary.LittleEndian.Uint64(content[:8])
	activeBytesDelimiter := int(activeBytesLength + 8)
	if contentLength < activeBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the active ContentKeys size of the Content instance, %d provided", activeBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	builder := app.builder.Create()
	if activeBytesLength > 0 {
		active, err := app.contentKeysAdapter.ToContentKeys(content[8:activeBytesDelimiter])
		if err != nil {
			return nil, err
		}

		builder.WithActive(active)
	}

	pendingBytesLengthDelimiter := activeBytesDelimiter + 8
	pendingBytesLength := binary.LittleEndian.Uint64(content[activeBytesDelimiter:pendingBytesLengthDelimiter])
	pendingBytesDelimiter := pendingBytesLengthDelimiter + int(pendingBytesLength)
	if contentLength < pendingBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the pending ContentKeys size of the Content instance, %d provided", pendingBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	if pendingBytesLength > 0 {
		pendings, err := app.contentKeysAdapter.ToContentKeys(content[pendingBytesLengthDelimiter:pendingBytesDelimiter])
		if err != nil {
			return nil, err
		}

		builder.WithPendings(pendings)
	}

	deletedBytes := content[pendingBytesDelimiter:]
	if len(deletedBytes) > 0 {
		deleted, err := app.contentKeysAdapter.ToContentKeys(deletedBytes)
		if err != nil {
			return nil, err
		}

		builder.WithDeleted(deleted)
	}

	return builder.Now()
}
