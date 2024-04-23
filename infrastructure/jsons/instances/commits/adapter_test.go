package commits

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
	jsons_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
	jsons_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/resources"
)

func TestAdapter_Success(t *testing.T) {
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
	actions := actions.NewActionsForTests([]actions.Action{
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

	content := commits.NewContentForTests(actions)
	signature, err := signers.NewFactory().Create().Sign(content.Hash().String())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := commits.NewCommitForTests(content, signature)
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

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_WithPrevious_Success(t *testing.T) {
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
	prevContent := commits.NewContentForTests(
		actions.NewActionsForTests([]actions.Action{
			actions.NewActionWithInsertAndDeleteForTests(
				insert,
				del,
			),
		}),
	)

	prevSig, err := signers.NewFactory().Create().Sign(prevContent.Hash().String())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := commits.NewContentWithPreviousForTests(
		actions.NewActionsForTests([]actions.Action{
			actions.NewActionWithInsertForTests(
				insert,
			),
			actions.NewActionWithDeleteForTests(
				del,
			),
		}),
		commits.NewCommitForTests(
			prevContent,
			prevSig,
		),
	)

	signature, err := signers.NewFactory().Create().Sign(content.Hash().String())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := commits.NewCommitForTests(content, signature)
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

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
