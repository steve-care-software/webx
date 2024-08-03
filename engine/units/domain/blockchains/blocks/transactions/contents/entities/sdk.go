package entities

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents/authorizations"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents/entities/contents"
)

// Entity represents an entity
type Entity interface {
	Hash() hash.Hash
	Content() contents.Content
	Authorization() authorizations.Authorization
}
