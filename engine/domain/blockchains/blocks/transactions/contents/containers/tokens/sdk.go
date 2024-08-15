package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents/containers/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Elements() elements.Elements
	Owner() []hash.Hash
}
