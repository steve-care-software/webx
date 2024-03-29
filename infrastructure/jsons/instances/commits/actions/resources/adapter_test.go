package resources

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
	jsons_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
)

func TestAdapter_Success(t *testing.T) {
	pIdentifier, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := resources.NewResourceForTests(
		[]string{
			"first",
			"second",
		},
		pointers.NewPointerForTests(
			[]string{
				"first",
				"second",
				"third",
			},
			*pIdentifier,
		),
	)

	adapter := NewAdapter(
		createTestInstanceAdapter(
			jsons_pointers.NewAdapter(),
		),
	)

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(ins, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
