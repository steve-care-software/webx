package references

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type contentKeyBuilder struct {
	pHash      *hash.Hash
	pKind      *uint8
	content    Pointer
	pTrx       *hash.Hash
	pCreatedOn *time.Time
}

func createContentKeyBuilder() ContentKeyBuilder {
	out := contentKeyBuilder{
		pHash:      nil,
		pKind:      nil,
		content:    nil,
		pTrx:       nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentKeyBuilder) Create() ContentKeyBuilder {
	return createContentKeyBuilder()
}

// WithHash adds an hash to the builder
func (app *contentKeyBuilder) WithHash(hash hash.Hash) ContentKeyBuilder {
	app.pHash = &hash
	return app
}

// WithKind adds a kind to the builder
func (app *contentKeyBuilder) WithKind(kind uint8) ContentKeyBuilder {
	app.pKind = &kind
	return app
}

// WithContent adds a content to the builder
func (app *contentKeyBuilder) WithContent(content Pointer) ContentKeyBuilder {
	app.content = content
	return app
}

// WithTransaction adds a transaction to the builder
func (app *contentKeyBuilder) WithTransaction(trx hash.Hash) ContentKeyBuilder {
	app.pTrx = &trx
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentKeyBuilder) CreatedOn(createdOn time.Time) ContentKeyBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new ContentKey instance
func (app *contentKeyBuilder) Now() (ContentKey, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a ContentKey instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a ContentKey instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a ContentKey instance")
	}

	if app.pTrx == nil {
		return nil, errors.New("the transaction is mandatory in order to build a ContentKey instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a ContentKey instance")
	}

	return createContentKey(*app.pHash, *app.pKind, app.content, *app.pTrx, *app.pCreatedOn), nil
}
