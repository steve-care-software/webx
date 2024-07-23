package stacks

import (
	"errors"
	"os"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/signers"
)

type assignableBuilder struct {
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
	signer          signers.Signer
	signerPublicKey signers.PublicKey
	signature       signers.Signature
	vote            signers.Vote
	list            Assignables
	application     applications.Application
	filePtr         *os.File
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		pBool:           nil,
		pString:         nil,
		pFloat:          nil,
		pInt:            nil,
		bytes:           nil,
		hash:            nil,
		pUnsignedInt:    nil,
		instance:        nil,
		encryptor:       nil,
		encPublicKey:    nil,
		signer:          nil,
		signerPublicKey: nil,
		signature:       nil,
		vote:            nil,
		list:            nil,
		application:     nil,
		filePtr:         nil,
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

// WithString adds a string to the builder
func (app *assignableBuilder) WithString(stringValue string) AssignableBuilder {
	app.pString = &stringValue
	return app
}

// WithFloat adds a float to the builder
func (app *assignableBuilder) WithFloat(floatVal float64) AssignableBuilder {
	app.pFloat = &floatVal
	return app
}

// WithInt adds an int to the builder
func (app *assignableBuilder) WithInt(intVal int) AssignableBuilder {
	app.pInt = &intVal
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

// WithUnsignedInt add unsigned int to the builder
func (app *assignableBuilder) WithUnsignedInt(unsignedInt uint) AssignableBuilder {
	app.pUnsignedInt = &unsignedInt
	return app
}

// WithInstance adds instanxce to the builder
func (app *assignableBuilder) WithInstance(ins instances.Instance) AssignableBuilder {
	app.instance = ins
	return app
}

// WithEncryptor adds encryptor to the builder
func (app *assignableBuilder) WithEncryptor(encryptor encryptors.Encryptor) AssignableBuilder {
	app.encryptor = encryptor
	return app
}

// WithEncryptorPubKey adds encryptor public key to the builder
func (app *assignableBuilder) WithEncryptorPubKey(encryptorPubKey encryptors.PublicKey) AssignableBuilder {
	app.encPublicKey = encryptorPubKey
	return app
}

// WithSigner adds signer to the builder
func (app *assignableBuilder) WithSigner(signer signers.Signer) AssignableBuilder {
	app.signer = signer
	return app
}

// WithSignerPubKey adds signer public key to the builder
func (app *assignableBuilder) WithSignerPubKey(signerPubKey signers.PublicKey) AssignableBuilder {
	app.signerPublicKey = signerPubKey
	return app
}

// WithSignature adds signature to the builder
func (app *assignableBuilder) WithSignature(signature signers.Signature) AssignableBuilder {
	app.signature = signature
	return app
}

// WithVote adds vote to the builder
func (app *assignableBuilder) WithVote(vote signers.Vote) AssignableBuilder {
	app.vote = vote
	return app
}

// WithList adds list to the builder
func (app *assignableBuilder) WithList(list Assignables) AssignableBuilder {
	app.list = list
	return app
}

// WithApplication adds an application to the builder
func (app *assignableBuilder) WithApplication(application applications.Application) AssignableBuilder {
	app.application = application
	return app
}

// WithFilePointer adds a file pointer to the builder
func (app *assignableBuilder) WithFilePointer(filePtr os.File) AssignableBuilder {
	app.filePtr = &filePtr
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	if app.pBool != nil {
		return createAssignableWithBool(app.pBool), nil
	}

	if app.pString != nil {
		return createAssignableWithString(app.pString), nil
	}

	if app.pFloat != nil {
		return createAssignableWithFloat(app.pFloat), nil
	}

	if app.pInt != nil {
		return createAssignableWithInt(app.pInt), nil
	}

	if app.bytes != nil {
		return createAssignableWithBytes(app.bytes), nil
	}

	if app.hash != nil {
		return createAssignableWithHash(app.hash), nil
	}

	if app.pUnsignedInt != nil {
		return createAssignableWithUnsignedInt(app.pUnsignedInt), nil
	}

	if app.instance != nil {
		return createAssignableWithInstance(app.instance), nil
	}

	if app.encryptor != nil {
		return createAssignableWithEncryptor(app.encryptor), nil
	}

	if app.encPublicKey != nil {
		return createAssignableWithEncryptorPublicKey(app.encPublicKey), nil
	}

	if app.signer != nil {
		return createAssignableWithSigner(app.signer), nil
	}

	if app.signerPublicKey != nil {
		return createAssignableWithSignerPublicKey(app.signerPublicKey), nil
	}

	if app.signature != nil {
		return createAssignableWithSignature(app.signature), nil
	}

	if app.vote != nil {
		return createAssignableWithVote(app.vote), nil
	}

	if app.list != nil {
		return createAssignableWithList(app.list), nil
	}

	if app.application != nil {
		return createAssignableWithApplication(app.application), nil
	}

	if app.filePtr != nil {
		return createAssignableWithFilePointer(app.filePtr), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
