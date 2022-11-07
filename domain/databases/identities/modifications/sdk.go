package modifications

import (
	"time"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithName(name string) Builder
	WithSignature(sig []byte) Builder
	WithEncryption(enc []byte) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Modification, error)
}

// Modification represents a modifucation
type Modification interface {
	Entity() entities.Entity
	Content() Content
	CreatedOn() time.Time
}

// Content represents a modification content
type Content interface {
	HasName() bool
	Name() string
	HasSignature() bool
	Signature() []byte
	HasEncryption() bool
	Encryption() []byte
}
