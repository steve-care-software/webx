package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	loaders_identities "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

// Application represents the identity application
type Application interface {
	Create(input loaders_identities.Identity, name string, description string, password []byte) (loaders_identities.Identity, error)
	Authenticate(input loaders_identities.Identity, name string, password []byte) (loaders_identities.Identity, error)
	SetPassword(input loaders_identities.Identity, newPassword []byte) (loaders_identities.Identity, error) // update the password of the authenticated user
	SetUser(input loaders_identities.Identity, name string) (loaders_identities.Identity, error)
	Follow(input loaders_identities.Identity, namespace string) (loaders_identities.Identity, error)

	// actions:
	Encrypt(input loaders_identities.Identity, message []byte) ([]byte, error)
	Decrypt(input loaders_identities.Identity, cipher []byte) ([]byte, error)
	Sign(input loaders_identities.Identity, message []byte) (signers.Signature, error)
	ValidateSignature(input loaders_identities.Identity, message []byte, sig signers.Signature) (bool, error)
	Vote(input loaders_identities.Identity, message []byte, ring []signers.PublicKey) (signers.Vote, error)
	ValidateVote(input loaders_identities.Identity, message []byte, vote signers.Vote, ring []hash.Hash) (bool, error)
}
