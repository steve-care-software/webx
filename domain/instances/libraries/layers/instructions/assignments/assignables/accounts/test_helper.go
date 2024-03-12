package accounts

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
)

// NewAccountWithListForTests creates a new account with list for tests
func NewAccountWithListForTests(list string) Account {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithCredentialsForTests creates a new account with credentials for tests
func NewAccountWithCredentialsForTests(credentials credentials.Credentials) Account {
	ins, err := NewBuilder().Create().WithCredentials(credentials).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithRetrieveForTests creates a new account with retrieve for tests
func NewAccountWithRetrieveForTests(retrieve retrieves.Retrieve) Account {
	ins, err := NewBuilder().Create().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithCommunicationForTests creates a new account with communication for tests
func NewAccountWithCommunicationForTests(communication communications.Communication) Account {
	ins, err := NewBuilder().Create().WithCommunication(communication).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithEncryptionForTests creates a new account with encryption for tests
func NewAccountWithEncryptionForTests(encryption encryptions.Encryption) Account {
	ins, err := NewBuilder().Create().WithEncryption(encryption).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
