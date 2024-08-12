package identities

const noCurrentUserErr = "there is no authenticated current user"

// Application represents the identity application
type Application interface {
	/*List(input loaders_identities.Identity) []string
	Create(input loaders_identities.Identity, name string, description string, password []byte) (loaders_identities.Identity, error)
	Authenticate(input loaders_identities.Identity, name string, password []byte) (loaders_identities.Identity, error)
	SetPassword(input loaders_identities.Identity, newPassword []byte) (loaders_identities.Identity, error) // update the password of the authenticated user
	SetUser(input loaders_identities.Identity, name string) (loaders_identities.Identity, error)
	Follow(input loaders_identities.Identity, namespace string, password []byte) (loaders_identities.Identity, error)

	// actions:
	Encrypt(input loaders_identities.Identity, message []byte) ([]byte, error)
	Decrypt(input loaders_identities.Identity, cipher []byte) ([]byte, error)
	Sign(input loaders_identities.Identity, message []byte) (signers.Signature, error)
	ValidateSignature(input loaders_identities.Identity, message []byte, sig signers.Signature) (bool, error)
	Vote(input loaders_identities.Identity, message []byte, ring []signers.PublicKey) (signers.Vote, error)
	ValidateVote(input loaders_identities.Identity, message []byte, vote signers.Vote, ring []hash.Hash) (bool, error)

	// namespaces
	Namespaces(input loaders_identities.Identity) ([]string, error)
	Namespace(input loaders_identities.Identity, name string) error
	Dive(input loaders_identities.Identity) (namespaces.Application, error)*/
}
