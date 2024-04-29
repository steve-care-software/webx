package accounts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/retrieves"
)

func TestAdapter_withList_Success(t *testing.T) {
	ins := accounts.NewAccountWithListForTests(
		"myList",
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

func TestAdapter_withCredentials_Success(t *testing.T) {
	ins := accounts.NewAccountWithCredentialsForTests(
		credentials.NewCredentialsForTests(
			"myUsername",
			"myPassword",
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

func TestAdapter_withRetrieve_Success(t *testing.T) {
	ins := accounts.NewAccountWithRetrieveForTests(
		retrieves.NewRetrieveForTests(
			"myPassword",
			"myCredentials",
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

func TestAdapter_withCommunication_Success(t *testing.T) {
	ins := accounts.NewAccountWithCommunicationForTests(
		communications.NewCommunicationWithSignForTests(
			signs.NewSignForTests("myMessage", "myAccount"),
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

func TestAdapter_withEncryption_Success(t *testing.T) {
	ins := accounts.NewAccountWithEncryptionForTests(
		encryptions.NewEncryptionWithDecryptForTests(
			decrypts.NewDecryptForTests(
				"myCipher",
				"myAccount",
			),
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
