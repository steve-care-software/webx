package tokens

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/units"
	"github.com/steve-care-software/webx/engine/units/domain/units/transfers/tokens/validations"
)

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Amount() uint64
	Answer() []byte
	Remaining() units.Units
	HasValidation() bool
	Validation() validations.Validation
}
