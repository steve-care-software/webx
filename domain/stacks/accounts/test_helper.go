package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// NewAccountWithAccountForTests creates an account with account for tests
func NewAccountWithAccountForTests(value accounts.Account) Account {
	ins, err := NewBuilder().Create().WithAccount(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithCredentialsForTests creates an account with credentials for tests
func NewAccountWithCredentialsForTests(value credentials.Credentials) Account {
	ins, err := NewBuilder().Create().WithCredentials(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithRingForTests creates an account with ring for tests
func NewAccountWithRingForTests(value []signers.PublicKey) Account {
	ins, err := NewBuilder().Create().WithRing(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithSignatureForTests creates an account with signature for tests
func NewAccountWithSignatureForTests(value signers.Signature) Account {
	ins, err := NewBuilder().Create().WithSignature(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAccountWithVoteForTests creates an account with vote for tests
func NewAccountWithVoteForTests(value signers.Vote) Account {
	ins, err := NewBuilder().Create().WithVote(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
