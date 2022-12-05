package references

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type blockchainKeyBuilder struct {
	pHash      *hash.Hash
	content    Pointer
	pCreatedOn *time.Time
}

func createBlockchainKeyBuilder() BlockchainKeyBuilder {
	out := blockchainKeyBuilder{
		pHash:      nil,
		content:    nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockchainKeyBuilder) Create() BlockchainKeyBuilder {
	return createBlockchainKeyBuilder()
}

// WithHash adds an hash to the builder
func (app *blockchainKeyBuilder) WithHash(hash hash.Hash) BlockchainKeyBuilder {
	app.pHash = &hash
	return app
}

// WithContent adds a content to the builder
func (app *blockchainKeyBuilder) WithContent(content Pointer) BlockchainKeyBuilder {
	app.content = content
	return app
}

// CreatedOn adds a creation time to the builder
func (app *blockchainKeyBuilder) CreatedOn(createdOn time.Time) BlockchainKeyBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new BlockchainKey instance
func (app *blockchainKeyBuilder) Now() (BlockchainKey, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a BlockchainKey instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a BlockchainKey instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a BlockchainKey instance")
	}

	return createBlockchainKey(*app.pHash, app.content, *app.pCreatedOn), nil
}
