package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
)

type testInstanceAdapter struct {
	pointerAdapter pointers.Adapter
}

func createTestInstanceAdapter(
	pointerAdapter pointers.Adapter,
) instances.Adapter {
	out := testInstanceAdapter{
		pointerAdapter: pointerAdapter,
	}

	return &out
}

// ToBytes converts instances to bytes
func (app *testInstanceAdapter) ToBytes(path []string, ins instances.Instance) ([]byte, error) {
	return app.pointerAdapter.ToBytes(ins.(pointers.Pointer))
}

// ToInstance converts bytes to instance
func (app *testInstanceAdapter) ToInstance(path []string, data []byte) (instances.Instance, error) {
	return app.pointerAdapter.ToInstance(data)
}
