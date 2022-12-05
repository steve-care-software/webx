package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/identities/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/identities/domain/identities"
)

// Application represents a program application
type Application interface {
	New(name string) error
	applications.Database
	Database
	Software
}

// Software represents the program software application
type Software interface {
	Sign(identity identities.Identity, hash hash.Hash) (signatures.RingSignature, error)
}

// Database represents the program database application
type Database interface {
	List() ([]string, error)
	Retrieve(context uint, name string, password []byte) (identities.Identity, error)
	Insert(context uint, identity identities.Identity, password []byte) error
	Update(context uint, identity identities.Identity, currentPassword []byte, newPassword []byte) error
}
