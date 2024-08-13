package updates

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// NewUpdateWithDataForTests creates a new update with data for tests
func NewUpdateWithDataForTests(name string, data []byte) Update {
	ins, err := NewBuilder().Create().WithName(name).WithData(data).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithAdditionForTests creates a new update with aaddition
func NewUpdateWithAdditionForTests(name string, addition []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithWhiteListAddition(addition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithRemovalForTests creates a new update with removal
func NewUpdateWithRemovalForTests(name string, removal []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithWhiteListRemoval(removal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithDataAndAdditionForTests creates a new update with data and addition
func NewUpdateWithDataAndAdditionForTests(name string, data []byte, addition []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithData(data).WithWhiteListAddition(addition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithDataAndRemovalForTests creates a new update with data and removal
func NewUpdateWithDataAndRemovalForTests(name string, data []byte, removal []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithData(data).WithWhiteListRemoval(removal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithAdditionAndRemovalForTests creates a new update with addition and removal
func NewUpdateWithAdditionAndRemovalForTests(name string, addition []hash.Hash, removal []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithWhiteListAddition(addition).WithWhiteListRemoval(removal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewUpdateWithDataAndAdditionAndRemovalForTests creates a new update with data and addition and removal
func NewUpdateWithDataAndAdditionAndRemovalForTests(name string, data []byte, addition []hash.Hash, removal []hash.Hash) Update {
	ins, err := NewBuilder().Create().WithName(name).WithData(data).WithWhiteListAddition(addition).WithWhiteListRemoval(removal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
