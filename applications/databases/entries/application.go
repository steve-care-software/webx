package entries

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entries"
	"github.com/steve-care-software/webx/domain/databases/references"
)

type application struct {
	builder          entries.Builder
	entryBuilder     entries.EntryBuilder
	relationsBuilder entries.RelationsBuilder
	relationBuilder  entries.RelationBuilder
	linksBuilder     entries.LinksBuilder
	linkBuilder      entries.LinkBuilder
	additionBuilder  entries.AdditionBuilder
	reference        references.Reference
	absFilePath      string
}

func createApplication(
	builder entries.Builder,
	entryBuilder entries.EntryBuilder,
	relationsBuilder entries.RelationsBuilder,
	relationBuilder entries.RelationBuilder,
	linksBuilder entries.LinksBuilder,
	linkBuilder entries.LinkBuilder,
	additionBuilder entries.AdditionBuilder,
	reference references.Reference,
	absFilePath string,
) Application {
	out := application{
		builder:          builder,
		entryBuilder:     entryBuilder,
		relationsBuilder: relationsBuilder,
		relationBuilder:  relationBuilder,
		linksBuilder:     linksBuilder,
		linkBuilder:      linkBuilder,
		additionBuilder:  additionBuilder,
		reference:        reference,
		absFilePath:      absFilePath,
	}

	return &out
}

// List lists the content by kind
func (app *application) List(kind uint8) ([][]byte, error) {
	return nil, nil
}

// List lists the content by master entry and slave kind
func (app *application) ListByLink(masterKind uint8, masterHash hash.Hash, slaveKind uint8) ([][]byte, error) {
	return nil, nil
}

// Retrieve retrieves an entry by hash
func (app *application) Retrieve(kind uint8, hash hash.Hash) ([]byte, error) {
	return nil, nil
}

// Insert inserts an entry
func (app *application) Insert(entry entries.Entry) error {
	return nil
}

// Add inserts an addition to an entry
func (app *application) Add(addition entries.Addition) error {
	return nil
}
