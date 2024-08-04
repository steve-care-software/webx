package layers

import "github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers/pointers"

type layer struct {
	isDeleted bool
	pointers  pointers.Pointers
}

func createLayer(
	isDeleted bool,
) Layer {
	return createLayerInternally(isDeleted, nil)
}

func createLayerWithPointers(
	isDeleted bool,
	pointers pointers.Pointers,
) Layer {
	return createLayerInternally(isDeleted, pointers)
}

func createLayerInternally(
	isDeleted bool,
	pointers pointers.Pointers,
) Layer {
	out := layer{
		isDeleted: isDeleted,
		pointers:  pointers,
	}

	return &out
}

// IsDeleted returns true if deleted, false otherwise
func (obj *layer) IsDeleted() bool {
	return obj.isDeleted
}

// HasPointers returns true if there is pointers, false otherwise
func (obj *layer) HasPointers() bool {
	return obj.pointers != nil
}

// Pointers returns the pointers, if any
func (obj *layer) Pointers() pointers.Pointers {
	return obj.pointers
}
