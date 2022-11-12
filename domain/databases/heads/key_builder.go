package heads

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type keyBuilder struct {
	pHash      *hash.Hash
	pIndex     *uint
	pKind      *uint8
	content    Pointer
	isEntity   bool
	pCreatedOn *time.Time
	pTrx       *uint
}

func createKeyBuilder() KeyBuilder {
	out := keyBuilder{
		pHash:      nil,
		pIndex:     nil,
		pKind:      nil,
		content:    nil,
		isEntity:   false,
		pCreatedOn: nil,
		pTrx:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *keyBuilder) Create() KeyBuilder {
	return createKeyBuilder()
}

// WithHash adds an hash to the builder
func (app *keyBuilder) WithHash(hash hash.Hash) KeyBuilder {
	app.pHash = &hash
	return app
}

// WithIndex adds an index to the builder
func (app *keyBuilder) WithIndex(index uint) KeyBuilder {
	app.pIndex = &index
	return app
}

// WithKind adds a kind to the builder
func (app *keyBuilder) WithKind(kind uint8) KeyBuilder {
	app.pKind = &kind
	return app
}

// WithContent adds a content to the builder
func (app *keyBuilder) WithContent(content Pointer) KeyBuilder {
	app.content = content
	return app
}

// WithTransaction adds a transaction to the builder
func (app *keyBuilder) WithTransaction(trx uint) KeyBuilder {
	app.pTrx = &trx
	return app
}

// CreatedOn adds a creation time to the builder
func (app *keyBuilder) CreatedOn(createdOn time.Time) KeyBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// IsEntity flags the builder as an entity
func (app *keyBuilder) IsEntity() KeyBuilder {
	app.isEntity = true
	return app
}

// Now builds a new Key instance
func (app *keyBuilder) Now() (Key, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Key instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Key instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Key instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Key instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Key instance")
	}

	kind := *app.pKind
	if kind <= KindApplication {
		str := fmt.Sprintf("the kind must be a uint8 with a value between %d and %d, %d provided", KindBlockchain, KindApplication, kind)
		return nil, errors.New(str)
	}

	if app.pTrx != nil {
		return createKeyWithTransaction(*app.pHash, *app.pIndex, *app.pKind, app.content, app.isEntity, *app.pCreatedOn, app.pTrx), nil
	}

	if kind != KindBlockchain && kind != KindBlockchainBlock && kind != KindBlockchainTransaction {
		return nil, errors.New("the key must have a transaction when the kind is NOT blockchain related")
	}

	return createKey(*app.pHash, *app.pIndex, *app.pKind, app.content, app.isEntity, *app.pCreatedOn), nil
}
