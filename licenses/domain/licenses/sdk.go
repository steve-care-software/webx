package licenses

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/currencies/domain/spends"
	"github.com/steve-care-software/webx/licenses/domain/products"
)

// Licenses represents licenses
type Licenses interface {
	List() []License
}

// License represents a license
type License interface {
	Hash() hash.Hash
	Product() products.Product
	Payment() spends.Spend
	Owner() []hash.Hash
}
