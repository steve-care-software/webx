package actions

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
	jsons_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
	jsons_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/resources"
)

func TestActionAdapter_Success(t *testing.T) {
	pIdentifier, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	insert := resources.NewResourceForTests(
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

	path := []string{
		"first",
		"second",
	}

	pDelIdentifier, err := hash.NewAdapter().FromBytes([]byte("this is some other bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	del := pointers.NewPointerForTests(path, *pDelIdentifier)
	ins := actions.NewActionsForTests([]actions.Action{
		actions.NewActionWithInsertForTests(
			insert,
		),
		actions.NewActionWithDeleteForTests(
			del,
		),
		actions.NewActionWithInsertAndDeleteForTests(
			insert,
			del,
		),
	})

	adapter := NewAdapter(
		jsons_resources.NewTestInstanceAdapter(
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
