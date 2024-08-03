package transfers

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/units"
	"github.com/steve-care-software/webx/engine/units/domain/units/transfers/tokens"
)

// Transfers represents a transfr
type Transfers interface {
	Hash() hash.Hash
	List() []Transfer
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	From() units.Unit
	Token() tokens.Token
}
