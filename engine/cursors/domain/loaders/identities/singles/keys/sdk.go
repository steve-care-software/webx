package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/signers"
)

// Adapter represents the keys adapter
type Adapter interface {
	InstancesToBytes(ins Keys) ([]byte, error)
	BytesToInstances(data []byte) (Keys, error)
	InstanceToBytes(ins Key) ([]byte, error)
	BytesToInstance(data []byte) (Key, error)
}

// Builder represents the keys builder
type Builder interface {
	Create() Builder
	WithList(list []Key) Builder
	Now() (Keys, error)
}

// Keys represents keys
type Keys interface {
	List() []Key
}

// KeyBuilder represents a key builder
type KeyBuilder interface {
	Create() KeyBuilder
	WithName(name string) KeyBuilder
	WithEncryptor(encryptor encryptors.Encryptor) KeyBuilder
	WithSigner(signer signers.Signer) KeyBuilder
	CreatedOn() time.Time
	Now() (Key, error)
}

// Key represents a key
type Key interface {
	Name() string
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
	CreatedOn() time.Time
}
