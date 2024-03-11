package repositories

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isSkeleton  bool
	isHeight    bool
	list        string
	retrieve    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isSkeleton:  false,
		isHeight:    false,
		list:        "",
		retrieve:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list string) Builder {
	app.list = list
	return app
}

// WithRetrieve adds a retrieve to the builder
func (app *builder) WithRetrieve(retrieve string) Builder {
	app.retrieve = retrieve
	return app
}

// IsSkeleton flags the builder as a skeleton
func (app *builder) IsSkeleton() Builder {
	app.isSkeleton = true
	return app
}

// IsHeight flags the builder as a height
func (app *builder) IsHeight() Builder {
	app.isHeight = true
	return app
}

// Nwo builds a new Repository instance
func (app *builder) Now() (Repository, error) {
	data := [][]byte{}
	if app.isSkeleton {
		data = append(data, []byte("isSkeleton"))
	}

	if app.isHeight {
		data = append(data, []byte("isHeight"))
	}

	if app.retrieve != "" {
		data = append(data, []byte("retrieve"))
		data = append(data, []byte(app.retrieve))
	}

	if app.list != "" {
		data = append(data, []byte("list"))
		data = append(data, []byte(app.list))
	}

	if app.isSkeleton || app.isHeight {
		if len(data) != 1 {
			return nil, errors.New("the Repository is invalid")
		}
	}

	if app.retrieve != "" || app.list != "" {
		if len(data) != 2 {
			return nil, errors.New("the Repository is invalid")
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isSkeleton {
		return createRepositoryWithSkeleton(*pHash), nil
	}

	if app.isHeight {
		return createRepositoryWithHeight(*pHash), nil
	}

	if app.retrieve != "" {
		return createRepositoryWithRetrieve(*pHash, app.retrieve), nil
	}

	return createRepositoryWithList(*pHash, app.list), nil
}
