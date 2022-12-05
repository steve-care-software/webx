package genesis

import (
	"math/big"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/currencies/domain/shares"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

// Genesis represents a currency genesis
type Genesis interface {
	Hash() hash.Hash
	Name() string
	Shares() shares.Shares
	TotalSupply() big.Int
	Vesting() VestingPeriod
}

// VestingPeriod represents a currency's vesting period
type VestingPeriod interface {
	Hash() hash.Hash
	Program() programs.Program
	TimeInputVariable() string
	AmountOutputVariable() string
}
