package references

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type contentKeyBuilder struct {
	pHash   *hash.Hash
	pKind   *uint
	content Pointer
	pCommit *hash.Hash
}

func createContentKeyBuilder() ContentKeyBuilder {
	out := contentKeyBuilder{
		pHash:   nil,
		pKind:   nil,
		content: nil,
		pCommit: nil,
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
func (app *contentKeyBuilder) WithKind(kind uint) ContentKeyBuilder {
	app.pKind = &kind
	return app
}

// WithContent adds a content to the builder
func (app *contentKeyBuilder) WithContent(content Pointer) ContentKeyBuilder {
	app.content = content
	return app
}

// WithCommit adds a commit to the builder
func (app *contentKeyBuilder) WithCommit(commit hash.Hash) ContentKeyBuilder {
	app.pCommit = &commit
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

	if app.pCommit == nil {
		return nil, errors.New("the commit is mandatory in order to build a ContentKey instance")
	}

	return createContentKey(*app.pHash, *app.pKind, app.content, *app.pCommit), nil
}
