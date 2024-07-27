package contents

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/entities"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents/singles"
)

// Content represents a transaction content
type Content interface {
	Hash() hash.Hash
	IsEntity() bool
	Entity() entities.Entity
	IsSingle() bool
	Single() singles.Single
}
