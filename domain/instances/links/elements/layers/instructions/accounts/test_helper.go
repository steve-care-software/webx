package accounts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates"
)

// NewAccountWithInsertForTests creates an account with insert for tests
func NewAccountWithInsertForTests(insert inserts.Insert) Account {
	ins, err := NewBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithUpdateForTests creates an account with update for tests
func NewAccountWithUpdateForTests(update updates.Update) Account {
	ins, err := NewBuilder().Create().WithUpdate(update).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithDeleteForTests creates an account with delete for tests
func NewAccountWithDeleteForTests(value string) Account {
	ins, err := NewBuilder().Create().WithDelete(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
