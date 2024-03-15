package logics

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics/locations"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

// NewLogicForTests creates a new logic for tests
func NewLogicForTests(layer layers.Layer, location locations.Location) Logic {
	ins, err := NewBuilder().Create().WithLayer(layer).WithLocation(location).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
