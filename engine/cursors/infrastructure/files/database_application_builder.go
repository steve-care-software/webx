package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/cursors/applications/databases"
)

type databaseApplicationBuilder struct {
	original    []string
	destination []string
}

func createDatabaseApplicationBuilder() databases.FileBuilder {
	out := databaseApplicationBuilder{
		original:    nil,
		destination: nil,
	}

	return &out
}

// Create initializes the builder
func (app *databaseApplicationBuilder) Create() databases.FileBuilder {
	return createDatabaseApplicationBuilder()
}

// WithOriginal adds an original path to the builder
func (app *databaseApplicationBuilder) WithOriginal(original []string) databases.FileBuilder {
	app.original = original
	return app
}

// WithDestination adds a destination path to the builder
func (app *databaseApplicationBuilder) WithDestination(destination []string) databases.FileBuilder {
	app.destination = destination
	return app
}

// Now builds a new Application builder
func (app *databaseApplicationBuilder) Now() (databases.Application, error) {
	if app.original != nil && len(app.original) <= 0 {
		app.original = nil
	}

	if app.destination != nil && len(app.destination) <= 0 {
		app.destination = nil
	}

	if app.original == nil {

	}

	if app.destination == nil {

	}

	// open the origin path:
	originPath := filepath.Join(app.original...)
	pOriginal, err := os.OpenFile(originPath, os.O_RDWR, os.ModeAppend)
	if err != nil {
		str := fmt.Sprintf("failed to open file: %s", err.Error())
		return nil, errors.New(str)
	}

	// if the destination file doesn't exists, create it:
	var pDestination *os.File
	destinationPath := filepath.Join(app.destination...)
	if _, err := os.Stat(destinationPath); os.IsNotExist(err) {
		dir := filepath.Dir(destinationPath)
		err := os.MkdirAll(dir, os.ModePerm) // Create the directory path
		if err != nil {
			return nil, err
		}

		// Create the file
		pDestination, err = os.Create(destinationPath)
		if err != nil {
			return nil, err
		}
	}

	// lock the destination file:
	pLock := fslock.New(destinationPath)
	err = pLock.TryLock()
	if err != nil {
		str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
		return nil, errors.New(str)
	}

	if pDestination == nil {
		pOpenFile, err := os.OpenFile(destinationPath, os.O_RDWR, os.ModeAppend)
		if err != nil {
			str := fmt.Sprintf("failed to open file: %s", err.Error())
			return nil, errors.New(str)
		}

		pDestination = pOpenFile
	}

	return createDatabaseApplication(
		originPath,
		pOriginal,
		pDestination,
		pLock,
	), nil
}
