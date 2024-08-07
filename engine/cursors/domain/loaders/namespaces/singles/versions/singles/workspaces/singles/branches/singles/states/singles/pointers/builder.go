package pointers

import (
	"errors"
	"sort"
)

type builder struct {
	list []Pointer
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Pointer) Builder {
	app.list = list
	return app
}

// Now builds a new Pointers instance
func (app *builder) Now() (Pointers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Pointer in order to build a Pointers instance")
	}

	// order by index
	lastIndex := uint64(0)
	indices := []int{}
	orderedPointers := map[uint64]Pointer{}
	for idx, onePointer := range app.list {
		delimiter := onePointer.Storage().Delimiter()
		index := delimiter.Index()
		if idx <= 0 {
			lastIndex = index
			indices = append(indices, int(index))
			orderedPointers[lastIndex] = onePointer
			continue
		}

		orderedPointers[index] = onePointer
		indices = append(indices, int(index))
		lastIndex = index
	}

	length := len(indices)
	sort.Ints(indices)
	output := []Pointer{}
	for i := 0; i < length; i++ {
		idx := uint64(indices[i])
		output = append(output, orderedPointers[idx])
	}

	return createPointers(
		output,
	), nil
}
