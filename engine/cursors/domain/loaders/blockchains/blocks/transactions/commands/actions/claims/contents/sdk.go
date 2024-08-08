package contents

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Content represents a claim content
type Content interface {
	Hash() hash.Hash
	Transfer() hash.Hash
	Answer() []byte
	Condition() hash.Hash
	ConditionExpireIn() uint64
}
