package validations

import "github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"

// Validation represents a validation
type Validation interface {
	Hash() hash.Hash
	Layer() hash.Hash
	Input() hash.Hash
}
