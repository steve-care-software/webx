package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/signers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the key adapter
type Adapter interface {
	ToBytes(ins Key) ([]byte, error)
	ToInstance(data []byte) (Key, error)
}

// Builder represents a key builder
type Builder interface {
	Create() Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Key, error)
}

// Key represents a key
type Key interface {
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
	CreatedOn() time.Time
}
