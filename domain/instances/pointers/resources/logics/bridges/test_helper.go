package bridges

import "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"

// NewBridgesForTests creates new bridges for tests
func NewBridgesForTests(list []Bridge) Bridges {
	ins, err := NewBuiler().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBridgeForTests creates a new bridge for tests
func NewBridgeForTests(path []string, layer layers.Layer) Bridge {
	ins, err := NewBridgeBuilder().Create().WithPath(path).WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
