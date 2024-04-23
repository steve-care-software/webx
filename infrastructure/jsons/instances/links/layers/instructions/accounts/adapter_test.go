package accounts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates/criterias"
)

func TestAdapter_withInsert_Success(t *testing.T) {
	ins := accounts.NewAccountWithInsertForTests(
		inserts.NewInsertForTests("myUser", "myPass"),
	)

	adapter := NewAdapter()

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

func TestAdapter_withUpdate_Success(t *testing.T) {
	ins := accounts.NewAccountWithUpdateForTests(
		updates.NewUpdateForTests(
			"myCredentials",
			criterias.NewCriteriaForTests(true, false),
		),
	)

	adapter := NewAdapter()

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

func TestAdapter_withDelete_Success(t *testing.T) {
	ins := accounts.NewAccountWithDeleteForTests("myAccount")
	adapter := NewAdapter()

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
