package files

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/steve-care-software/datastencil/states/domain/files"
)

type fileRepository struct {
	basePath []string
}

func createFileRepository(
	basePath []string,
) files.Repository {
	out := fileRepository{
		basePath: basePath,
	}

	return &out
}

// Exists returns true if the file exists, false otherwise
func (app *fileRepository) Exists(path []string) bool {
	filePath := createFilePath(app.basePath, path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// Retrieve retrieves bytes by path
func (app *fileRepository) Retrieve(path []string) ([]byte, error) {
	filePath := createFilePath(app.basePath, path)
	if !app.Exists(path) {
		str := fmt.Sprintf("the file (path: %s) does not exists", filePath)
		return nil, errors.New(str)
	}

	return ioutil.ReadFile(filePath)
}
