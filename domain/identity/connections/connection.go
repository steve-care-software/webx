package connections

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
)

type connection struct {
	id         uuid.UUID
	public     uuid.UUID
	encryption keys.PrivateKey
}

func createConnection(
	id uuid.UUID,
	public uuid.UUID,
	encryption keys.PrivateKey,
) Connection {
	return createConnectionInternally(id, public, encryption)
}

func createConnectionInternally(
	id uuid.UUID,
	public uuid.UUID,
	encryption keys.PrivateKey,
) Connection {
	out := connection{
		id:         id,
		public:     public,
		encryption: encryption,
	}

	return &out
}

// ID returns the id
func (obj *connection) ID() uuid.UUID {
	return obj.id
}

// Public returns the public
func (obj *connection) Public() uuid.UUID {
	return obj.public
}

// Encryption returns the encryption
func (obj *connection) Encryption() keys.PrivateKey {
	return obj.encryption
}
