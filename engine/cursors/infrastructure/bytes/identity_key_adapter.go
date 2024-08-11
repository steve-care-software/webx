package bytes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

type identityKeyAdapter struct {
	signerAdapter    signers.Adapter
	encryptorAdapter signers.Adapter
	builder          keys.Builder
}

func createIdentityKeyAdapter(
	signerAdapter signers.Adapter,
	encryptorAdapter signers.Adapter,
	builder keys.Builder,
) keys.Adapter {
	out := identityKeyAdapter{
		signerAdapter:    signerAdapter,
		encryptorAdapter: encryptorAdapter,
		builder:          builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *identityKeyAdapter) ToBytes(ins keys.Key) ([]byte, error) {
	return nil, nil
}

// ToInstance converts bytes to instance
func (app *identityKeyAdapter) ToInstance(data []byte) (keys.Key, error) {
	return nil, nil
}
