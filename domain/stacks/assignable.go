package stacks

import (
	"github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/hash"
)

type assignable struct {
	pBool              *bool
	bytes              []byte
	encryptorPublicKey encryptors.PublicKey
	signerPubKey       signers.PublicKey
	signerPubKeys      []signers.PublicKey
	signature          signers.Signature
	vote               signers.Vote
	hashList           []hash.Hash
	hash               hash.Hash
}

func createAssignableWithBool(
	pBool *bool,
) Assignable {
	return createAssignableInternally(
		pBool,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithBytes(
	bytes []byte,
) Assignable {
	return createAssignableInternally(
		nil,
		bytes,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithEncryptorPublicKey(
	encryptorPublicKey encryptors.PublicKey,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		encryptorPublicKey,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithSignerPublicKey(
	signerPubKey signers.PublicKey,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		signerPubKey,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithSignerPublicKeys(
	signerPubKeys []signers.PublicKey,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		signerPubKeys,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithSignature(
	signature signers.Signature,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		signature,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithVote(
	vote signers.Vote,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		vote,
		nil,
		nil,
	)
}

func createAssignableWithHashList(
	hashList []hash.Hash,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		hashList,
		nil,
	)
}

func createAssignableWithHash(
	hash hash.Hash,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		hash,
	)
}

func createAssignableInternally(
	pBool *bool,
	bytes []byte,
	encryptorPublicKey encryptors.PublicKey,
	signerPubKey signers.PublicKey,
	signerPubKeys []signers.PublicKey,
	signature signers.Signature,
	vote signers.Vote,
	hashList []hash.Hash,
	hash hash.Hash,
) Assignable {
	out := assignable{
		pBool:              pBool,
		bytes:              bytes,
		encryptorPublicKey: encryptorPublicKey,
		signerPubKey:       signerPubKey,
		signerPubKeys:      signerPubKeys,
		signature:          signature,
		vote:               vote,
		hashList:           hashList,
		hash:               hash,
	}

	return &out
}

// IsBool returns true if bool, false otherwise
func (obj *assignable) IsBool() bool {
	return obj.pBool != nil
}

// Bool returns bool, if any
func (obj *assignable) Bool() *bool {
	return obj.pBool
}

// IsBytes returns true if bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns bytes, if any
func (obj *assignable) Bytes() []byte {
	return obj.bytes
}

// IsEncryptorPublicKey returns true if encryptor public key, false otherwise
func (obj *assignable) IsEncryptorPublicKey() bool {
	return obj.encryptorPublicKey != nil
}

// EncryptorPublicKey returns encryptor public key, if any
func (obj *assignable) EncryptorPublicKey() encryptors.PublicKey {
	return obj.encryptorPublicKey
}

// IsSignerPublicKey returns true if signer public key, false otherwise
func (obj *assignable) IsSignerPublicKey() bool {
	return obj.signerPubKey != nil
}

// SignerPublicKey returns signer public key, if any
func (obj *assignable) SignerPublicKey() signers.PublicKey {
	return obj.signerPubKey
}

// IsSignerPublicKeys returns true if signer public keys, false otherwise
func (obj *assignable) IsSignerPublicKeys() bool {
	return obj.signerPubKeys != nil
}

// SignerPublicKeys returns signer public keys, if any
func (obj *assignable) SignerPublicKeys() []signers.PublicKey {
	return obj.signerPubKeys
}

// IsSignature returns true if signatures, false otherwise
func (obj *assignable) IsSignature() bool {
	return obj.signature != nil
}

// Signature returns signature, if any
func (obj *assignable) Signature() signers.Signature {
	return obj.signature
}

// IsVote returns true if vote, false otherwise
func (obj *assignable) IsVote() bool {
	return obj.vote != nil
}

// Vote returns vote, if any
func (obj *assignable) Vote() signers.Vote {
	return obj.vote
}

// IsHashList returns true if hashList, false otherwise
func (obj *assignable) IsHashList() bool {
	return obj.hashList != nil
}

// HashList returns hashList, if any
func (obj *assignable) HashList() []hash.Hash {
	return obj.hashList
}

// IsHash returns true if hash, false otherwise
func (obj *assignable) IsHash() bool {
	return obj.hash != nil
}

// Hash returns hash, if any
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}
