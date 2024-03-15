package sqllites

import (
	"database/sql"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/queries"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons"
)

type instanceRepository struct {
	hashAdapter              hash.Adapter
	buildInstances           map[string]buildInstanceFn
	elementsToListInstanceFn map[string]elementsToListInstanceFn
	skeleton                 skeletons.Skeleton
	pDB                      *sql.DB
}

func createInstanceReposiory(
	hashAdapter hash.Adapter,
	buildInstances map[string]buildInstanceFn,
	elementsToListInstanceFn map[string]elementsToListInstanceFn,
	skeleton skeletons.Skeleton,
	pDB *sql.DB,
) instances.Repository {
	out := instanceRepository{
		hashAdapter:              hashAdapter,
		buildInstances:           buildInstances,
		elementsToListInstanceFn: elementsToListInstanceFn,
		skeleton:                 skeleton,
		pDB:                      pDB,
	}
	return &out
}

// Height returns the current commit height
func (app *instanceRepository) Height() (*uint, error) {
	return nil, nil
}

// List returns the hashes list related to the query
func (app *instanceRepository) List(query queries.Query) ([]hash.Hash, error) {
	return nil, nil
}

// ListByPath returns the hashes list related to the path
func (app *instanceRepository) ListByPath(path []string) ([]hash.Hash, error) {
	return nil, nil
}

// Exists returns true if the instance exists, false otherwise
func (app *instanceRepository) Exists(path []string, hash hash.Hash) bool {
	return false
}

// Retrieve returns the instance by query
func (app *instanceRepository) Retrieve(query queries.Query) (instances.Instance, error) {
	return nil, nil
}

// RetrieveByPathAndHash returns the instance by path and hash
func (app *instanceRepository) RetrieveByPathAndHash(path []string, hash hash.Hash) (instances.Instance, error) {
	return nil, nil
}
