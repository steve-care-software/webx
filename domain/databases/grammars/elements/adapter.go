package elements

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type adapter struct {
	cardinalityAdapter CardinalityAdapter
	hashAdapter        hash.Adapter
	builder            Builder
}

func createAdapter(
	cardinalityAdapter CardinalityAdapter,
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		cardinalityAdapter: cardinalityAdapter,
		hashAdapter:        hashAdapter,
		builder:            builder,
	}

	return &out
}

// ToContent converts an element to bytes
func (app *adapter) ToContent(ins Element) ([]byte, error) {
	return nil, nil
}

// ToElement converts bytes to an Element instance
func (app *adapter) ToElement(content []byte) (Element, error) {
	return nil, nil
}
