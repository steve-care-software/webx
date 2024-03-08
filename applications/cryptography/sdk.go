package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Application represents the cryptography application
type Application interface {
	Hash(bytes []byte) (hash.Hash, error)
	Encrypt(message []byte, password []byte) ([]byte, error)
	Decrypt(cipher []byte, password []byte) ([]byte, error)
}
