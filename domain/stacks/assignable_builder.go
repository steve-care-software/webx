package stacks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
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
	action          actions.Action
	commit          commits.Commit
	database        databases.Database
	delete          deletes.Delete
	modification    modifications.Modification
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
		action:          nil,
		commit:          nil,
		database:        nil,
		delete:          nil,
		modification:    nil,
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

// WithAction adds action to the builder
func (app *assignableBuilder) WithAction(action actions.Action) AssignableBuilder {
	app.action = action
	return app
}

// WithCommit adds commit to the builder
func (app *assignableBuilder) WithCommit(commit commits.Commit) AssignableBuilder {
	app.commit = commit
	return app
}

// WithDatabase adds database to the builder
func (app *assignableBuilder) WithDatabase(database databases.Database) AssignableBuilder {
	app.database = database
	return app
}

// WithDelete adds delete to the builder
func (app *assignableBuilder) WithDelete(delete deletes.Delete) AssignableBuilder {
	app.delete = delete
	return app
}

// WithModification adds modification to the builder
func (app *assignableBuilder) WithModification(modification modifications.Modification) AssignableBuilder {
	app.modification = modification
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

	if app.action != nil {
		return createAssignableWithAction(app.action), nil
	}

	if app.commit != nil {
		return createAssignableWithCommit(app.commit), nil
	}

	if app.database != nil {
		return createAssignableWithDatabase(app.database), nil
	}

	if app.delete != nil {
		return createAssignableWithDelete(app.delete), nil
	}

	if app.modification != nil {
		return createAssignableWithModification(app.modification), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
