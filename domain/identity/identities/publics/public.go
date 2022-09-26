package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type public struct {
	id         uuid.UUID
	name       string
	encryption keys.PublicKey
	signature  hash.Hash
	host       string
	port       uint
}

func createPublic(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
) Public {
	out := public{
		id:         id,
		name:       name,
		encryption: encryption,
		signature:  signature,
		host:       host,
		port:       port,
	}

	return &out
}

// ID returns the id
func (obj *public) ID() uuid.UUID {
	return obj.id
}

// Name returns the name
func (obj *public) Name() string {
	return obj.name
}

// Encryption returns the encryption
func (obj *public) Encryption() keys.PublicKey {
	return obj.encryption
}

// Encryption returns the encryption
func (obj *public) Signature() hash.Hash {
	return obj.signature
}

// Host returns the host
func (obj *public) Host() string {
	return obj.host
}

// Port returns the port
func (obj *public) Port() uint {
	return obj.port
}
