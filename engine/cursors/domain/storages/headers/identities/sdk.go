package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Identities represents identities
type Identities interface {
	List() []Identity
}

// Identity represents an identity
type Identity interface {
	Original() originals.Original
	Signer() []byte
	Encryptor() []byte
	HasWallets() bool
	Wallets() delimiters.Delimiter
}
