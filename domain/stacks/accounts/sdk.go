package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	WithCredentials(credentials credentials.Credentials) Builder
	WithRing(ring []signers.PublicKey) Builder
	WithSignature(sig signers.Signature) Builder
	WithVote(vote signers.Vote) Builder
	Now() (Account, error)
}

// Account represents an account assignable
type Account interface {
	IsAccount() bool
	Account() accounts.Account
	IsCredentials() bool
	Credentials() credentials.Credentials
	IsRing() bool
	Rign() []signers.PublicKey
	IsSignature() bool
	Signature() signers.Signature
	IsVote() bool
	Vote() signers.Vote
}
