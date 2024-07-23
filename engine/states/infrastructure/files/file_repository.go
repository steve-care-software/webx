package files

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/steve-care-software/webx/engine/states/domain/files"
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

// Open opens a file
func (app *fileRepository) Open(path []string, permission uint) (*os.File, error) {
	filePath := createFilePath(app.basePath, path)
	return os.Open(filePath)
}

// Exists returns true if the file exists, false otherwise
func (app *fileRepository) Exists(path []string) bool {
	filePath := createFilePath(app.basePath, path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// RetrieveChunk retrieves a portion of data from a file
func (app *fileRepository) RetrieveChunk(identifier *os.File, index uint, length uint) ([]byte, error) {
	startIndex := int64(index)
	_, err := identifier.Seek(startIndex, 0)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, int64(length))
	_, err = identifier.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// RetrieveFrom retrieves data from an index til the eof
func (app *fileRepository) RetrieveFrom(identifier *os.File, index uint) ([]byte, error) {
	startIndex := int64(index)
	_, err := identifier.Seek(startIndex, 0)
	if err != nil {
		return nil, err
	}

	buffer, err := io.ReadAll(identifier)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// RetrieveAll retrieves all data from identifier
func (app *fileRepository) RetrieveAll(identifier *os.File) ([]byte, error) {
	buffer, err := io.ReadAll(identifier)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// RetrieveFromPath retrieves all bytes from path
func (app *fileRepository) RetrieveFromPath(path []string) ([]byte, error) {
	filePath := createFilePath(app.basePath, path)
	if !app.Exists(path) {
		str := fmt.Sprintf("the file (path: %s) does not exists", filePath)
		return nil, errors.New(str)
	}

	return os.ReadFile(filePath)
}
