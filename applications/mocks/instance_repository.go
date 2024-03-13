package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/queries"
)

type instanceRepository struct {
	height    uint
	hashList  []hash.Hash
	instances map[string]instances.Instance
}

func createInstanceRepository(
	height uint,
	hashList []hash.Hash,
	instances map[string]instances.Instance,
) instances.Repository {
	out := instanceRepository{
		height:    height,
		hashList:  hashList,
		instances: instances,
	}

	return &out
}

// Height returns the current commit height
func (app *instanceRepository) Height() (*uint, error) {
	value := app.height
	return &value, nil
}

// List returns the hashes list related to the query
func (app *instanceRepository) List(query queries.Query) ([]hash.Hash, error) {
	if app.hashList == nil {
		return nil, errors.New("the list was requested to fail")
	}

	return app.hashList, nil
}

// Exists returns true if the instance exists, false otherwise
func (app *instanceRepository) Exists(path []string, hash hash.Hash) bool {
	_, err := app.RetrieveByPathAndHash(path, hash)
	if err != nil {
		return false
	}

	return true
}

// Retrieve returns the instance by query
func (app *instanceRepository) Retrieve(query queries.Query) (instances.Instance, error) {
	if ins, ok := app.instances[query.Hash().String()]; ok {
		return ins, nil
	}

	return nil, errors.New("the Retrieve was requested to fail")
}

// RetrieveByPathAndHash returns the instance by path and hash
func (app *instanceRepository) RetrieveByPathAndHash(path []string, hash hash.Hash) (instances.Instance, error) {
	if ins, ok := app.instances[hash.String()]; ok {
		return ins, nil
	}

	return nil, errors.New("the RetrieveByPathAndHash was requested to fail")
}
