package governances

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/governances/shares"
	"github.com/steve-care-software/webx/engine/units/domain/stocks"
)

// Governance represents the governance
type Governance interface {
	Hash() hash.Hash
	Developers() stocks.Stocks
	Shares() shares.Share
}
