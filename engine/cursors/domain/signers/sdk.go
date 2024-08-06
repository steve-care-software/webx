package signers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	kyber "go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
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

// NewFactory creates a Signer factory
func NewFactory() Factory {
	return createFactory()
}

// NewAdapter creates a new Signer adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// NewPublicKeyAdapter creates a new publicKey adapter
func NewPublicKeyAdapter() PublicKeyAdapter {
	return createPublicKeyAdapter()
}

// NewSignatureAdapter creates a signature adapter
func NewSignatureAdapter() SignatureAdapter {
	return createSignatureAdapter()
}

// NewVoteAdapter creates a ring signature adapter
func NewVoteAdapter() VoteAdapter {
	hashAdapter := hash.NewAdapter()
	return createVoteAdapter(hashAdapter)
}

// Factory represents a signer factory
type Factory interface {
	Create() Signer
}

// Adapter represents a signer adapter
type Adapter interface {
	ToSigner(pk string) (Signer, error)
}

// Signer represents a private key
type Signer interface {
	PublicKey() PublicKey
	Sign(msg string) (Signature, error)
	Vote(msg string, ringPubKeys []PublicKey) (Vote, error)
	String() string
}

// PublicKeyAdapter represents a publicKey adapter
type PublicKeyAdapter interface {
	ToPublicKey(pubKey string) (PublicKey, error)
}

// PublicKey represents the public key
type PublicKey interface {
	Point() kyber.Point
	Equals(pubKey PublicKey) bool
	String() string
}

// SignatureAdapter represents a signature adapter
type SignatureAdapter interface {
	ToSignature(sig string) (Signature, error)
}

// Signature represents a signature
type Signature interface {
	PublicKey(msg string) PublicKey
	Verify() bool
	String() string
}

// VoteAdapter represents a ring signature adapter
type VoteAdapter interface {
	ToSignature(sig string) (Vote, error)
	ToVerification(sig Vote, msg string, pubKeyHashes []hash.Hash) (bool, error)
}

// Vote represents a Vote
type Vote interface {
	Ring() []PublicKey
	Verify(msg string) bool
	String() string
}
