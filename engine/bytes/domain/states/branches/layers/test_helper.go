package layers

import "github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers/pointers"

// NewLayersForTests creates a new layers for tests
func NewLayersForTests(list []Layer) Layers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(isDeleted bool) Layer {
	builder := NewLayerBuilder().Create()
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func NewLayerWithPointersForTests(isDeleted bool, pointers pointers.Pointers) Layer {
	builder := NewLayerBuilder().Create().WithPointers(pointers)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
