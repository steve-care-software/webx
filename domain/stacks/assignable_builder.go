package stacks

import (
	"errors"

	"github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/hash"
)

type assignableBuilder struct {
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

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		pBool:              nil,
		bytes:              nil,
		encryptorPublicKey: nil,
		signerPubKey:       nil,
		signerPubKeys:      nil,
		signature:          nil,
		vote:               nil,
		hashList:           nil,
		hash:               nil,
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

// WithEncryptorPublicKey adds an encryptorPublicKey to the builder
func (app *assignableBuilder) WithEncryptorPublicKey(encryptorPublicKey encryptors.PublicKey) AssignableBuilder {
	app.encryptorPublicKey = encryptorPublicKey
	return app
}

// WithSignerPublicKey add signerPublicKey to the builder
func (app *assignableBuilder) WithSignerPublicKey(signerPublicKey signers.PublicKey) AssignableBuilder {
	app.signerPubKey = signerPublicKey
	return app
}

// WithSignerPublicKeys add signerPubKeys to the builder
func (app *assignableBuilder) WithSignerPublicKeys(signerPubKeys []signers.PublicKey) AssignableBuilder {
	app.signerPubKeys = signerPubKeys
	return app
}

// WithSignature add signature to the builder
func (app *assignableBuilder) WithSignature(signature signers.Signature) AssignableBuilder {
	app.signature = signature
	return app
}

// WithVote add vote to the builder
func (app *assignableBuilder) WithVote(vote signers.Vote) AssignableBuilder {
	app.vote = vote
	return app
}

// WithHashList add hashList to the builder
func (app *assignableBuilder) WithHashList(hashList []hash.Hash) AssignableBuilder {
	app.hashList = hashList
	return app
}

// WithHash add hash to the builder
func (app *assignableBuilder) WithHash(hash hash.Hash) AssignableBuilder {
	app.hash = hash
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

	if app.encryptorPublicKey != nil {
		return createAssignableWithEncryptorPublicKey(app.encryptorPublicKey), nil
	}

	if app.signerPubKey != nil {
		return createAssignableWithSignerPublicKey(app.signerPubKey), nil
	}

	if app.signerPubKeys != nil && len(app.signerPubKeys) <= 0 {
		app.signerPubKeys = nil
	}

	if app.signerPubKeys != nil {
		return createAssignableWithSignerPublicKeys(app.signerPubKeys), nil
	}

	if app.signature != nil {
		return createAssignableWithSignature(app.signature), nil
	}

	if app.vote != nil {
		return createAssignableWithVote(app.vote), nil
	}

	if app.hashList != nil && len(app.hashList) <= 0 {
		app.hashList = nil
	}

	if app.hashList != nil {
		return createAssignableWithHashList(app.hashList), nil
	}

	if app.hash != nil {
		return createAssignableWithHash(app.hash), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
