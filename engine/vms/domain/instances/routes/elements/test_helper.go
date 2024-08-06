package elements

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// NewElementsForTests creates elements for tests
func NewElementsForTests(list []Element) Elements {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithLayerForTests creates element with layer for tests
func NewElementWithLayerForTests(layer hash.Hash) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithBytesForTests creates element with bytes for tests
func NewElementWithBytesForTests(bytes []byte) Element {
	ins, err := NewElementBuilder().Create().WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithStringForTests creates element with string for tests
func NewElementWithStringForTests(str string) Element {
	ins, err := NewElementBuilder().Create().WithString(str).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
