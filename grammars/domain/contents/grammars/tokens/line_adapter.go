package tokens

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type lineAdapter struct {
	hashAdapter hash.Adapter
	builder     LineBuilder
}

func createLineAdapter(
	hashAdapter hash.Adapter,
	builder LineBuilder,
) LineAdapter {
	out := lineAdapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToContent converts a Line instance to content
func (app *lineAdapter) ToContent(ins Line) ([]byte, error) {
	output := []byte{}
	elements := ins.Elements()
	for _, oneElement := range elements {
		output = append(output, oneElement.Bytes()...)
	}

	return output, nil
}

// ToLine converts content to a Line instance
func (app *lineAdapter) ToLine(content []byte) (Line, error) {
	length := len(content)
	if length%hash.Size != 0 {
		str := fmt.Sprintf("the content's length (%d) was expected to be a multiple of %d in order to convert it to a Line instance", length, hash.Size)
		return nil, errors.New(str)
	}

	elements := []hash.Hash{}
	amount := len(content) / hash.Size
	for i := 0; i < amount; i++ {
		startsOn := i * hash.Size
		endsOn := startsOn + hash.Size
		pHash, err := app.hashAdapter.FromBytes(content[startsOn:endsOn])
		if err != nil {
			return nil, err
		}

		elements = append(elements, *pHash)
	}

	return app.builder.Create().WithElements(elements).Now()
}
