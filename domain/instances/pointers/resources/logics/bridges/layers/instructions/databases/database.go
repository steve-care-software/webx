package databases

import "github.com/steve-care-software/datastencil/domain/hash"

type database struct {
	hash   hash.Hash
	save   string
	delete string
}

func createDatabaseWithSave(
	hash hash.Hash,
	save string,
) Database {
	return createDatabaseInternally(hash, save, "")
}

func createDatabaseWithDelete(
	hash hash.Hash,
	delete string,
) Database {
	return createDatabaseInternally(hash, "", delete)
}

func createDatabaseInternally(
	hash hash.Hash,
	save string,
	delete string,
) Database {
	out := database{
		hash:   hash,
		save:   save,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// IsSave returns true if there is an save, false otherwise
func (obj *database) IsSave() bool {
	return obj.save != ""
}

// Save returns the save, if any
func (obj *database) Save() string {
	return obj.save
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *database) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *database) Delete() string {
	return obj.delete
}
