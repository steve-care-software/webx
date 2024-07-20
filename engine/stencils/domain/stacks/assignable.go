package stacks

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/signers"
)

type assignable struct {
	pBool           *bool
	pString         *string
	pFloat          *float64
	pInt            *int
	bytes           []byte
	hash            hash.Hash
	pUnsignedInt    *uint
	instance        instances.Instance
	encryptor       encryptors.Encryptor
	encPublicKey    encryptors.PublicKey
	signature       signers.Signature
	signer          signers.Signer
	signerPublicKey signers.PublicKey
	vote            signers.Vote
	list            Assignables
	application     applications.Application
}

func createAssignableWithBool(
	pBool *bool,
) Assignable {
	return createAssignableInternally(pBool, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithString(
	pString *string,
) Assignable {
	return createAssignableInternally(nil, pString, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithFloat(
	pFloat *float64,
) Assignable {
	return createAssignableInternally(nil, nil, pFloat, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithInt(
	pInt *int,
) Assignable {
	return createAssignableInternally(nil, nil, nil, pInt, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithBytes(
	bytes []byte,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, bytes, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithHash(
	hash hash.Hash,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, hash, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithUnsignedInt(
	pUnsignedInt *uint,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, pUnsignedInt, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithInstance(
	instance instances.Instance,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, instance, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithEncryptor(
	encryptor encryptors.Encryptor,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, encryptor, nil, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithEncryptorPublicKey(
	encPublicKey encryptors.PublicKey,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, encPublicKey, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithSigner(
	signer signers.Signer,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, signer, nil, nil, nil, nil, nil)
}

func createAssignableWithSignerPublicKey(
	signerPubKey signers.PublicKey,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, signerPubKey, nil, nil, nil, nil)
}

func createAssignableWithSignature(
	signature signers.Signature,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, signature, nil, nil, nil)
}

func createAssignableWithVote(
	vote signers.Vote,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, vote, nil, nil)
}

func createAssignableWithList(
	list Assignables,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, list, nil)
}

func createAssignableWithApplication(
	application applications.Application,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, application)
}

func createAssignableInternally(
	pBool *bool,
	pString *string,
	pFloat *float64,
	pInt *int,
	bytes []byte,
	hash hash.Hash,
	pUnsignedInt *uint,
	instance instances.Instance,
	encryptor encryptors.Encryptor,
	encPublicKey encryptors.PublicKey,
	signer signers.Signer,
	signerPublicKey signers.PublicKey,
	signature signers.Signature,
	vote signers.Vote,
	list Assignables,
	application applications.Application,
) Assignable {
	out := assignable{
		pBool:           pBool,
		pString:         pString,
		pFloat:          pFloat,
		pInt:            pInt,
		bytes:           bytes,
		hash:            hash,
		pUnsignedInt:    pUnsignedInt,
		instance:        instance,
		encryptor:       encryptor,
		encPublicKey:    encPublicKey,
		signer:          signer,
		signerPublicKey: signerPublicKey,
		signature:       signature,
		vote:            vote,
		list:            list,
		application:     application,
	}

	return &out
}

// IsBool returns true if bool, false otherwose
func (obj *assignable) IsBool() bool {
	return obj.pBool != nil
}

// Bool returns the bool, if any
func (obj *assignable) Bool() *bool {
	return obj.pBool
}

// IsString returns true if string, false otherwise
func (obj *assignable) IsString() bool {
	return obj.pString != nil
}

// String returns the string, if any
func (obj *assignable) String() *string {
	return obj.pString
}

// IsFloat returns true if float, false otherwise
func (obj *assignable) IsFloat() bool {
	return obj.pFloat != nil
}

// Float returns the float, if any
func (obj *assignable) Float() *float64 {
	return obj.pFloat
}

// IsInt returns true if int, false otherwise
func (obj *assignable) IsInt() bool {
	return obj.pInt != nil
}

// Int returns the int, if any
func (obj *assignable) Int() *int {
	return obj.pInt
}

// IsBytes returns true if bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() []byte {
	return obj.bytes
}

// IsHash returns true if hash, false otherwise
func (obj *assignable) IsHash() bool {
	return obj.hash != nil
}

// Hash returns the hash, if any
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsUnsignedInt returns true if unsigned int, false otherwise
func (obj *assignable) IsUnsignedInt() bool {
	return obj.pUnsignedInt != nil
}

// UnsignedInt returns the unsigned int, if any
func (obj *assignable) UnsignedInt() *uint {
	return obj.pUnsignedInt
}

// IsInstance returns true if instance, false otherwise
func (obj *assignable) IsInstance() bool {
	return obj.instance != nil
}

// Instance returns the instance, if any
func (obj *assignable) Instance() instances.Instance {
	return obj.instance
}

// IsEncryptor returns true if encryptor, false otherwise
func (obj *assignable) IsEncryptor() bool {
	return obj.encryptor != nil
}

// Encryptor returns the encryptor, if any
func (obj *assignable) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// IsEncryptorPublicKey returns true if encryptor public key, false otherwise
func (obj *assignable) IsEncryptorPublicKey() bool {
	return obj.encPublicKey != nil
}

// EncryptorPublicKey returns the encryptor public key, if any
func (obj *assignable) EncryptorPublicKey() encryptors.PublicKey {
	return obj.encPublicKey
}

// IsSigner returns true if signer, false otherwise
func (obj *assignable) IsSigner() bool {
	return obj.signer != nil
}

// Signer returns the signer, if any
func (obj *assignable) Signer() signers.Signer {
	return obj.signer
}

// IsSignerPublicKey returns true if signer public key, false otherwise
func (obj *assignable) IsSignerPublicKey() bool {
	return obj.signerPublicKey != nil
}

// SignerPublicKey returns the signer public key, if any
func (obj *assignable) SignerPublicKey() signers.PublicKey {
	return obj.signerPublicKey
}

// IsSignature returns true if signature, false otherwise
func (obj *assignable) IsSignature() bool {
	return obj.signature != nil
}

// Signature returns the signature, if any
func (obj *assignable) Signature() signers.Signature {
	return obj.signature
}

// IsVote returns true if vote, false otherwise
func (obj *assignable) IsVote() bool {
	return obj.vote != nil
}

// Vote returns the vote, if any
func (obj *assignable) Vote() signers.Vote {
	return obj.vote
}

// IsList returns true if list, false otherwise
func (obj *assignable) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *assignable) List() Assignables {
	return obj.list
}

// IsApplication returns true if application, false otherwise
func (obj *assignable) IsApplication() bool {
	return obj.application != nil
}

// Applicattion returns the application, if any
func (obj *assignable) Applicattion() applications.Application {
	return obj.application
}
