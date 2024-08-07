package files

import (
	"io"
	"math"
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/cursors/applications/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type databaseApplication struct {
	originalPath string
	pOriginal    *os.File
	pDestination *os.File
	pLock        *fslock.Lock
	cursor       uint64
}

func createDatabaseApplication(
	originalPath string,
	pOriginal *os.File,
	pDestination *os.File,
	pLock *fslock.Lock,
) databases.Application {
	out := databaseApplication{
		originalPath: originalPath,
		pOriginal:    pOriginal,
		pDestination: pDestination,
		pLock:        pLock,
		cursor:       0,
	}

	return &out
}

// Reset resets the application
func (app *databaseApplication) Reset() {
	app.cursor = 0
}

// Read reads data from the file at the delimiter
func (app *databaseApplication) Read(delimiter delimiters.Delimiter) ([]byte, error) {
	index := delimiter.Index()
	length := delimiter.Length()
	return app.read(index, length)
}

// CopyBeforeThenWrite copy data from the file before the index if needed, then write the provided data
func (app *databaseApplication) CopyBeforeThenWrite(startAtIndex uint64, index uint64, bytes []byte) error {
	// copy the data if needed:
	diff := index - startAtIndex
	if diff > 0 {
		amountLoops := int(diff / readChunkSize)
		nextChkIndex := startAtIndex
		for i := 0; i < amountLoops; i++ {
			chk, err := app.read(nextChkIndex, readChunkSize)
			if err != nil {
				return err
			}

			// write:
			err = app.write(nextChkIndex, chk)
			if err != nil {
				return err
			}

			// set the last chunk index:
			nextChkIndex += uint64(len(chk))
		}

		// write the last chunk if any
		remaining := uint64(math.Mod(float64(diff), float64(readChunkSize)))
		if remaining > 0 {
			// read the last chunk:
			lastChk, err := app.read(nextChkIndex, remaining)
			if err != nil {
				return err
			}

			// write it:
			err = app.write(nextChkIndex, lastChk)
			if err != nil {
				return err
			}
		}
	}

	// write the data:
	return app.write(index, bytes)
}

// CloseThenDeleteOriginal closes the files then deletes the original file
func (app *databaseApplication) CloseThenDeleteOriginal() error {
	err := app.Close()
	if err != nil {
		return err
	}

	return os.Remove(app.originalPath)
}

// Close closes the file
func (app *databaseApplication) Close() error {
	if app.pLock != nil {
		err := app.pLock.Unlock()
		if err != nil {
			return err
		}
	}

	if app.pOriginal != nil {
		err := app.pOriginal.Close()
		if err != nil {
			return err
		}
	}

	if app.pDestination != nil {
		err := app.pDestination.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *databaseApplication) read(index uint64, length uint64) ([]byte, error) {
	// seek:
	_, err := app.pOriginal.Seek(int64(index), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// read for the provided length:
	output := make([]byte, length)
	_, err = app.pOriginal.Read(output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (app *databaseApplication) write(index uint64, bytes []byte) error {
	// seek:
	_, err := app.pDestination.Seek(int64(index), io.SeekStart)
	if err != nil {
		return err
	}

	// write:
	_, err = app.pDestination.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
