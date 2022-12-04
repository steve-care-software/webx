package elements

import (
	"encoding/binary"
	"errors"
)

type cardinalityAdapter struct {
	builder CardinalityBuilder
}

func createCardinalityAdapter(
	builder CardinalityBuilder,
) CardinalityAdapter {
	out := cardinalityAdapter{
		builder: builder,
	}

	return &out
}

// ToContent converts a Cardinality to bytes
func (app *cardinalityAdapter) ToContent(ins Cardinality) ([]byte, error) {
	minBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(minBytes, uint64(ins.Min()))

	output := []byte{}
	output = append(output, minBytes...)

	if ins.HasMax() {
		pMax := ins.Max()
		maxBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(maxBytes, uint64(*pMax))

		output = append(output, maxBytes...)
	}

	return output, nil
}

// ToCardinality converts bytes to a Cardinality instance
func (app *cardinalityAdapter) ToCardinality(content []byte) (Cardinality, error) {
	contentLength := len(content)
	if contentLength < 8 {
		str := "the content was expected to contain at least %d bytes in order to convert data to a Cardinality instance, %d provided"
		return nil, errors.New(str)
	}

	min := binary.LittleEndian.Uint64(content[:8])
	remaining := content[8:]
	builder := app.builder.Create().WithMin(uint(min))
	if len(remaining) == 8 {
		max := binary.LittleEndian.Uint64(content[8 : 8*2])
		builder.WithMax(uint(max))
	}

	return builder.Now()
}
