package contents

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type adapter struct {
	builder        Builder
	contentAdapter ContentAdapter
}

func createAdapter(
	builder Builder,
	contentAdapter ContentAdapter,
) Adapter {
	out := adapter{
		builder:        builder,
		contentAdapter: contentAdapter,
	}

	return &out
}

// ToContent converts contents to bytes
func (app *adapter) ToContent(ins Contents) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneContent := range list {
		oneBytes, err := app.contentAdapter.ToContent(oneContent)
		if err != nil {
			return nil, err
		}

		sizeBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(sizeBytes, uint64(len(oneBytes)))

		output = append(output, sizeBytes...)
		output = append(output, oneBytes...)
	}

	return output, nil
}

// ToInstance converts bytes to contents instance
func (app *adapter) ToInstance(content []byte) (Contents, error) {
	contentLength := len(content)
	if contentLength < minContentsSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Contents instance, %d provided", minContentsSize, contentLength)
		return nil, errors.New(str)
	}

	list := []Content{}
	remaining := content
	for {
		if len(remaining) <= 0 {
			break
		}

		sizeDelimiter := 8
		size := binary.LittleEndian.Uint64(remaining[:sizeDelimiter])

		contentDelimiter := sizeDelimiter + int(size)
		contentIns, err := app.contentAdapter.ToInstance(remaining[sizeDelimiter:contentDelimiter])
		if err != nil {
			return nil, err
		}

		list = append(list, contentIns)
		remaining = remaining[contentDelimiter:]
	}

	return app.builder.Create().
		WithList(list).
		Now()
}
