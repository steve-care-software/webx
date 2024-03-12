package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

type account struct {
	accountIns  accounts.Account
	credentials credentials.Credentials
	ring        []signers.PublicKey
	signature   signers.Signature
	vote        signers.Vote
}

func createAccountWithAccount(
	accountIns accounts.Account,
) Account {
	return createAccountInternally(accountIns, nil, nil, nil, nil)
}

func createAccountWithCredentials(
	credentials credentials.Credentials,
) Account {
	return createAccountInternally(nil, credentials, nil, nil, nil)
}

func createAccountWithRing(
	ring []signers.PublicKey,
) Account {
	return createAccountInternally(nil, nil, ring, nil, nil)
}

func createAccountWithSignature(
	signature signers.Signature,
) Account {
	return createAccountInternally(nil, nil, nil, signature, nil)
}

func createAccountWithVote(
	vote signers.Vote,
) Account {
	return createAccountInternally(nil, nil, nil, nil, vote)
}

func createAccountInternally(
	accountIns accounts.Account,
	credentials credentials.Credentials,
	ring []signers.PublicKey,
	signature signers.Signature,
	vote signers.Vote,
) Account {
	out := account{
		accountIns:  accountIns,
		credentials: credentials,
		ring:        ring,
		signature:   signature,
		vote:        vote,
	}

	return &out
}

// IsAccount returns true if there is an account, false otherwise
func (obj *account) IsAccount() bool {
	return obj.accountIns != nil
}

// Account returns the account, if any
func (obj *account) Account() accounts.Account {
	return obj.accountIns
}

// IsCredentials returns true if there is a credentials, false otherwise
func (obj *account) IsCredentials() bool {
	return obj.credentials != nil
}

// Credentials returns the credentials, if any
func (obj *account) Credentials() credentials.Credentials {
	return obj.credentials
}

// IsRing returns true if there is a ring, false otherwise
func (obj *account) IsRing() bool {
	return obj.ring != nil
}

// Ring returns the ring, if any
func (obj *account) Ring() []signers.PublicKey {
	return obj.ring
}

// IsSignature returns true if there is a signature, false otherwise
func (obj *account) IsSignature() bool {
	return obj.signature != nil
}

// Signature returns the signature, if any
func (obj *account) Signature() signers.Signature {
	return obj.signature
}

// IsVote returns true if there is a vote, false otherwise
func (obj *account) IsVote() bool {
	return obj.vote != nil
}

// Vote returns the vote, if any
func (obj *account) Vote() signers.Vote {
	return obj.vote
}
