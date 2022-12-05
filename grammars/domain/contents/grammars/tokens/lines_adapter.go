package tokens

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type linesAdapter struct {
	builder     LinesBuilder
	lineAdapter LineAdapter
}

func createLinesAdapter(
	builder LinesBuilder,
	lineAdapter LineAdapter,
) LinesAdapter {
	out := linesAdapter{
		builder:     builder,
		lineAdapter: lineAdapter,
	}

	return &out
}

// ToContent converts lines to content
func (app *linesAdapter) ToContent(ins Lines) ([]byte, error) {
	list := ins.List()
	output := []byte{}
	for _, oneLine := range list {
		content, err := app.lineAdapter.ToContent(oneLine)
		if err != nil {
			return nil, err
		}

		lengthBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(lengthBytes, uint64(len(content)))

		output = append(output, lengthBytes...)
		output = append(output, content...)
	}

	return output, nil
}

// ToLines converts content to lines instance
func (app *linesAdapter) ToLines(content []byte) (Lines, error) {
	list := []Line{}
	remaining := content
	for len(remaining) > 0 {
		contentLength := len(remaining)
		if contentLength < 8 {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Lines instance, %d provided", 8, contentLength)
			return nil, errors.New(str)
		}

		bytesLength := int(binary.LittleEndian.Uint64(remaining[:8]))
		delimiter := bytesLength + 8
		if contentLength < delimiter {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Lines instance, %d provided", delimiter, contentLength)
			return nil, errors.New(str)
		}

		line, err := app.lineAdapter.ToLine(remaining[8:delimiter])
		if err != nil {
			return nil, err
		}

		list = append(list, line)
		remaining = remaining[delimiter:]
	}

	return app.builder.Create().WithList(list).Now()
}
