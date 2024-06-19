package files

import (
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/instances/pointers"
)

type pointerRepository struct {
	pointersAdapter pointers.Adapter
	pointersBuilder pointers.Builder
	basePath        []string
}

func createPointerRepository(
	pointersAdapter pointers.Adapter,
	pointersBuilder pointers.Builder,
	basePath []string,
) pointers.Repository {
	out := pointerRepository{
		pointersAdapter: pointersAdapter,
		pointersBuilder: pointersBuilder,
		basePath:        basePath,
	}

	return &out
}

// Retrieve retrieve pointers by path
func (app *pointerRepository) Retrieve(path []string) (pointers.Pointers, error) {
	filePath := filepath.Join(append(app.basePath, path...)...)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return app.pointersAdapter.ToInstance(bytes)
}

// Match retrieve pointers by path, then returns the subset that matches the provided history
func (app *pointerRepository) Match(path []string, history [][]string) (pointers.Pointers, error) {
	pointers, err := app.Retrieve(path)
	if err != nil {
		return nil, err
	}

	list := pointers.Match(history)
	return app.pointersBuilder.Create().
		WithList(list).
		Now()
}

// Fetch executes a match and then returns the data of the file where the selected pointer's path points to
func (app *pointerRepository) Fetch(path []string, history [][]string) ([]byte, error) {
	pointers, err := app.Match(path, history)
	if err != nil {
		return nil, err
	}

	basePath := append(app.basePath, path...)
	basePathWithoutEndingFile := basePath[:len(basePath)-1]

	resPath := pointers.First().Path()
	resFullPath := append(basePathWithoutEndingFile, resPath...)
	pathAsString := filepath.Join(resFullPath...)
	return ioutil.ReadFile(pathAsString)
}
