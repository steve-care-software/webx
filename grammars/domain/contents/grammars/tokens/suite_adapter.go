package tokens

import (
	"errors"
	"fmt"
)

type suiteAdapter struct {
	builder SuiteBuilder
}

func createSuiteAdapter(
	builder SuiteBuilder,
) SuiteAdapter {
	out := suiteAdapter{
		builder: builder,
	}

	return &out
}

// ToContent converts a Suite instance to content
func (app *suiteAdapter) ToContent(ins Suite) ([]byte, error) {
	isValidByte := byte(0)
	if ins.IsValid() {
		isValidByte = byte(1)
	}

	output := []byte{
		isValidByte,
	}

	return append(output, ins.Content()...), nil
}

// ToSuite converts content to a Suite instance
func (app *suiteAdapter) ToSuite(content []byte) (Suite, error) {
	contentLength := len(content)
	if contentLength < minSuiteSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Suite instance, %d provided", minSuiteSize, contentLength)
		return nil, errors.New(str)
	}

	builder := app.builder.Create().WithContent(content[1:])
	if content[:1][0] != 0 {
		builder.IsValid()
	}

	return builder.Now()
}
