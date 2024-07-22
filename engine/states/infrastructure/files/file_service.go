package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/states/domain/files"
)

type fileService struct {
	repository files.Repository
	basePath   []string
	locks      map[string]*fslock.Lock
}

func createFileService(
	repository files.Repository,
	basePath []string,
) files.Service {
	out := fileService{
		repository: repository,
		basePath:   basePath,
		locks:      map[string]*fslock.Lock{},
	}

	return &out
}

// Init initializes a file
func (app *fileService) Init(path []string) error {
	filePath := createFilePath(app.basePath, path[0:len(path)-1])
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// Lock locks a file
func (app *fileService) Lock(path []string) error {
	filePath := createFilePath(app.basePath, path)
	lock := fslock.New(filePath)
	err := lock.TryLock()
	if err != nil {
		str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
		return errors.New(str)
	}

	app.locks[filePath] = lock
	return nil
}

// Unlock unlocks a file
func (app *fileService) Unlock(path []string) error {
	filePath := createFilePath(app.basePath, path)
	if pLock, ok := app.locks[filePath]; ok {
		pLock.Unlock()
		return nil
	}

	str := fmt.Sprintf("there is no lock on the provided file path: %s", filePath)
	return errors.New(str)
}

// Save saves data in a file
func (app *fileService) Save(path []string, bytes []byte) error {
	filePath := createFilePath(app.basePath, path)
	return os.WriteFile(filePath, bytes, fs.ModePerm)
}

// Transact transact bytes to a file
func (app *fileService) Transact(path []string, bytes []byte) error {
	if !app.repository.Exists(path) {
		err := app.Init(path)
		if err != nil {
			return err
		}
	}

	err := app.Lock(path)
	if err != nil {
		return err
	}

	defer app.Unlock(path)
	return app.Save(path, bytes)
}
