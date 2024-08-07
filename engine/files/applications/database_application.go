package applications

import (
	"io"
	"math"
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/cursors/applications/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type fileApplication struct {
	pOriginal    *os.File
	pDestination *os.File
	pLock        *fslock.Lock
	cursor       uint64
}

func createFileApplication(
	pOriginal *os.File,
	pDestination *os.File,
	pLock *fslock.Lock,
) databases.Application {
	out := fileApplication{
		pOriginal:    pOriginal,
		pDestination: pDestination,
		pLock:        pLock,
		cursor:       0,
	}

	return &out
}

// Reset resets the application
func (app *fileApplication) Reset() {
	app.cursor = 0
}

// Read reads data from the file at the delimiter
func (app *fileApplication) Read(delimiter delimiters.Delimiter) ([]byte, error) {
	index := delimiter.Index()
	length := delimiter.Length()
	return app.read(index, length)
}

// CopyBeforeThenWrite copy data from the file before the index if needed, then write the provided data
func (app *fileApplication) CopyBeforeThenWrite(index uint64, bytes []byte) error {
	// copy the data if needed:
	if app.cursor < index {
		originalIndex := app.cursor
		originalLength := index - originalIndex
		amountLoops := int(math.Floor(float64(originalLength / readChunkSize)))
		lastChkIndex := originalIndex
		for i := 0; i < amountLoops; i++ {
			chkIndex := uint64((i * amountLoops)) + originalIndex
			chk, err := app.read(uint64(chkIndex), readChunkSize)
			if err != nil {
				return err
			}

			// write:
			err = app.write(chkIndex, chk)
			if err != nil {
				return err
			}

			// set the last chunk index:
			lastChkIndex = chkIndex
		}

		// write the last chunk if any
		if lastChkIndex < originalIndex {
			// read the last chunk:
			lastChk, err := app.read(uint64(lastChkIndex), originalIndex-lastChkIndex)
			if err != nil {
				return err
			}

			// write it:
			err = app.write(lastChkIndex, lastChk)
			if err != nil {
				return err
			}
		}
	}

	// write the data:
	return app.write(index, bytes)
}

// Close closes the file
func (app *fileApplication) Close() error {
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

func (app *fileApplication) read(index uint64, length uint64) ([]byte, error) {
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

func (app *fileApplication) write(index uint64, bytes []byte) error {
	// seek:
	_, err := app.pDestination.Seek(int64(index), io.SeekStart)
	if err != nil {
		return err
	}

	// write:
	_, err = app.pOriginal.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
