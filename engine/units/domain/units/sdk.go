package units

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

// Units represents units
type Units interface {
	Hash() hash.Hash
	List() []Unit
}

// Unit represents a unit
type Unit interface {
	Hash() hash.Hash
	Amount() hash.Hash
}
