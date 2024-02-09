package layers

import "github.com/steve-care-software/identity/domain/hash"

type signer struct {
	hash             hash.Hash
	sign             string
	vote             Vote
	genSignerPubKeys uint
	hashPublicKeys   string
	voteVerify       VoteVerify
	signatureVerify  SignatureVerify
	bytes            string
	isPubKey         bool
}

func createSignerWithSign(
	hash hash.Hash,
	sign string,
) Signer {
	return createSignerInternally(
		hash,
		sign,
		nil,
		0,
		"",
		nil,
		nil,
		"",
		false,
	)
}

func createSignerWithVote(
	hash hash.Hash,
	vote Vote,
) Signer {
	return createSignerInternally(
		hash,
		"",
		vote,
		0,
		"",
		nil,
		nil,
		"",
		false,
	)
}

func createSignerWithGenerateSignerKeys(
	hash hash.Hash,
	genSignerPubKeys uint,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		genSignerPubKeys,
		"",
		nil,
		nil,
		"",
		false,
	)
}

func createSignerWithHashPublicKeys(
	hash hash.Hash,
	hashPublicKeys string,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		0,
		hashPublicKeys,
		nil,
		nil,
		"",
		false,
	)
}

func createSignerWithVoteVerify(
	hash hash.Hash,
	voteVerify VoteVerify,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		0,
		"",
		voteVerify,
		nil,
		"",
		false,
	)
}

func createSignerWithSignatureVerify(
	hash hash.Hash,
	signatureVerify SignatureVerify,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		0,
		"",
		nil,
		signatureVerify,
		"",
		false,
	)
}

func createSignerWithBytes(
	hash hash.Hash,
	bytes string,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		0,
		"",
		nil,
		nil,
		bytes,
		false,
	)
}

func createSignerWithIsPublicKey(
	hash hash.Hash,
) Signer {
	return createSignerInternally(
		hash,
		"",
		nil,
		0,
		"",
		nil,
		nil,
		"",
		true,
	)
}

func createSignerInternally(
	hash hash.Hash,
	sign string,
	vote Vote,
	genSignerPubKeys uint,
	hashPublicKeys string,
	voteVerify VoteVerify,
	signatureVerify SignatureVerify,
	bytes string,
	isPubKey bool,
) Signer {
	out := signer{
		hash:             hash,
		sign:             sign,
		vote:             vote,
		genSignerPubKeys: genSignerPubKeys,
		hashPublicKeys:   hashPublicKeys,
		voteVerify:       voteVerify,
		signatureVerify:  signatureVerify,
		bytes:            bytes,
		isPubKey:         isPubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *signer) Hash() hash.Hash {
	return obj.hash
}

// IsSign returns true if there is a sign, false otherwise
func (obj *signer) IsSign() bool {
	return obj.sign != ""
}

// Sign returns the sign, if any
func (obj *signer) Sign() string {
	return obj.sign
}

// IsVote returns true if there is a vote, false otherwise
func (obj *signer) IsVote() bool {
	return obj.vote != nil
}

// Vote returns the vote, if any
func (obj *signer) Vote() Vote {
	return obj.vote
}

// IsGenerateSignerPublicKeys returns true if there is a generateSignerPubKeys, false otherwise
func (obj *signer) IsGenerateSignerPublicKeys() bool {
	return obj.genSignerPubKeys > 0
}

// GenerateSignerPublicKeys returns true the generateSignerPubKeys, if any
func (obj *signer) GenerateSignerPublicKeys() uint {
	return obj.genSignerPubKeys
}

// IsHashPublicKeys returns true if there is an hashPublicKeys, false otherwise
func (obj *signer) IsHashPublicKeys() bool {
	return obj.hashPublicKeys != ""
}

// HashPublicKeys returns true the hashPublicKeys, if any
func (obj *signer) HashPublicKeys() string {
	return obj.hashPublicKeys
}

// IsVoteVerify returns true if there is a voteVerify, false otherwise
func (obj *signer) IsVoteVerify() bool {
	return obj.voteVerify != nil
}

// VoteVerify returns the voteVerify, if any
func (obj *signer) VoteVerify() VoteVerify {
	return obj.voteVerify
}

// IsSignatureVerify returns true if there is a signatureVerify, false otherwise
func (obj *signer) IsSignatureVerify() bool {
	return obj.signatureVerify != nil
}

// SignatureVerify returns the signatureVerify, if any
func (obj *signer) SignatureVerify() SignatureVerify {
	return obj.signatureVerify
}

// IsBytes returns true if there is a bytes, false otherwise
func (obj *signer) IsBytes() bool {
	return obj.bytes != ""
}

// Bytes returns the bytes, if any
func (obj *signer) Bytes() string {
	return obj.bytes
}

// IsPublicKey returns true if there is isPubKey, false otherwise
func (obj *signer) IsPublicKey() bool {
	return obj.isPubKey
}
