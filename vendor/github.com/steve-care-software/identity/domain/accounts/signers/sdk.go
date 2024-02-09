package signers

import (
	kyber "go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"github.com/steve-care-software/identity/domain/hash"
)

/*
 * H'(m, s, e) = H(m || s * G + e * P)
 * P = x * G
 * e = H(m || k * G)
 * k = s + e * x
 * s = k – e * x
 * k = H(m || x) -> to generate a new k, since nobody but us knows x
 * where ...
 * 1. H is a hash function, for instance SHA256.
 * 2. s and e are 2 numbers forming the ring signature
 * 3. s and r are a pubKey and a number forming a signature
 * 4. m is the message we want to sign
 * 5. P is the public key.
 * 6. G is the random base
 * 7. k is a number chosen randomly.  A new one every time we sign must be generated
 * 8. x is the private key
 */

const delimiter = "#"
const elementDelimiter = "|"

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewFactory creates a new factory
func NewFactory() Factory {
	return createFactory()
}

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// NewPublicKeyAdapter creates a new public key adapter
func NewPublicKeyAdapter() PublicKeyAdapter {
	return createPublicKeyAdapter()
}

// NewSignatureAdapter creates a new signature adapter
func NewSignatureAdapter() SignatureAdapter {
	return createSignatureAdapter()
}

// NewVoteAdapter creates a new vote adapter
func NewVoteAdapter() VoteAdapter {
	hashAdapter := hash.NewAdapter()
	return createVoteAdapter(
		hashAdapter,
	)
}

// Factory represents a signer factory
type Factory interface {
	Create() Signer
}

// Adapter represents a signer adapter
type Adapter interface {
	ToSigner(pk []byte) (Signer, error)
}

// Signer represents a signer
type Signer interface {
	PublicKey() PublicKey
	Sign(msg []byte) (Signature, error)
	Vote(msg []byte, ring []PublicKey) (Vote, error)
	Bytes() ([]byte, error)
}

// PublicKeyAdapter represents a publicKey adapter
type PublicKeyAdapter interface {
	ToPublicKey(pubKey []byte) (PublicKey, error)
}

// PublicKey represents the public key
type PublicKey interface {
	Point() kyber.Point
	Equals(pubKey PublicKey) bool
	Bytes() ([]byte, error)
}

// SignatureAdapter represents a signature adapter
type SignatureAdapter interface {
	ToSignature(bytes []byte) (Signature, error)
}

// Signature represents a signature
type Signature interface {
	PublicKey(msg []byte) (PublicKey, error)
	Verify() bool
	Bytes() ([]byte, error)
}

// VoteAdapter represents a vote adapter
type VoteAdapter interface {
	ToVote(bytes []byte) (Vote, error)
	ToVerification(vote Vote, msg []byte, pubKeyHashes []hash.Hash) (bool, error)
}

// Vote represents a vote
type Vote interface {
	Ring() []PublicKey
	Verify(msg []byte) bool
	Bytes() ([]byte, error)
}
