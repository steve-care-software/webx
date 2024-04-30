package stacks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
)

type assignableBuilder struct {
	pBool           *bool
	bytes           []byte
	hash            hash.Hash
	hashList        []hash.Hash
	stringList      []string
	pUnsignedInt    *uint
	instance        instances.Instance
	encryptor       encryptors.Encryptor
	encryptorPubKey encryptors.PublicKey
	signer          signers.Signer
	signerPubKey    signers.PublicKey
	signature       signers.Signature
	vote            signers.Vote
	list            Assignables
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		pBool:           nil,
		bytes:           nil,
		hash:            nil,
		hashList:        nil,
		stringList:      nil,
		pUnsignedInt:    nil,
		instance:        nil,
		encryptor:       nil,
		encryptorPubKey: nil,
		signer:          nil,
		signerPubKey:    nil,
		signature:       nil,
		vote:            nil,
		list:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithBool adds a bool to the builder
func (app *assignableBuilder) WithBool(boolValue bool) AssignableBuilder {
	app.pBool = &boolValue
	return app
}

// WithBytes add bytes to the builder
func (app *assignableBuilder) WithBytes(bytes []byte) AssignableBuilder {
	app.bytes = bytes
	return app
}

// WithHash add hash to the builder
func (app *assignableBuilder) WithHash(hash hash.Hash) AssignableBuilder {
	app.hash = hash
	return app
}

// WithHashList adds an hash list to the builder
func (app *assignableBuilder) WithHashList(hashList []hash.Hash) AssignableBuilder {
	app.hashList = hashList
	return app
}

// WithStringList adds a string list to the builder
func (app *assignableBuilder) WithStringList(strList []string) AssignableBuilder {
	app.stringList = strList
	return app
}

// WithUnsignedInt adds an uint to the builder
func (app *assignableBuilder) WithUnsignedInt(unsignedInt uint) AssignableBuilder {
	app.pUnsignedInt = &unsignedInt
	return app
}

// WithInstance adds an instance to the builder
func (app *assignableBuilder) WithInstance(instance instances.Instance) AssignableBuilder {
	app.instance = instance
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *assignableBuilder) WithEncryptor(encryptor encryptors.Encryptor) AssignableBuilder {
	app.encryptor = encryptor
	return app
}

// WithEncryptorPubKey adds an encryptor pubKey to the builder
func (app *assignableBuilder) WithEncryptorPubKey(encryptorPubKey encryptors.PublicKey) AssignableBuilder {
	app.encryptorPubKey = encryptorPubKey
	return app
}

// WithSigner adds a signer pk to the builder
func (app *assignableBuilder) WithSigner(signer signers.Signer) AssignableBuilder {
	app.signer = signer
	return app
}

// WithSignerPubKey adds a signer pubKey to the builder
func (app *assignableBuilder) WithSignerPubKey(signerPubKey signers.PublicKey) AssignableBuilder {
	app.signerPubKey = signerPubKey
	return app
}

// WithSignature adds a signature to the builder
func (app *assignableBuilder) WithSignature(signature signers.Signature) AssignableBuilder {
	app.signature = signature
	return app
}

// WithVote adds a vote to the builder
func (app *assignableBuilder) WithVote(vote signers.Vote) AssignableBuilder {
	app.vote = vote
	return app
}

// WithList adds a list to the builder
func (app *assignableBuilder) WithList(list Assignables) AssignableBuilder {
	app.list = list
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	if app.pBool != nil {
		return createAssignableWithBool(app.pBool), nil
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createAssignableWithBytes(app.bytes), nil
	}

	if app.hash != nil {
		return createAssignableWithHash(app.hash), nil
	}

	if app.hashList != nil && len(app.hashList) <= 0 {
		app.hashList = nil
	}

	if app.hashList != nil {
		return createAssignableWithHashList(app.hashList), nil
	}

	if app.stringList != nil {
		return createAssignableWithStringList(app.stringList), nil
	}

	if app.pUnsignedInt != nil {
		return createAssignableWithUnsignedInt(app.pUnsignedInt), nil
	}

	if app.instance != nil {
		return createAssignableWithInstance(app.instance), nil
	}

	if app.encryptor != nil {

	}

	if app.encryptorPubKey != nil {

	}

	if app.signer != nil {

	}

	if app.signerPubKey != nil {

	}

	if app.signature != nil {

	}

	if app.vote != nil {

	}

	if app.list != nil {

	}

	return nil, errors.New("the Assignable is invalid")
}
