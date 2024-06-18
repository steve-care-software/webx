package pointers

import "github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(path []string, isActive bool) Pointer {
	builder := NewPointerBuilder().Create().
		WithPath(path)

	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerWithLoaderForTests creates a new pointer with loader for tests
func NewPointerWithLoaderForTests(path []string, isActive bool, loader conditions.Condition) Pointer {
	builder := NewPointerBuilder().Create().
		WithPath(path).
		WithLoader(loader)

	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerWithCancellerForTests creates a new pointer with canceller for tests
func NewPointerWithCancellerForTests(path []string, isActive bool, canceller conditions.Condition) Pointer {
	builder := NewPointerBuilder().Create().
		WithPath(path).
		WithCanceller(canceller)

	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerWithLoaderAndCancellerForTests creates a new pointer with loader and canceller for tests
func NewPointerWithLoaderAndCancellerForTests(path []string, isActive bool, loader conditions.Condition, canceller conditions.Condition) Pointer {
	builder := NewPointerBuilder().Create().
		WithPath(path).
		WithLoader(loader).
		WithCanceller(canceller)

	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
