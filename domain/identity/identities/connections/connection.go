package connections

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/identities/publics"
)

type connection struct {
	id         uuid.UUID
	public     publics.Public
	encryption keys.PrivateKey
}

func createConnection(
	id uuid.UUID,
	public publics.Public,
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
func (obj *connection) Public() publics.Public {
	return obj.public
}

// Encryption returns the encryption
func (obj *connection) Encryption() keys.PrivateKey {
	return obj.encryption
}
