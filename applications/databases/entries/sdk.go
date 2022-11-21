package entries

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entries"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	basePath string,
) Builder {
	builder := entries.NewBuilder()
	entryBuilder := entries.NewEntryBuilder()
	relationsBuilder := entries.NewRelationsBuilder()
	relationBuilder := entries.NewRelationBuilder()
	linksBuilder := entries.NewLinksBuilder()
	linkBuilder := entries.NewLinkBuilder()
	additionBuilder := entries.NewAdditionBuilder()
	return createBuilder(
		builder,
		entryBuilder,
		relationsBuilder,
		relationBuilder,
		linksBuilder,
		linkBuilder,
		additionBuilder,
		basePath,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithRelativeFilePath(relFilePath string) Builder
	Now() (Application, error)
}

// Application represents the pendings application
type Application interface {
	List(kind uint8) ([][]byte, error)
	ListByLink(masterKind uint8, masterHash hash.Hash, slaveKind uint8) ([][]byte, error)
	Retrieve(kind uint8, hash hash.Hash) ([]byte, error)
	Insert(entry entries.Entry) error
	Add(addition entries.Addition) error
}
