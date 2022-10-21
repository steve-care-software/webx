package databases

import (
	"github.com/steve-care-software/webx/domain/databases/blockchains"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/schemas"
)

type content struct {
	hash      hash.Hash
	name      string
	schema    schemas.Schema
	reference blockchains.Blockchain
	encryptTo keys.PublicKey
	migration Migration
}

func createContent(
	hash hash.Hash,
	name string,
	schema schemas.Schema,
	reference blockchains.Blockchain,
	encryptTo keys.PublicKey,
) Content {
	return createContentInternally(hash, name, schema, reference, encryptTo, nil)
}

func createContentWithMigration(
	hash hash.Hash,
	name string,
	schema schemas.Schema,
	reference blockchains.Blockchain,
	encryptTo keys.PublicKey,
	migration Migration,
) Content {
	return createContentInternally(hash, name, schema, reference, encryptTo, migration)
}

func createContentInternally(
	hash hash.Hash,
	name string,
	schema schemas.Schema,
	reference blockchains.Blockchain,
	encryptTo keys.PublicKey,
	migration Migration,
) Content {
	out := content{
		hash:      hash,
		name:      name,
		schema:    schema,
		reference: reference,
		encryptTo: encryptTo,
		migration: migration,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *content) Name() string {
	return obj.name
}

// Schema returns the schema
func (obj *content) Schema() schemas.Schema {
	return obj.schema
}

// Reference returns the reference blockchain
func (obj *content) Reference() blockchains.Blockchain {
	return obj.reference
}

// EncryptoTo returns the encryption public key
func (obj *content) EncryptoTo() keys.PublicKey {
	return obj.encryptTo
}

// HasMigration returns true if there is a migration, false otherwise
func (obj *content) HasMigration() bool {
	return obj.migration != nil
}

// Migration returns the migration, if any
func (obj *content) Migration() Migration {
	return obj.migration
}
