package databases

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/steve-care-software/webx/applications/databases"
	"github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/applications/databases/transactions"
	"github.com/steve-care-software/webx/domain/databases/references"
)

type application struct {
	contentAppBuilder contents.Builder
	trxAppBuilder     transactions.Builder
	referenceAdapter  references.Adapter
	referenceFactory  references.Factory
	dirPath           string
}

func createApplication(
	contentAppBuilder contents.Builder,
	trxAppBuilder transactions.Builder,
	referenceAdapter references.Adapter,
	referenceFactory references.Factory,
	dirPath string,
) databases.Application {
	out := application{
		contentAppBuilder: contentAppBuilder,
		trxAppBuilder:     trxAppBuilder,
		referenceAdapter:  referenceAdapter,
		referenceFactory:  referenceFactory,
		dirPath:           dirPath,
	}

	return &out
}

// List lists the databases
func (app *application) List() ([]string, error) {
	files, err := ioutil.ReadDir(app.dirPath)
	if err != nil {
		return []string{}, nil
	}

	output := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		output = append(output, file.Name())
	}

	return output, nil
}

// New creates a new database by name
func (app *application) New(name string) error {
	// if the directory path does not exists, create it:
	if _, err := os.Stat(app.dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(app.dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// create the path:
	path := filepath.Join(app.dirPath, name)

	// make sure the file doesn't already exists:
	if _, err := os.Stat(path); os.IsExist(err) {
		str := fmt.Sprintf("the name (%s) already exists", name)
		return errors.New(str)
	}

	// generate a reference instance:
	referenceIns, err := app.referenceFactory.Create()
	if err != nil {
		return err
	}

	// convert the reference to content:
	referenceContent, err := app.referenceAdapter.ToContent(referenceIns)
	if err != nil {
		return err
	}

	// declare the data array:
	data := []byte{}

	// add the reference's content length to the bytes array:
	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(len(referenceContent)))
	data = append(data, lengthBytes...)

	// add the reference's content to the array:
	data = append(data, referenceContent...)

	// write the bites on disk:
	ioutil.WriteFile(path, data, 0644)
	return nil
}

// Content returns a content application using the given name
func (app *application) Content(name string) (contents.Application, error) {
	return app.contentAppBuilder.Create().
		WithName(name).
		Now()
}

// Transaction returns a transaction application using the given name
func (app *application) Transaction(name string) (transactions.Application, error) {
	return app.trxAppBuilder.Create().
		WithName(name).
		Now()
}
