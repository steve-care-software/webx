package spends

import (
	"math/big"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/currencies/domain/shares"
	"github.com/steve-care-software/webx/identities/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

// Spends represents spends
type Spends interface {
	Hash() hash.Hash
	List() []Spend
}

// Spend represents a spend currency action
type Spend interface {
	Hash() hash.Hash
	Content() Content
	Signature() signatures.RingSignature
}

// Content represents content
type Content interface {
	Hash() hash.Hash
	Parent() Parent
	Amount() big.Int
	Safes() Safes
	Condition() programs.Program
}

// Safes represents safes
type Safes interface {
	List() []Safe
}

// Safe represents a safe
type Safe interface {
	Amount() big.Int
	Lock() []hash.Hash
}

// Parent represents the parent unit to spend
type Parent interface {
	Hash() hash.Hash
	IsShare() bool
	Share() shares.Share
	IsSpends() bool
	Spends() Spends
}
