package tokens

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type suitesAdapter struct {
	builder      SuitesBuilder
	suiteAdapter SuiteAdapter
}

func createSuitesAdapter(
	builder SuitesBuilder,
	suiteAdapter SuiteAdapter,
) SuitesAdapter {
	out := suitesAdapter{
		builder:      builder,
		suiteAdapter: suiteAdapter,
	}

	return &out
}

// ToContent converts suites to content
func (app *suitesAdapter) ToContent(ins Suites) ([]byte, error) {
	list := ins.List()
	output := []byte{}
	for _, oneSuite := range list {
		content, err := app.suiteAdapter.ToContent(oneSuite)
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

// ToSuites converts content to suites instance
func (app *suitesAdapter) ToSuites(content []byte) (Suites, error) {
	list := []Suite{}
	remaining := content
	for len(remaining) > 0 {
		contentLength := len(remaining)
		if contentLength < 8 {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Suites instance, %d provided", 8, contentLength)
			return nil, errors.New(str)
		}

		bytesLength := int(binary.LittleEndian.Uint64(remaining[:8]))
		delimiter := bytesLength + 8
		if contentLength < delimiter {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Suites instance, %d provided", delimiter, contentLength)
			return nil, errors.New(str)
		}

		suite, err := app.suiteAdapter.ToSuite(remaining[8:delimiter])
		if err != nil {
			return nil, err
		}

		list = append(list, suite)
		remaining = remaining[delimiter:]
	}

	return app.builder.Create().WithList(list).Now()
}
